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

type BlacklistIdCard struct{
	Controllers
}

func (this *BlacklistIdCard) List(ctx iris.Context) {
	param := new(api.BkBlacklistIdcardList)

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

	ll,err := new(models.BlacklistIdcard).ParseList(param).ListWithConds(param.Page, param.Size, nil, nil)
	if err != nil{
		ZapLog().With(zap.Error(err)).Error("Update err")
		this.ExceptionSerive(ctx, apibackend.BASERR_DATABASE_ERROR.Code(), apibackend.BASERR_DATABASE_ERROR.Desc())
		return
	}
	this.Response(ctx, ll)
}

func (this *BlacklistIdCard) Update(ctx iris.Context) {
	param := new(api.BkBlacklistIdcard)

	err := Tools.ShouldBindJSON(ctx, param)
	if err != nil {
		this.ExceptionSerive(ctx, apibackend.BASERR_INVALID_PARAMETER.Code(), apibackend.BASERR_INVALID_PARAMETER.Desc())
		ZapLog().Error( "param err", zap.Error(err))
		return
	}

	err = new(models.BlacklistIdcard).Parse(param).Update()
	if err != nil {
		ZapLog().With(zap.Error(err)).Error("Update err")
		this.ExceptionSerive(ctx, apibackend.BASERR_DATABASE_ERROR.Code(), apibackend.BASERR_DATABASE_ERROR.Desc())
		return
	}
	this.Response(ctx, nil)
}

func (this *BlacklistIdCard) Add(ctx iris.Context) {
	param := new(api.BkBlacklistIdcardAdd)

	err := Tools.ShouldBindJSON(ctx, param)
	if err != nil {
		this.ExceptionSerive(ctx, apibackend.BASERR_INVALID_PARAMETER.Code(), apibackend.BASERR_INVALID_PARAMETER.Desc())
		ZapLog().Error( "param err", zap.Error(err))
		return
	}

	err = new(models.BlacklistIdcard).ParseAdd(param).Add()
	if err != nil {
		ZapLog().With(zap.Error(err)).Error("Update err")
		this.ExceptionSerive(ctx, apibackend.BASERR_DATABASE_ERROR.Code(), apibackend.BASERR_DATABASE_ERROR.Desc())
		return
	}
	this.Response(ctx, nil)
}

