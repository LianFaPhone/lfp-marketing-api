package controllers

import (
	apibackend "LianFaPhone/lfp-api/errdef"
	. "LianFaPhone/lfp-base/log/zap"
	"LianFaPhone/lfp-marketing-api/api"
	"LianFaPhone/lfp-marketing-api/common"
	"LianFaPhone/lfp-marketing-api/config"
	"LianFaPhone/lfp-marketing-api/db"
	"LianFaPhone/lfp-marketing-api/models"
	"github.com/kataras/iris"
	"go.uber.org/zap"
)

var phoneLimiter *common.BusLimiter
var ipLimiter *common.BusLimiter

type Sms struct {
	Controllers
}

func (this *Sms) Send(ctx iris.Context) {
	param := new(api.SmsSend)
	err := Tools.ShouldBindJSON(ctx, param)
	if err != nil {
		this.ExceptionSerive(ctx, apibackend.BASERR_INVALID_PARAMETER.Code(), apibackend.BASERR_INVALID_PARAMETER.Desc())
		ZapLog().Error("param err", zap.Error(err))
		return
	}

	if phoneLimiter == nil || ipLimiter == nil {
		this.ExceptionSerive(ctx, apibackend.BASERR_UNKNOWN_BUG.Code(), apibackend.BASERR_UNKNOWN_BUG.Desc())
		ZapLog().Error("just bug err")
		return
	}
	if param.CountryCode == nil || len(*param.CountryCode) == 0 {
		param.CountryCode = new(string)
		*param.CountryCode = "0086"
	}
	if param.Language == nil {
		param.Language = new(string)
		*param.Language = "zh-CN"
	}

	//业务频率限制下
	recipient := *param.CountryCode + *param.Phone
	limitFlag, err := phoneLimiter.Check(recipient)
	if err != nil {
		ZapLog().With(zap.Error(err)).Error("redis err")
		this.ExceptionSerive(ctx, apibackend.BASERR_DATABASE_ERROR.Code(), apibackend.BASERR_DATABASE_ERROR.Desc())
		return
	}
	if limitFlag {
		ZapLog().With(zap.Error(err)).Error("phone sms limit err")
		this.ExceptionSerive(ctx, apibackend.BASERR_OPERATE_FREQUENT.Code(), apibackend.BASERR_OPERATE_FREQUENT.Desc())
		return
	}

	limitFlag, err = ipLimiter.Check(common.GetRealIp(ctx))
	if err != nil {
		ZapLog().With(zap.Error(err)).Error("redis err")
		this.ExceptionSerive(ctx, apibackend.BASERR_DATABASE_ERROR.Code(), apibackend.BASERR_DATABASE_ERROR.Desc())
		return
	}
	if limitFlag {
		ZapLog().With(zap.Error(err)).Error("ip sms limit err")
		this.ExceptionSerive(ctx, apibackend.BASERR_OPERATE_FREQUENT.Code(), apibackend.BASERR_OPERATE_FREQUENT.Desc())
		return
	}

	verification := common.NewVerification(&db.GRedis, "sim", "")
	verifyId, err := verification.GenerateSms(0, recipient, config.GConfig.BasNotify.VerifyCodeSmsTmp, *param.Language, param.PlayTp)
	if err != nil {
		ZapLog().With(zap.Error(err)).Error("GenerateSms err")
		this.ExceptionSerive(ctx, apibackend.BASERR_DATABASE_ERROR.Code(), apibackend.BASERR_DATABASE_ERROR.Desc())
		return
	}

	res := models.Sms{
		VerifyUUId: &verifyId,
	}
	this.Response(ctx, res)
}

func (this *Sms) Verify(ctx iris.Context) {
	param := new(api.SmsVerify)
	err := Tools.ShouldBindJSON(ctx, param)
	if err != nil {
		this.ExceptionSerive(ctx, apibackend.BASERR_INVALID_PARAMETER.Code(), apibackend.BASERR_INVALID_PARAMETER.Desc())
		ZapLog().Error("param err", zap.Error(err))
		return
	}

	if param.CountryCode == nil || len(*param.CountryCode) == 0 {
		param.CountryCode = new(string)
		*param.CountryCode = "0086"
	}

	recipient := *param.CountryCode + *param.Phone
	verification := common.NewVerification(&db.GRedis, "sim", "")
	flag, err := verification.Verify(*param.VerifyUUId, 0, *param.VerifyCode, recipient)
	if err != nil {
		ZapLog().With(zap.Error(err)).Error("Verify err")
		this.ExceptionSerive(ctx, apibackend.BASERR_DATABASE_ERROR.Code(), apibackend.BASERR_DATABASE_ERROR.Desc())
		return
	}
	if !flag {
		ZapLog().With(zap.Error(err)).Error("Verify err")
		this.ExceptionSerive(ctx, apibackend.BASERR_ACTIVITY_FISSIONSHARE_SMS_INCORRECT_VERIFYCODE.Code(), apibackend.BASERR_ACTIVITY_FISSIONSHARE_SMS_INCORRECT_VERIFYCODE.OriginDesc())
		return
	}
	this.Response(ctx, nil)
}
