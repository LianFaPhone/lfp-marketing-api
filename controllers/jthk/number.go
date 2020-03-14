package jthk

import (
	"LianFaPhone/lfp-marketing-api/api"
	. "LianFaPhone/lfp-marketing-api/thirdcard-api/ydjthk"
	"go.uber.org/zap"
	"github.com/kataras/iris"
	apibackend "LianFaPhone/lfp-api/errdef"
	. "LianFaPhone/lfp-base/log/zap"
)

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
