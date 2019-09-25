package controllers

import (
	"github.com/kataras/iris"
	"LianFaPhone/lfp-marketing-api/api"
	"go.uber.org/zap"
	apibackend "LianFaPhone/lfp-api/errdef"
	. "LianFaPhone/lfp-base/log/zap"
	"LianFaPhone/lfp-marketing-api/models"
)

type PhoneNumberPool struct{
	Controllers
}

func (this *PhoneNumberPool) FtList(ctx iris.Context) {
	param := new(api.FtPhoneNumberList)

	err := Tools.ShouldBindJSON(ctx, param)
	if err != nil {
		this.ExceptionSerive(ctx, apibackend.BASERR_INVALID_PARAMETER.Code(), apibackend.BASERR_INVALID_PARAMETER.Desc())
		ZapLog().Error( "param err", zap.Error(err))
		return
	}
	param.UseFlag = 0
	param.Valid = 1

	ll,err := new(models.PhoneNumberPool).FtParseList(param).ListWithConds(param.Page, param.Size, []string{"number", "level"}, nil)
	if err != nil{
		ZapLog().With(zap.Error(err)).Error("Update err")
		this.ExceptionSerive(ctx, apibackend.BASERR_DATABASE_ERROR.Code(), apibackend.BASERR_DATABASE_ERROR.Desc())
		return
	}
	this.Response(ctx, ll)
}

func (this *PhoneNumberPool) BkList(ctx iris.Context) {
	param := new(api.BkPhoneNumberList)

	err := Tools.ShouldBindJSON(ctx, param)
	if err != nil {
		this.ExceptionSerive(ctx, apibackend.BASERR_INVALID_PARAMETER.Code(), apibackend.BASERR_INVALID_PARAMETER.Desc())
		ZapLog().Error( "param err", zap.Error(err))
		return
	}

	condPair := make([]*models.SqlPairCondition, 0, 2)
	if param.StartCreatedAt != nil {
		condPair= append(condPair, &models.SqlPairCondition{"created_at >= ?", param.StartCreatedAt})
	}
	if param.EndCreatedAt != nil {
		condPair= append(condPair, &models.SqlPairCondition{"created_at <= ?", param.EndCreatedAt})
	}

	ll,err := new(models.PhoneNumberPool).BkParseList(param).ListWithConds(param.Page, param.Size, nil, condPair)
	if err != nil{
		ZapLog().With(zap.Error(err)).Error("Update err")
		this.ExceptionSerive(ctx, apibackend.BASERR_DATABASE_ERROR.Code(), apibackend.BASERR_DATABASE_ERROR.Desc())
		return
	}
	this.Response(ctx, ll)
}

func (this *PhoneNumberPool) Get(ctx iris.Context) {
	param := new(api.BkPhoneNumberGet)

	err := Tools.ShouldBindJSON(ctx, param)
	if err != nil {
		this.ExceptionSerive(ctx, apibackend.BASERR_INVALID_PARAMETER.Code(), apibackend.BASERR_INVALID_PARAMETER.Desc())
		ZapLog().Error( "param err", zap.Error(err))
		return
	}

	if (param.Id == nil) && (param.Number == nil) {
		this.ExceptionSerive(ctx, apibackend.BASERR_INVALID_PARAMETER.Code(), apibackend.BASERR_INVALID_PARAMETER.Desc())
		ZapLog().Error( "param err", zap.Any("param", *param))
		return
	}

	res ,err := new(models.PhoneNumberPool).ParseGet(param).Get()
	if err != nil {
		ZapLog().With(zap.Error(err)).Error("Update err")
		this.ExceptionSerive(ctx, apibackend.BASERR_DATABASE_ERROR.Code(), apibackend.BASERR_DATABASE_ERROR.Desc())
		return
	}
	this.Response(ctx, res)
}

func (this *PhoneNumberPool) NumberCheck(ctx iris.Context) {
	param := new(api.FtPhoneNumberCheck)

	err := Tools.ShouldBindJSON(ctx, param)
	if err != nil {
		this.ExceptionSerive(ctx, apibackend.BASERR_INVALID_PARAMETER.Code(), apibackend.BASERR_INVALID_PARAMETER.Desc())
		ZapLog().Error( "param err", zap.Error(err))
		return
	}

	res ,err := new(models.PhoneNumberPool).GetByNumber(*param.Number)
	if err != nil {
		ZapLog().With(zap.Error(err)).Error("Update err")
		this.ExceptionSerive(ctx, apibackend.BASERR_DATABASE_ERROR.Code(), apibackend.BASERR_DATABASE_ERROR.Desc())
		return
	}
	if res == nil {
		this.ExceptionSerive(ctx, apibackend.BASERR_OBJECT_NOT_FOUND.Code(), apibackend.BASERR_OBJECT_NOT_FOUND.Desc())
		return

	}
	if res.UseFlag != nil && *res.UseFlag == 1 {
		this.ExceptionSerive(ctx, apibackend.BASERR_OBJECT_BE_USED.Code(), apibackend.BASERR_OBJECT_BE_USED.Desc())
		return
	}
	this.Response(ctx, nil)
}

func (this *PhoneNumberPool) Update(ctx iris.Context) {
	param := new(api.BkPhoneNumber)

	err := Tools.ShouldBindJSON(ctx, param)
	if err != nil {
		this.ExceptionSerive(ctx, apibackend.BASERR_INVALID_PARAMETER.Code(), apibackend.BASERR_INVALID_PARAMETER.Desc())
		ZapLog().Error( "param err", zap.Error(err))
		return
	}

	err = new(models.PhoneNumberPool).Parse(param).Update()
	if err != nil {
		ZapLog().With(zap.Error(err)).Error("Update err")
		this.ExceptionSerive(ctx, apibackend.BASERR_DATABASE_ERROR.Code(), apibackend.BASERR_DATABASE_ERROR.Desc())
		return
	}
	this.Response(ctx, nil)
}

func (this *PhoneNumberPool) Adds(ctx iris.Context) {
	params := make([]*api.BkPhoneNumberSave, 0)

	if err := ctx.ReadJSON(&params); err != nil {
		this.ExceptionSerive(ctx, apibackend.BASERR_INVALID_PARAMETER.Code(), apibackend.BASERR_INVALID_PARAMETER.Desc())
		ZapLog().Error( "param err", zap.Error(err))
		return
	}


	res := &api.BkResPhoneNumberSave{
		SuccNumber: make([]*string, 0),
		FailNumber: make([]*string, 0),
	}

	for i:=0; i < len(params); i++ {
		uniqueFlag,err := new(models.PhoneNumberPool).UniqueByNumber(*params[i].Number)
		if err != nil {
			ZapLog().With(zap.Error(err)).Error("UniqueByNumber err")
			//this.ExceptionSerive(ctx, apibackend.BASERR_DATABASE_ERROR.Code(), apibackend.BASERR_DATABASE_ERROR.Desc())
			res.FailCount = len(res.FailNumber)
			res.SuccCount = len(params) - res.FailCount
			this.ExceptionSeriveWithData(ctx, apibackend.BASERR_DATABASE_ERROR.Code(), res)
			return
		}
		if !uniqueFlag {
			res.FailNumber = append(res.FailNumber, params[i].Number)
			continue
		}
		err = new(models.PhoneNumberPool).ParseAdd(*params[i].Number, *params[i].Level).Add()
		if err != nil {
			//ZapLog().With(zap.Error(err), zap.String("number", *params[i].Number), zap.Int("level", *params[i].Level)).Error("Number Add err")
			//this.ExceptionSerive(ctx, apibackend.BASERR_DATABASE_ERROR.Code(), apibackend.BASERR_DATABASE_ERROR.Desc())
			res.FailCount = len(res.FailNumber)
			res.SuccCount = len(params) - res.FailCount
			this.ExceptionSeriveWithData(ctx, apibackend.BASERR_DATABASE_ERROR.Code(), res)
			return
		}
		res.SuccNumber = append(res.SuccNumber, params[i].Number)
	}

	res.FailCount = len(res.FailNumber)
	res.SuccCount = len(params) - res.FailCount
	this.ExceptionSeriveWithData(ctx, apibackend.BASERR_SUCCESS.Code(), res)
}
