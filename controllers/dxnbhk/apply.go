package dxnbhk

import (
	apibackend "LianFaPhone/lfp-api/errdef"
	. "LianFaPhone/lfp-base/log/zap"
	"LianFaPhone/lfp-marketing-api/api"
	"LianFaPhone/lfp-marketing-api/config"
	. "LianFaPhone/lfp-marketing-api/controllers"
	"LianFaPhone/lfp-marketing-api/models"
	"LianFaPhone/lfp-marketing-api/thirdcard-api/dxnbhk"
	"github.com/kataras/iris"
	"go.uber.org/zap"
	"strings"
	"sync"
	"time"
)

type Dxnbhk struct{
	Controllers
	fastApplyLocker sync.Mutex
}

func (this *Dxnbhk) FastApply(ctx iris.Context) {
	param := new(api.FtDxnbhkFastApply)
	err := Tools.ShouldBindJSON(ctx, param)
	if err != nil {
		this.ExceptionSerive(ctx, apibackend.BASERR_INVALID_PARAMETER.Code(), apibackend.BASERR_INVALID_PARAMETER.Desc())
		ZapLog().Error("param err", zap.Error(err))
		return
	}

	var partnerGoods *models.PdPartnerGoods
	var reseller *config.DxnbhkReseller
	{ //怕又bug导致死锁
		this.fastApplyLocker.Lock()
		defer this.fastApplyLocker.Unlock()
		for i:=0;i< len(config.GConfig.Dxnbhk.Partners); i++{
			if config.GConfig.Dxnbhk.Partners[config.GConfig.Dxnbhk.PartnerIndex].Count < config.GConfig.Dxnbhk.Partners[config.GConfig.Dxnbhk.PartnerIndex].MaxNum {
				//config.GConfig.Dxnbhk.PartnerIndex = i
				config.GConfig.Dxnbhk.Partners[config.GConfig.Dxnbhk.PartnerIndex].Count++
				reseller = &config.GConfig.Dxnbhk.Partners[config.GConfig.Dxnbhk.PartnerIndex]
				partnerGoods,err = new(models.PdPartnerGoods).GetByCodeFromCache(reseller.PartnerGoodsCode)
				if err != nil {
					continue
				}
				if partnerGoods == nil {
					ZapLog().Error("fastapply", zap.String("code", reseller.PartnerGoodsCode), zap.Int("len", len(config.GConfig.Dxnbhk.Partners)),zap.Int("i", i))
					config.GConfig.Dxnbhk.PartnerIndex= (config.GConfig.Dxnbhk.PartnerIndex+1)%len(config.GConfig.Dxnbhk.Partners)
					continue
				}
				break
			}
			config.GConfig.Dxnbhk.Partners[config.GConfig.Dxnbhk.PartnerIndex].Count = 0
			config.GConfig.Dxnbhk.PartnerIndex= (config.GConfig.Dxnbhk.PartnerIndex+1)%len(config.GConfig.Dxnbhk.Partners)
		}
	}

	if partnerGoods == nil || reseller == nil {
		ZapLog().Error("no find partnerGoods", zap.Any("pds", partnerGoods), zap.Any("resller", reseller))
		new(models.CardOrderLog).FtParseAdd2(nil, &param.OrderNo, "电信宁波花卡|快速下单失败|商品未找到").Add()
		this.ExceptionSerive(ctx, apibackend.BASERR_UNKNOWN_BUG.Code(), "商品未匹配")
		return
	}

	newOrder := &models.CardOrder{
		OrderNo:          &param.OrderNo,
		Isp:              new(int),
		PartnerId:        partnerGoods.PartnerId,
		PartnerGoodsCode: partnerGoods.Code,
		Status:           new(int),
		ThirdOrderAt:     new(int64),
	}
	*newOrder.ThirdOrderAt = time.Now().Unix()
	*newOrder.Isp = models.CONST_ISP_Dianxin
	*newOrder.Status = models.CONST_OrderStatus_New_Apply_Doing

	count,err := newOrder.UpdatesByOrderNo();
	if err != nil {
		ZapLog().With(zap.Error(err)).Error("database err")
		new(models.CardOrderLog).FtParseAdd2(nil, &param.OrderNo, "电信宁波花卡|快速下单成功|更新状态失败，"+err.Error()).Add()
		this.ExceptionSerive(ctx, apibackend.BASERR_DATABASE_ERROR.Code(), apibackend.BASERR_DATABASE_ERROR.Desc())
	}

	if count <=0 {
		this.ExceptionSerive(ctx, apibackend.BASERR_OBJECT_NOT_FOUND.Code(), "订单未找到")
		return
	}

	this.Response(ctx, nil)
	go this.aysnFastApply(ctx, param.OrderNo, partnerGoods, reseller)

}

func (this *Dxnbhk) aysnFastApply(ctx iris.Context, orderNo string,  partnerGoods *models.PdPartnerGoods, reseller *config.DxnbhkReseller) {
	order,err := new(models.CardOrder).GetByOrderNo(orderNo)
	if err != nil {
		new(models.CardOrderLog).FtParseAdd2(order.Id, order.OrderNo, "电信宁波花卡|快速下单失败|数据库错误,"+err.Error()).Add()
		//this.ExceptionSerive(ctx, apibackend.BASERR_OBJECT_DATA_NOT_FOUND.Code(), "数据不完整")
		return
	}
	if order == nil {
		new(models.CardOrderLog).FtParseAdd2(order.Id, order.OrderNo, "电信宁波花卡|快速下单失败|订单未找到").Add()
		//this.ExceptionSerive(ctx, apibackend.BASERR_OBJECT_DATA_NOT_FOUND.Code(), "数据不完整")
		return
	}


	if order.Province == nil || order.City ==nil || order.Area == nil || len(*order.Province) < 2 || len(*order.City) < 3{
		new(models.CardOrderLog).FtParseAdd2(order.Id, order.OrderNo, "电信宁波花卡|快速下单失败|数据不完整").Add()
		//this.ExceptionSerive(ctx, apibackend.BASERR_OBJECT_DATA_NOT_FOUND.Code(), "数据不完整")
		new(models.CardOrder).UpdateStatusByOrderNo(orderNo, models.CONST_OrderStatus_Fail_Already_Retry)
		return
	}

	if order.Status == nil {
		new(models.CardOrderLog).FtParseAdd2(order.Id, order.OrderNo, "电信宁波花卡|快速下单失败|数据不完整").Add()
		//this.ExceptionSerive(ctx, apibackend.BASERR_UNKNOWN_BUG.Code(), "系统错误")
		new(models.CardOrder).UpdateStatusByOrderNo(orderNo, models.CONST_OrderStatus_Fail_Already_Retry)

		return
	}

	if (*order.Status < models.CONST_OrderStatus_MinFail) && (*order.Status > models.CONST_OrderStatus_MaxFail) {
		new(models.CardOrderLog).FtParseAdd2(order.Id, order.OrderNo, "电信宁波花卡|快速下单失败|状态不允许").Add()
		//this.ExceptionSerive(ctx, apibackend.BASERR_OBJECT_EXISTS.Code(), "订单已完成")
		new(models.CardOrder).UpdateStatusByOrderNo(orderNo, models.CONST_OrderStatus_Fail_Already_Retry)
		return
	}

	provinceName :=""
	if strings.HasPrefix(*order.Province, "内蒙古") {
		provinceName = "内蒙古"
	}else if strings.HasPrefix(*order.Province, "黑龙江") {
		provinceName = "黑龙江"
	}else{
		temp := []rune(*order.Province)
		provinceName= string(temp[0:2])
	}

	province,err := new(models.DxnbhkProvice).GetByName(provinceName)
	if err != nil {
		new(models.CardOrderLog).FtParseAdd2(order.Id, order.OrderNo, "电信宁波花卡|快速下单失败|数据库异常").Add()
		ZapLog().With(zap.Error(err)).Error("database err")
		new(models.CardOrder).UpdateStatusByOrderNo(orderNo, models.CONST_OrderStatus_Fail_Already_Retry)
		return
	}
	if province == nil {
		new(models.CardOrderLog).FtParseAdd2(order.Id, order.OrderNo, "电信宁波花卡|快速下单失败|省份匹配不上").Add()
		new(models.CardOrder).UpdateStatusByOrderNo(orderNo, models.CONST_OrderStatus_Fail_Already_Retry)
		return
	}

	city,err := new(models.DxnbhkCity).GetByProvinceIdAndName(*province.Id, *order.City)
	if err != nil {
		new(models.CardOrderLog).FtParseAdd2(order.Id, order.OrderNo, "电信宁波花卡|快速下单失败|数据库异常").Add()
		//ZapLog().With(zap.Error(err)).Error("database err")
		new(models.CardOrder).UpdateStatusByOrderNo(orderNo, models.CONST_OrderStatus_Fail_Already_Retry)
		return
	}
	if city == nil {
		new(models.CardOrderLog).FtParseAdd2(order.Id, order.OrderNo, "电信宁波花卡|快速下单失败|城市匹配不上,"+*order.City).Add()
		new(models.CardOrder).UpdateStatusByOrderNo(orderNo, models.CONST_OrderStatus_Fail_Already_Retry)
		return
	}

	area,err := new(models.DxnbhkArea).GetByCityIdAndName(*city.Id, *order.Area)
	if err != nil {
		new(models.CardOrderLog).FtParseAdd2(order.Id, order.OrderNo, "电信宁波花卡|快速下单失败|数据库异常").Add()
		//ZapLog().With(zap.Error(err)).Error("database err")
		new(models.CardOrder).UpdateStatusByOrderNo(orderNo, models.CONST_OrderStatus_Fail_Already_Retry)
		return
	}
	if area == nil {
		new(models.CardOrderLog).FtParseAdd2(order.Id, order.OrderNo, "电信宁波花卡|快速下单失败|区县匹配不上,"+*order.Area).Add()
		new(models.CardOrder).UpdateStatusByOrderNo(orderNo, models.CONST_OrderStatus_Fail_Already_Retry)
		return
	}

	reOrderSubmit := &dxnbhk.ReOrderSubmit{
		ContactNumber: *order.Phone,
		UserName    : *order.TrueName,
		IdNumber  : *order.IdCard,
		AgentMark   :reseller.ResellerId,
		ProductId   : reseller.ProductId,
		Province :  provinceName,
		City :      *city.Name,
		Area    : *area.Name,
		Address   : *order.Address,
	}

	//var lastErr error
	var res *dxnbhk.ResOrderSubmit
	sucFlag := false
	res,err = reOrderSubmit.Send()
	if err != nil {
		ZapLog().With(zap.Error(err)).Error("reOrderSubmit send err")
		new(models.CardOrderLog).FtParseAdd2(order.Id, order.OrderNo, "电信宁波花卡|快速下单失败|对接电信, "+err.Error()).Add()
		new(models.CardOrder).UpdateStatusByOrderNo(orderNo, models.CONST_OrderStatus_Fail_Already_Retry)
		return
	}

	switch res.Success.(type) {
		case bool:
			sucFlag,_ = res.Success.(bool)
		case string:
			sucStr,_ := res.Success.(string)
			if sucStr == "true" {
				sucFlag = true
			}
	}

	if sucFlag == false {
		ZapLog().With(zap.String("errMsg", res.Msg)).Error("reOrderSubmit send err")
		new(models.CardOrderLog).FtParseAdd2(order.Id, order.OrderNo, "电信宁波花卡|快速下单失败|对接电信, "+res.Msg).Add()
		new(models.CardOrder).UpdateStatusByOrderNo(orderNo, models.CONST_OrderStatus_Fail_Already_Retry)
		return
	}


	new(models.CardOrderLog).FtParseAdd2(order.Id, order.OrderNo, "电信宁波花卡|快速下单成功|"+res.Msg).Add()

	newOrder := &models.CardOrder{
		OrderNo:          &orderNo,
		Status:           new(int),
		NewPhone:         &res.SelectNumber,
		ThirdOrderNo:     &res.OrderNumber,
		ThirdResp:        nil,
		ThirdOrderAt:     new(int64),
	}
	*newOrder.ThirdOrderAt = time.Now().Unix()
	*newOrder.Status = models.CONST_OrderStatus_New

	if _,err := newOrder.UpdatesByOrderNo(); err != nil {
			ZapLog().With(zap.Error(err)).Error("database err")
			new(models.CardOrderLog).FtParseAdd2(order.Id, order.OrderNo, "电信宁波花卡|快速下单成功|更新状态失败，"+err.Error()).Add()
	}

}
