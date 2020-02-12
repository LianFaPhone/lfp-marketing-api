package controllers

import (
	apibackend "LianFaPhone/lfp-api/errdef"
	. "LianFaPhone/lfp-base/log/zap"
	"LianFaPhone/lfp-marketing-api/api"
	"LianFaPhone/lfp-marketing-api/models"
	"github.com/kataras/iris"
	"go.uber.org/zap"
)

type PhoneNumberLevel struct {
	Controllers
}

func (this *PhoneNumberLevel) FtGets(ctx iris.Context) {
	param := new(api.BkPhoneNumberLevelGets)

	err := Tools.ShouldBindJSON(ctx, param)
	if err != nil {
		this.ExceptionSerive(ctx, apibackend.BASERR_INVALID_PARAMETER.Code(), apibackend.BASERR_INVALID_PARAMETER.Desc())
		ZapLog().Error("param err", zap.Error(err))
		return
	}

	if param.Valid == nil {
		param.Valid = new(int)
		*param.Valid = 1
	}

	ll, err := new(models.PhoneNumberLevel).ParseGets(param).Gets()
	if err != nil {
		ZapLog().With(zap.Error(err)).Error("Update err")
		this.ExceptionSerive(ctx, apibackend.BASERR_DATABASE_ERROR.Code(), apibackend.BASERR_DATABASE_ERROR.Desc())
		return
	}
	this.Response(ctx, ll)
}

func (this *PhoneNumberLevel) Update(ctx iris.Context) {
	param := new(api.BkPhoneNumberLevelUpdate)

	err := Tools.ShouldBindJSON(ctx, param)
	if err != nil {
		this.ExceptionSerive(ctx, apibackend.BASERR_INVALID_PARAMETER.Code(), apibackend.BASERR_INVALID_PARAMETER.Desc())
		ZapLog().Error("param err", zap.Error(err))
		return
	}

	err = new(models.PhoneNumberLevel).ParseUpdate(param).Update()
	if err != nil {
		ZapLog().With(zap.Error(err)).Error("Update err")
		this.ExceptionSerive(ctx, apibackend.BASERR_DATABASE_ERROR.Code(), apibackend.BASERR_DATABASE_ERROR.Desc())
		return
	}
	this.Response(ctx, nil)
}

func (this *PhoneNumberLevel) Add(ctx iris.Context) {
	param := new(api.BkPhoneNumberLevelAdd)

	err := Tools.ShouldBindJSON(ctx, param)
	if err != nil {
		this.ExceptionSerive(ctx, apibackend.BASERR_INVALID_PARAMETER.Code(), apibackend.BASERR_INVALID_PARAMETER.Desc())
		ZapLog().Error("param err", zap.Error(err))
		return
	}

	uFlag, err := new(models.PhoneNumberLevel).UniqueByLevel(*param.Level)
	if err != nil {
		ZapLog().With(zap.Error(err)).Error("UniqueByLevel err")
		this.ExceptionSerive(ctx, apibackend.BASERR_DATABASE_ERROR.Code(), apibackend.BASERR_DATABASE_ERROR.Desc())
		return
	}
	if !uFlag {
		this.ExceptionSerive(ctx, apibackend.BASERR_OBJECT_EXISTS.Code(), apibackend.BASERR_OBJECT_EXISTS.Desc())
		return
	}

	err = new(models.PhoneNumberLevel).ParseAdd(param).Add()
	if err != nil {
		ZapLog().With(zap.Error(err)).Error("Update err")
		this.ExceptionSerive(ctx, apibackend.BASERR_DATABASE_ERROR.Code(), apibackend.BASERR_DATABASE_ERROR.Desc())
		return
	}
	this.Response(ctx, nil)
}

func (this *PhoneNumberLevel) List(ctx iris.Context) {
	param := new(api.BkPhoneNumberLevelList)

	err := Tools.ShouldBindJSON(ctx, param)
	if err != nil {
		this.ExceptionSerive(ctx, apibackend.BASERR_INVALID_PARAMETER.Code(), apibackend.BASERR_INVALID_PARAMETER.Desc())
		ZapLog().Error("param err", zap.Error(err))
		return
	}

	ll, err := new(models.PhoneNumberLevel).ParseList(param).ListWithConds(param.Page, param.Size, nil, nil)
	if err != nil {
		ZapLog().With(zap.Error(err)).Error("Update err")
		this.ExceptionSerive(ctx, apibackend.BASERR_DATABASE_ERROR.Code(), apibackend.BASERR_DATABASE_ERROR.Desc())
		return
	}
	this.Response(ctx, ll)
}
