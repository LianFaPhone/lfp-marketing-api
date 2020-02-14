package order

import (
	apibackend "LianFaPhone/lfp-api/errdef"
	. "LianFaPhone/lfp-base/log/zap"
	"LianFaPhone/lfp-marketing-api/api"

	. "LianFaPhone/lfp-marketing-api/controllers"
	"LianFaPhone/lfp-marketing-api/models"

	"github.com/kataras/iris"
	"go.uber.org/zap"
)

type CardOrderLog struct{
	Controllers
}

func (this *CardOrderLog) BkList(ctx iris.Context) {
	param := new(api.BkCardOrderLogList)

	err := Tools.ShouldBindJSON(ctx, param)
	if err != nil {
		this.ExceptionSerive(ctx, apibackend.BASERR_INVALID_PARAMETER.Code(), apibackend.BASERR_INVALID_PARAMETER.Desc())
		ZapLog().Error("param err", zap.Error(err))
		return
	}

	ll,err :=new(models.CardOrderLog).BkParseList(param).ListWithConds(param.Page, param.Size, nil ,nil)
	if err != nil {
		ZapLog().With(zap.Error(err)).Error("database err")
		this.ExceptionSerive(ctx, apibackend.BASERR_DATABASE_ERROR.Code(), apibackend.BASERR_DATABASE_ERROR.Desc())
		return
	}
	this.Response(ctx, ll)
}
