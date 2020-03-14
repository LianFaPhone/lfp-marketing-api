package jthk

import (
	apibackend "LianFaPhone/lfp-api/errdef"
	. "LianFaPhone/lfp-base/log/zap"
	. "LianFaPhone/lfp-marketing-api/controllers"
	"github.com/kataras/iris"
	"LianFaPhone/lfp-marketing-api/api"
	"LianFaPhone/lfp-marketing-api/models"
	. "LianFaPhone/lfp-marketing-api/thirdcard-api/ydjthk"
	"go.uber.org/zap"
)

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
		ZapLog().With(zap.Error(err)).Error("ReIdCheckUrl send err")
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
			log := "照片上传网址|获取成功|存储失败："+err.Error()
			new(models.CardOrderLog).FtParseAdd(nil, param.OrderId, &log).Add()
			return
		}
		log := "照片上传网址|获取成功"
		new(models.CardOrderLog).FtParseAdd(nil, param.OrderId, &log).Add()
	}()



}
