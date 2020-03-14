package jthk

import (
	"LianFaPhone/lfp-marketing-api/api"
	"LianFaPhone/lfp-marketing-api/models"
	. "LianFaPhone/lfp-marketing-api/thirdcard-api/ydjthk"
	"go.uber.org/zap"
	"net/url"
	"strconv"
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
		ZapLog().With(zap.Error(err)).Error("Protocal send err")
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
		ZapLog().With(zap.Error(err)).Error("ReToken send err")
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
		ZapLog().With(zap.Error(err)).Error("GetAddr err")
		this.ExceptionSerive(ctx, apibackend.BASERR_INTERNAL_SERVICE_ACCESS_ERROR.Code(), err.Error())
		return
	}
	this.Response(ctx, ll)
}