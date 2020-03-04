package jthk

import (
	"LianFaPhone/lfp-marketing-api/api"
	"LianFaPhone/lfp-marketing-api/common"
	"LianFaPhone/lfp-marketing-api/config"
	"LianFaPhone/lfp-marketing-api/models"
	. "LianFaPhone/lfp-marketing-api/thirdcard-api/ydhk"
	"fmt"
	"go.uber.org/zap"
	"net/url"
	"strconv"
	"strings"
	"time"
)

import (
	. "LianFaPhone/lfp-marketing-api/controllers"
	"github.com/kataras/iris"
)
import (
	apibackend "LianFaPhone/lfp-api/errdef"
	. "LianFaPhone/lfp-base/log/zap"
)

type Ydhk struct{
	Controllers
}

func (this * Ydhk) GetProtocal(ctx iris.Context) {
	isoao,_ := ctx.URLParamBool("isOao")
	pCode := ctx.FormValue("province_code")
	token,err := new(ReProtocal).Send(isoao, pCode)
	if err != nil {
		ZapLog().With(zap.Error(err)).Error("Retoken send err")
		this.ExceptionSerive(ctx, apibackend.BASERR_INTERNAL_SERVICE_ACCESS_ERROR.Code(), err.Error())
		return
	}
	this.Response(ctx, token)
}

func (this * Ydhk) GetToken(ctx iris.Context) {
	isoao,_ := ctx.URLParamBool("isOao")
	channelId := ctx.URLParam("channelId")
	if len(channelId) == 0 {
		code := ctx.URLParam("code")
		if len(code) == 0 {
			this.ExceptionSerive(ctx, apibackend.BASERR_INVALID_PARAMETER.Code(), apibackend.BASERR_INVALID_PARAMETER.Desc())
			ZapLog().Error("param err")
			return
		}
		ppd, err := new(models.PdPartnerGoods).GetByCodeFromCache(code)
		if err != nil {
			this.ExceptionSerive(ctx, apibackend.BASERR_DATABASE_ERROR.Code(), apibackend.BASERR_DATABASE_ERROR.Desc())
			ZapLog().Error("database err", zap.Error(err))
			return
		}
		if ppd == nil {
			this.ExceptionSerive(ctx, apibackend.BASERR_OBJECT_NOT_FOUND.Code(), apibackend.BASERR_OBJECT_NOT_FOUND.Desc())
			ZapLog().Error("nofind err")
			return
		}
		if ppd.UrlParam == nil {
			this.ExceptionSerive(ctx, apibackend.BASERR_OBJECT_DATA_NOT_FOUND.Code(), apibackend.BASERR_OBJECT_DATA_NOT_FOUND.Desc())
			ZapLog().Error("urlparam nofind err")
			return
		}
		vv, _ := url.ParseQuery(*ppd.UrlParam)
		channelId = vv.Get("channelId")
		if len(channelId) == 0 {
			this.ExceptionSerive(ctx, apibackend.BASERR_OBJECT_DATA_NOT_FOUND.Code(), apibackend.BASERR_OBJECT_DATA_NOT_FOUND.Desc())
			ZapLog().Error("channelId nofind err")
			return
		}
		isoao,_ = strconv.ParseBool(vv.Get("isOao"))
	}

	token,err := new(ReToken).Send(isoao, channelId)
	if err != nil {
		ZapLog().With(zap.Error(err)).Error("Retoken send err")
		this.ExceptionSerive(ctx, apibackend.BASERR_INTERNAL_SERVICE_ACCESS_ERROR.Code(), err.Error())
		return
	}
	this.Response(ctx, &api.FtResYdhkToken{token})
}

func (this * Ydhk) GetsAddr(ctx iris.Context) {
	isoao,_ := ctx.URLParamBool("isOao")
	//channelTp := ctx.URLParam("channelType")
	ll, err := new(ReAddr).Send(isoao)
	if err != nil {
		ZapLog().With(zap.Error(err)).Error("Update err")
		this.ExceptionSerive(ctx, apibackend.BASERR_INTERNAL_SERVICE_ACCESS_ERROR.Code(), err.Error())
		return
	}
	this.Response(ctx, ll)
}

func (this * Ydhk) ListNumberPool(ctx iris.Context) {
	isoao,_ := ctx.URLParamBool("isOao")
	param := new(api.FtYdhkNumberPoolList)
	if err := ctx.ReadJSON(param); err != nil {
		this.ExceptionSerive(ctx, apibackend.BASERR_INVALID_PARAMETER.Code(), apibackend.BASERR_INVALID_PARAMETER.Desc())
		ZapLog().Error("param err", zap.Error(err))
		return
	}
	if param.Page >=1 {
		param.Page = param.Page -1
	}

	ll,err := new(ReCardSearch).Send(isoao, param.ProviceCode, param.Provice, param.CityCode, param.City, param.Searchkey, param.Page, param.Size)
	if err != nil {
		ZapLog().With(zap.Error(err)).Error("Retoken send err")
		this.ExceptionSerive(ctx, apibackend.BASERR_INTERNAL_SERVICE_ACCESS_ERROR.Code(), err.Error())
		return
	}
	this.Response(ctx, ll)
}

func (this * Ydhk) LockNumber(ctx iris.Context) {
	isoao,_ := ctx.URLParamBool("isOao")
	param := new(api.FtYdhkNumberPoolLock)
	if err := ctx.ReadJSON(param); err != nil {
		this.ExceptionSerive(ctx, apibackend.BASERR_INVALID_PARAMETER.Code(), apibackend.BASERR_INVALID_PARAMETER.Desc())
		ZapLog().Error("param err", zap.Error(err))
		return
	}

	flag,err := new(ReCloseNumber).Send(isoao, param.ProviceCode, param.CityCode,  param.Number, param.Token)
	if err != nil {
		ZapLog().With(zap.Error(err)).Error("Retoken send err")
		this.ExceptionSerive(ctx, apibackend.BASERR_CARDMARKET_PHONEPOOL_LOCK_FAIL.Code(), err.Error())
		return
	}
	if !flag {
		this.ExceptionSerive(ctx, apibackend.BASERR_CARDMARKET_PHONEPOOL_LOCK_FAIL.Code(), apibackend.BASERR_CARDMARKET_PHONEPOOL_LOCK_FAIL.OriginDesc())
		return
	}
	this.Response(ctx, nil)
}

func (this * Ydhk) Apply(ctx iris.Context) {
	isoao,_ := ctx.URLParamBool("isOao")
//	productType := ctx.URLParamDefault("productType", "19")
	productId := ctx.URLParam("productId")
	channelId := ctx.URLParam("channelId")

	param := new(api.FtYdhkApply) //不对过多参数做检测，提高性能，前端必须做好检测。
	if err := ctx.ReadJSON(param); err != nil {
		this.ExceptionSerive(ctx, apibackend.BASERR_INVALID_PARAMETER.Code(), apibackend.BASERR_INVALID_PARAMETER.Desc())
		ZapLog().Error("param err", zap.Error(err))
		return
	}
	if len(param.CertificateNo) <=0 || len(param.Phone) <= 0 {
		this.ExceptionSerive(ctx, apibackend.BASERR_INVALID_PARAMETER.Code(), apibackend.BASERR_INVALID_PARAMETER.Desc())
		ZapLog().Error("param err")
		return
	}

	if (len(channelId) == 0) && (param.PartnerGoodsCode != nil) {
		ppg,err := new(models.PdPartnerGoods).GetByCodeFromCache(*param.PartnerGoodsCode)
		if err != nil {
			this.ExceptionSerive(ctx, apibackend.BASERR_DATABASE_ERROR.Code(), apibackend.BASERR_DATABASE_ERROR.Desc())
			ZapLog().Error("datebase err", zap.Error(err))
			return
		}
		if ppg == nil {
			this.ExceptionSerive(ctx, apibackend.BASERR_OBJECT_NOT_FOUND.Code(), apibackend.BASERR_OBJECT_NOT_FOUND.Desc())
			ZapLog().Error("nofind ppg err")
			return
		}
		if ppg.UrlParam == nil {
			this.ExceptionSerive(ctx, apibackend.BASERR_OBJECT_DATA_NOT_FOUND.Code(), apibackend.BASERR_OBJECT_DATA_NOT_FOUND.Desc())
			ZapLog().Error("nofind urlparam err")
			return
		}
		vv, _ := url.ParseQuery(*ppg.UrlParam)
		channelId = vv.Get("channelId")
		productId = vv.Get("productId")
		isoao,_ = strconv.ParseBool(vv.Get("isOao"))
		if len(channelId) ==0 || len(productId) == 0 {
			this.ExceptionSerive(ctx, apibackend.BASERR_OBJECT_DATA_NOT_FOUND.Code(), apibackend.BASERR_OBJECT_DATA_NOT_FOUND.Desc())
			ZapLog().Error("nofind urlparam err")
			return
		}
	}

	errCode, orderId,oaoFlag,orderErr := new(ReOrderSubmit).Parse(channelId, productId, nil).Send(isoao, param.AccessToken, param.Phone,  param.NewPhone, param.LeagalName, param.CertificateNo, param.Address, param.Province, param.City, param.SendProvince, param.SendCity,param.SendDistrict)
	if orderErr != nil {
		ZapLog().With(zap.Error(orderErr)).Error("Retoken send err")
		this.ExceptionSerive(ctx, errCode.Code(), orderErr.Error())
		if errCode != apibackend.BASERR_CARDMARKET_PHONECARD_APPLY_FAID_AND_SHOW {
			return
		}
	}else{
		this.Response(ctx, &api.FtResYdhkApply{orderId,oaoFlag})
	}

	go func(){
		oldOrder,_ := new(models.CardOrder).GetByIdcardAndNewPhone(param.CertificateNo, param.NewPhone, &models.SqlPairCondition{"created_at > ?", time.Now().Unix() - 300})
		if oldOrder != nil { // 老订单
			this.recordOldOrder(oldOrder, errCode, oaoFlag, orderErr)
		}else{
			this.recordNewOrder(ctx, param, orderId, errCode, oaoFlag, orderErr)
		}

	}()

}

func (this * Ydhk) recordNewOrder(ctx iris.Context, param *api.FtYdhkApply, thirdOrderId string, errCode apibackend.EnumBasErr, oaoFlag bool, orderErr error) {
	orderNo := fmt.Sprintf("D%s%s%03d", config.GConfig.Server.DevId,time.Now().Format("060102030405000"), GIdGener.Gen())
	modelParam := &models.CardOrder{
		OrderNo:  &orderNo,
		TrueName: &param.LeagalName,
		IdCard:   &param.CertificateNo,
		Isp: param.ClassISP,
		PartnerId : param.PartnerId,              //手机卡套餐类型
		PartnerGoodsCode : param.PartnerGoodsCode,           //手机卡套餐类型

		//		CountryCode: p.CountryCode,
		Phone:        &param.Phone,
		Province:     &param.SendProvinceName,
		ProvinceCode: &param.Province,
		City:         &param.SendCityName,
		CityCode:     &param.City,
		Area:         &param.SendDistrictName,
		AreaCode:     &param.SendDistrict,
		//Town:         p.Town,
		Address:      &param.Address,
		IP:           param.IP,
		PhoneOSTp:    param.PhoneOSTp,

		NewPhone: &param.NewPhone,
		ThirdOrderNo: &thirdOrderId,
	}

	modelParam.Valid = new(int)
	*modelParam.Valid = 1



	if modelParam.PartnerId == nil || modelParam.Isp == nil {
		cc,err := new(models.PdPartnerGoods).GetByCodeFromCache(*param.PartnerGoodsCode)
		if err ==nil && cc != nil {
			modelParam.PartnerId = cc.PartnerId
			if modelParam.Isp == nil {
				partner,_ := new(models.PdPartner).GetByIdFromCache(*modelParam.PartnerId)
				if partner != nil {
					modelParam.Isp = partner.Isp
				}
			}
		}
	}

	if modelParam.IP == nil {
		IP := common.GetRealIp(ctx)
		modelParam.IP = &IP
	}

	if modelParam.PhoneOSTp == nil {
		device := ctx.GetHeader("User-Agent")
		if strings.Contains(device, "Android") {
			modelParam.PhoneOSTp = new(int)
			*modelParam.PhoneOSTp =  models.CONST_PHONEOS_Android
		}else if  strings.Contains(device, "iPhone")  {
			modelParam.PhoneOSTp = new(int)
			*modelParam.PhoneOSTp =  models.CONST_PHONEOS_Iphone
		}else if  strings.Contains(device, "iPad")  {
			modelParam.PhoneOSTp = new(int)
			*modelParam.PhoneOSTp =  models.CONST_PHONEOS_Ipad
		} else {
			modelParam.PhoneOSTp = new(int)
			*modelParam.PhoneOSTp =  models.CONST_PHONEOS_Other
		}
	}

	if modelParam.City != nil {
		pp,err := new(models.BsCity).GetByName(*modelParam.City)
		if err ==nil && pp != nil {
			modelParam.CityCode = pp.Code
		}
	}

	if modelParam.Area != nil {
		pp,err := new(models.BsArea).GetByName(*modelParam.Area)
		if err ==nil && pp != nil {
			modelParam.AreaCode = pp.Code
		}
	}

	modelParam.Status = new(int)
	if errCode != 0 {
		*modelParam.Status =	models.CONST_OrderStatus_Fail
	}else{
		if oaoFlag == true {
			*modelParam.Status =	models.CONST_OrderStatus_New
		}else{
			*modelParam.Status =	models.CONST_OrderStatus_New_UnFinish
		}
	}

	if err := modelParam.Add(); err != nil {
		ZapLog().With(zap.Error(err)).Error("database err")
		return
	}
	if errCode != 0 {
		log := orderErr.Error()
		if err := new(models.CardOrderLog).FtParseAdd(nil,&orderNo, &log).Add(); err != nil {
			ZapLog().With(zap.Error(err)).Error("database err")
		}
	}
}

func (this * Ydhk) recordOldOrder(oldOrder *models.CardOrder, errCode apibackend.EnumBasErr, oaoFlag bool, orderErr error) {
	newStatus := models.CONST_OrderStatus_Fail
	if errCode == 0 {
		if oaoFlag == true {
			newStatus =	models.CONST_OrderStatus_New
		}else{
			newStatus =	models.CONST_OrderStatus_New_UnFinish
		}
	}
	if (oldOrder.Status!=nil) && (*oldOrder.Status == models.CONST_OrderStatus_Fail) {
		if err := new(models.CardOrder).UpdateStatusByOrderNo(*oldOrder.OrderNo, newStatus); err != nil {
			ZapLog().With(zap.Error(err)).Error("database err")
		}
	}

	log := "订单申请成功"
	if errCode != 0 {
		log = orderErr.Error()
	}
	if err := new(models.CardOrderLog).FtParseAdd(nil,oldOrder.OrderNo, &log).Add(); err != nil {
		ZapLog().With(zap.Error(err)).Error("database err")
	}
}

func (this * Ydhk) OfflineActive(ctx iris.Context) {
	//isoao, _ := ctx.URLParamBool("isOao")
	//	productType := ctx.URLParamDefault("productType", "19")
	productId := ctx.URLParam("productId")
	channelId := ctx.URLParam("channelId")

	param := new(api.FtYdhkApply)
	if err := ctx.ReadJSON(param); err != nil {
		this.ExceptionSerive(ctx, apibackend.BASERR_INVALID_PARAMETER.Code(), apibackend.BASERR_INVALID_PARAMETER.Desc())
		ZapLog().Error("param err", zap.Error(err))
		return
	}

	oaoModel := new(string)
	*oaoModel = "2"
	errCode, orderId, oaoFlag, err := new(ReOrderSubmit).Parse(channelId, productId, oaoModel).OfflineActiveSend(param.AccessToken, param.Phone, param.NewPhone, param.LeagalName, param.CertificateNo, param.Address, param.Province, param.City, param.SendProvince, param.SendCity, param.SendDistrict)
	if err != nil {
		ZapLog().With(zap.Error(err)).Error("Retoken send err")
		this.ExceptionSerive(ctx, errCode.Code(), err.Error())
		return
	}
	this.Response(ctx, &api.FtResYdhkApply{orderId,oaoFlag})

	go func() {
		modelParam := &models.CardOrder{
			Status : new(int),
			ThirdOrderNo: &orderId,
		}

		if oaoFlag == true {
			*modelParam.Status =	models.CONST_OrderStatus_New
		}else{
			*modelParam.Status =	models.CONST_OrderStatus_New_UnFinish
		}

		condStr := "created_at >= " +fmt.Sprintf("%d", time.Now().Unix() - 12*3600)
		_, err := modelParam.UpdatesByNewPhoneAndIdcard(param.NewPhone, param.CertificateNo, condStr)
		if err != nil {
			ZapLog().With(zap.Error(err)).Error("database err")
			return
		}
	}()
}


func (this * Ydhk) UploadPhoto(ctx iris.Context) {
	//ll, err := new(models.CardClass).Gets()
	//if err != nil {
	//	ZapLog().With(zap.Error(err)).Error("Update err")
	//	this.ExceptionSerive(ctx, apibackend.BASERR_DATABASE_ERROR.Code(), apibackend.BASERR_DATABASE_ERROR.Desc())
	//	return
	//}
	//this.Response(ctx, ll)
}

func (this *Ydhk) FtConfirm(ctx iris.Context) {
	param := new(api.FtCardOrderConfirm)

	err := Tools.ShouldBindJSON(ctx, param)
	if err != nil {
		this.ExceptionSerive(ctx, apibackend.BASERR_INVALID_PARAMETER.Code(), apibackend.BASERR_INVALID_PARAMETER.Desc())
		ZapLog().Error("param err", zap.Error(err))
		return
	}
	if (param.SuccFlag != nil) && (*param.SuccFlag !=0) {
		err = new(models.CardOrder).UpdateStatusByOrderNo(param.OrderNo, models.CONST_OrderStatus_New)
		if err != nil {
			ZapLog().With(zap.Error(err)).Error("database err")
			this.ExceptionSerive(ctx, apibackend.BASERR_DATABASE_ERROR.Code(), apibackend.BASERR_DATABASE_ERROR.Desc())
			return
		}
	}


	if param.Log != nil && len(*param.Log) > 1 {
		go func (){
			new(models.CardOrderLog).FtParseAdd(nil, &param.OrderNo, param.Log).Add()
		}()
	}
	this.Response(ctx, nil)
}

func (this *Ydhk) FtIdCheckUrlGet(ctx iris.Context) {
	isOao,_ := ctx.URLParamBool("isOao")
	channelId := ctx.URLParam("channelId")
	param := new(api.FtIdCheckUrlGet)
	err := Tools.ShouldBindJSON(ctx, param)
	if err != nil {
		this.ExceptionSerive(ctx, apibackend.BASERR_INVALID_PARAMETER.Code(), apibackend.BASERR_INVALID_PARAMETER.Desc())
		ZapLog().Error("param err", zap.Error(err))
		return
	}

	url,err := new(ReIdCheckUrl).Send(isOao,channelId, param.ThirdOrderNo, param.NewPhone, param.Token)
	if err != nil {
		ZapLog().With(zap.Error(err)).Error("Retoken send err")
		this.ExceptionSerive(ctx, apibackend.BASERR_INTERNAL_SERVICE_ACCESS_ERROR.Code(), err.Error())
		return
	}
	this.Response(ctx, url)
	if param.OrderId == nil {
		return
	}
	go func(){
		err = new(models.CardOrderUrl).FtParseAdd(nil, param.OrderId, &url).Add()
		if err != nil {
			ZapLog().With(zap.Error(err)).Error("CardOrderUrl FtParseAdd err")
		}
	}()



}