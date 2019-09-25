package black

import (
	. "LianFaPhone/lfp-marketing-api/controllers"
"github.com/kataras/iris"
"LianFaPhone/lfp-marketing-api/api"
"go.uber.org/zap"
	apibackend "LianFaPhone/lfp-api/errdef"
. "LianFaPhone/lfp-base/log/zap"
"LianFaPhone/lfp-marketing-api/models"
)

type BlacklistPhone struct{
	Controllers
}

func (this *BlacklistPhone) List(ctx iris.Context) {
	param := new(api.BkBlacklistPhoneList)

	err := Tools.ShouldBindJSON(ctx, param)
	if err != nil {
		this.ExceptionSerive(ctx, apibackend.BASERR_INVALID_PARAMETER.Code(), apibackend.BASERR_INVALID_PARAMETER.Desc())
		ZapLog().Error( "param err", zap.Error(err))
		return
	}

	//condPair := make([]*models.SqlPairCondition, 0, 2)
	//if param.StartCreatedAt != nil {
	//	condPair= append(condPair, &models.SqlPairCondition{"created_at >= ?", param.StartCreatedAt})
	//}
	//if param.EndCreatedAt != nil {
	//	condPair= append(condPair, &models.SqlPairCondition{"created_at <= ?", param.EndCreatedAt})
	//}

	ll,err := new(models.BlacklistPhone).ParseList(param).ListWithConds(param.Page, param.Size, nil, nil)
	if err != nil{
		ZapLog().With(zap.Error(err)).Error("Update err")
		this.ExceptionSerive(ctx, apibackend.BASERR_DATABASE_ERROR.Code(), apibackend.BASERR_DATABASE_ERROR.Desc())
		return
	}
	this.Response(ctx, ll)
}

func (this *BlacklistPhone) Update(ctx iris.Context) {
	param := new(api.BkBlacklistPhone)

	err := Tools.ShouldBindJSON(ctx, param)
	if err != nil {
		this.ExceptionSerive(ctx, apibackend.BASERR_INVALID_PARAMETER.Code(), apibackend.BASERR_INVALID_PARAMETER.Desc())
		ZapLog().Error( "param err", zap.Error(err))
		return
	}

	err = new(models.BlacklistPhone).Parse(param).Update()
	if err != nil {
		ZapLog().With(zap.Error(err)).Error("Update err")
		this.ExceptionSerive(ctx, apibackend.BASERR_DATABASE_ERROR.Code(), apibackend.BASERR_DATABASE_ERROR.Desc())
		return
	}
	this.Response(ctx, nil)
}

func (this *BlacklistPhone) Add(ctx iris.Context) {
	param := new(api.BkBlacklistPhoneAdd)

	err := Tools.ShouldBindJSON(ctx, param)
	if err != nil {
		this.ExceptionSerive(ctx, apibackend.BASERR_INVALID_PARAMETER.Code(), apibackend.BASERR_INVALID_PARAMETER.Desc())
		ZapLog().Error( "param err", zap.Error(err))
		return
	}

	err = new(models.BlacklistPhone).ParseAdd(param).Add()
	if err != nil {
		ZapLog().With(zap.Error(err)).Error("Update err")
		this.ExceptionSerive(ctx, apibackend.BASERR_DATABASE_ERROR.Code(), apibackend.BASERR_DATABASE_ERROR.Desc())
		return
	}
	this.Response(ctx, nil)
}

