package jthk

import (
	"LianFaPhone/lfp-marketing-api/api"
	"LianFaPhone/lfp-marketing-api/common"
	"LianFaPhone/lfp-marketing-api/config"
	"LianFaPhone/lfp-marketing-api/models"
	"LianFaPhone/lfp-marketing-api/thirdcard-api/douyin"
	"LianFaPhone/lfp-marketing-api/thirdcard-api/kuaishou"
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
	if err := Tools.ShouldBindJSON(ctx, param); err != nil {
		this.ExceptionSerive(ctx, apibackend.BASERR_INVALID_PARAMETER.Code(), apibackend.BASERR_INVALID_PARAMETER.Desc())
		ZapLog().Error("param err", zap.Error(err))
		return
	}
	//if len(param.CertificateNo) <=0 || len(param.Phone) <= 0 {
	//	this.ExceptionSerive(ctx, apibackend.BASERR_INVALID_PARAMETER.Code(), apibackend.BASERR_INVALID_PARAMETER.Desc())
	//	ZapLog().Error("param err")
	//	return
	//}

	adTp := int64(0)
	if (len(channelId) == 0) && (param.PartnerGoodsCode != nil) {
		ppg,err := new(models.PdPartnerGoods).GetByCodeFromCache(*param.PartnerGoodsCode)
		if err != nil {
			this.ExceptionSerive(ctx, apibackend.BASERR_DATABASE_ERROR.Code(), apibackend.BASERR_DATABASE_ERROR.Desc())
			ZapLog().Error("datebase err", zap.Error(err))
			return
		}
		if ppg == nil || (ppg.Valid!=nil && *ppg.Valid == 0){
			this.ExceptionSerive(ctx, apibackend.BASERR_OBJECT_NOT_FOUND.Code(), apibackend.BASERR_OBJECT_NOT_FOUND.Desc())
			ZapLog().Error("nofind ppg err")
			return
		}
		partner ,err := new(models.PdPartner).GetByIdFromCache(*ppg.PartnerId)
		if err != nil {
			ZapLog().With(zap.Error(err)).Error("Update err")
			this.ExceptionSerive(ctx, apibackend.BASERR_DATABASE_ERROR.Code(), apibackend.BASERR_DATABASE_ERROR.Desc())
			return
		}
		if partner == nil || (partner.Valid!=nil && *partner.Valid == 0) {
			ZapLog().Error("nofind ppg err")
			this.ExceptionSerive(ctx, apibackend.BASERR_OBJECT_NOT_FOUND.Code(), apibackend.BASERR_OBJECT_NOT_FOUND.Desc())
			return
		}
		urlParam := ""
		if partner.UrlParam != nil  {
			urlParam = *partner.UrlParam
		}

		if len(urlParam) <= 0 {
			urlParam = *ppg.UrlParam
		}else{
			urlParam = urlParam + "&"+*ppg.UrlParam
		}

		if len(urlParam) <= 0 {
			this.ExceptionSerive(ctx, apibackend.BASERR_OBJECT_DATA_NOT_FOUND.Code(), apibackend.BASERR_OBJECT_DATA_NOT_FOUND.Desc())
			ZapLog().Error("nofind urlparam err")
			return
		}
		vv, _ := url.ParseQuery(urlParam)
		channelId = vv.Get("channelId")
		productId = vv.Get("productId")

		isoao,_ = strconv.ParseBool(vv.Get("isOao"))
		if len(channelId) ==0 || len(productId) == 0 {
			this.ExceptionSerive(ctx, apibackend.BASERR_OBJECT_DATA_NOT_FOUND.Code(), apibackend.BASERR_OBJECT_DATA_NOT_FOUND.Desc())
			ZapLog().Error("nofind urlparam err")
			return
		}
		if ppg.AdTp != nil {
			adTp = int64(*ppg.AdTp)
		}
		if adTp <= 0  {
			adTp,_  = strconv.ParseInt(vv.Get("ad_tp"), 10, 32)
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
		var  urlValues url.Values
		if param.UrlQueryString != nil {
			urlValues,_ = url.ParseQuery(*param.UrlQueryString)
		}
		if adTp <= 0 {
			adTp = ctx.URLParamInt64Default("ad_tp", 0)
			if adTp <=0 {//这个好像 没必要
				adTp,_ = strconv.ParseInt(urlValues.Get("ad_tp"), 10, 32)
			}
		}

		orderNo := ""
		oldOrder,_ := new(models.CardOrder).GetByIdcardAndNewPhone(param.CertificateNo, param.NewPhone, &models.SqlPairCondition{"created_at > ?", time.Now().Unix() - 300})
		if oldOrder != nil { // 老订单
			this.recordOldOrder(oldOrder, errCode, oaoFlag, orderErr)
			orderNo = *oldOrder.OrderNo
		}else{
			orderNo = this.recordNewOrder(ctx, param, orderId, errCode, oaoFlag, orderErr, int(adTp))
		}

		if !oaoFlag ||(errCode != apibackend.BASERR_SUCCESS) || (param.UrlQueryString == nil) {
			//ZapLog().Error("kuaishou  noneed kuaishousend ", zap.String("orderNo", orderNo), zap.Bool("oao", isoao), zap.Int("code", errCode.Code()), zap.Any("param.UrlQueryString", param.UrlQueryString))
			return
		}
		//urlValues,err := url.ParseQuery(*param.UrlQueryString)
		//if err != nil {
		//	//ZapLog().Error("kuaishou ParseQuery err", zap.Error(err))
		//	return
		//}

		/**********************/
		//if adTp <= 0 {
		//	adTp = ctx.URLParamInt64Default("ad_tp", 0)
		//	if adTp <=0 {
		//		adTp,_ = strconv.ParseInt(urlValues.Get("ad_tp"), 10, 32)
		//	}
		//}
		/*************************/

		log:="成功"
		pushFlag :=1
		succFlag := 1
		if adTp == models.CONST_ADTRACK_Tp_KuaiShou {

			adCallBack := urlValues.Get("callback")
			if len(adCallBack) <= 0 {
				adCallBack = ctx.URLParam("callback")
				if len(adCallBack) <= 0 {
					//ZapLog().Error("kuaishou nofind callback", zap.Any("urlparam", param.UrlQueryString))
					return
				}
			}
			if err := new(kuaishou.ReTracker).Send(adCallBack, "9", time.Now().UnixNano()/1000); err != nil {
				ZapLog().Error("kuaishou send err", zap.Error(err), zap.String("callback", adCallBack))
				log = "失败："+err.Error()
				pushFlag =0
				succFlag = 0
			}

			if err := new(models.AdTrack).FtParseAdd(&orderNo, &adCallBack,&log,  int(adTp), &pushFlag, &succFlag).Add(); err != nil {
				ZapLog().Error("AdTrack err", zap.Error(err), zap.String("callback", adCallBack))
			}
		}else if adTp == models.CONST_ADTRACK_Tp_DouYin{
			link := ctx.GetHeader("Origin")
			link = link+"/?"+ *param.UrlQueryString
			if err := new(douyin.ReTracker).Send(link, "TD", time.Now().Unix(), 3); err != nil {
				ZapLog().Error("douyin send err", zap.Error(err), zap.String("link", link))
				log = "失败："+err.Error()
				pushFlag =0
				succFlag = 0
			}

			if err := new(models.AdTrack).FtParseAdd(&orderNo, &link,&log,  int(adTp), &pushFlag, &succFlag).Add(); err != nil {
				ZapLog().Error("AdTrack err", zap.Error(err), zap.String("link", link))
			}
		}else{
			ZapLog().Error("nofind ad_tp", zap.Any("urlparam", param.UrlQueryString), zap.Int64("ad_tp", adTp))
		}


	}()

}

func (this * Ydhk) recordNewOrder(ctx iris.Context, param *api.FtYdhkApply, thirdOrderId string, errCode apibackend.EnumBasErr, oaoFlag bool, orderErr error, adTp int) string {
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

	//if adTp > 0 {
		modelParam.AdTp = &adTp
	//}


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
		*modelParam.Status = this.ParseFailStatus(orderErr.Error())
	}else{
		if oaoFlag == true {
			*modelParam.Status =	models.CONST_OrderStatus_New
		}else{
			*modelParam.Status =	models.CONST_OrderStatus_New_UnFinish
		}
	}

	if err := modelParam.Add(); err != nil {
		ZapLog().With(zap.Error(err)).Error("database err")
		return orderNo
	}
	if errCode != 0 {
		log := orderErr.Error()
		if err := new(models.CardOrderLog).FtParseAdd(nil,&orderNo, &log).Add(); err != nil {
			ZapLog().With(zap.Error(err)).Error("database err")
		}
	}
	return orderNo
}

func (this * Ydhk) recordOldOrder(oldOrder *models.CardOrder, errCode apibackend.EnumBasErr, oaoFlag bool, orderErr error) {
	newStatus := models.CONST_OrderStatus_Fail
	if errCode == 0 {
		if oaoFlag == true {
			newStatus =	models.CONST_OrderStatus_New
		}else{
			newStatus =	models.CONST_OrderStatus_New_UnFinish
		}
	}else{
		newStatus = this.ParseFailStatus(orderErr.Error())
	}
	if (oldOrder.Status!=nil) && (*oldOrder.Status >= models.CONST_OrderStatus_MinFail || *oldOrder.Status <= models.CONST_OrderStatus_MaxFail) {
		if err := new(models.CardOrder).UpdateStatusByOrderNo(*oldOrder.OrderNo, newStatus); err != nil {
			ZapLog().With(zap.Error(err)).Error("database err")
			log := "老订单状态更新失败"
			new(models.CardOrderLog).FtParseAdd(nil,oldOrder.OrderNo, &log).Add()
		}
	}

	log := "订单再次申请成功"
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
	go func(){
		if param.OrderId == nil {
			cardOrder,err := new(models.CardOrder).GetByThirdOrderNo(param.ThirdOrderNo)
			if err != nil {
				ZapLog().With(zap.Error(err)).Error("CardOrderUrl FtParseAdd err")
				return
			}
			if cardOrder == nil {
				return
			}
			param.OrderId = cardOrder.OrderNo
		}
		if param.OrderId == nil {
			return
		}
		err = new(models.CardOrderUrl).FtParseAdd(nil, param.OrderId, &url).Add()
		if err != nil {
			ZapLog().With(zap.Error(err)).Error("CardOrderUrl FtParseAdd err")
			log := "上传照片网址存储失败："+err.Error()
			new(models.CardOrderLog).FtParseAdd(nil, param.OrderId, &log).Add()

		}
	}()



}

func (this *Ydhk) ParseFailStatus(errMsg string) int {
	if strings.Contains(errMsg, "欠费号码") {
		return models.CONST_OrderStatus_Fail_Retry
	}else if strings.Contains(errMsg, "已超时") {
		return models.CONST_OrderStatus_Fail_Retry
	}else if strings.Contains(errMsg, "系统错误") {
		return models.CONST_OrderStatus_Fail_Retry
	}else if strings.Contains(errMsg, "订购的号码不存在") {
		return models.CONST_OrderStatus_Fail_Retry
	}else if strings.Contains(errMsg, "号码已被占用") {
		return models.CONST_OrderStatus_Fail_Retry
	}else if strings.Contains(errMsg, "系统忙") {
		return models.CONST_OrderStatus_Fail_Retry
	}else if strings.Contains(errMsg, "无法申请新号卡") {
		return models.CONST_OrderStatus_Fail_Retry
	}else if strings.Contains(errMsg, "锁定号码失败") {
		return models.CONST_OrderStatus_Fail_Retry
	}
	return models.CONST_OrderStatus_Fail
}