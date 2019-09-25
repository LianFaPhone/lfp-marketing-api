package order

import "go.uber.org/zap"

import(
	. "LianFaPhone/lfp-marketing-api/controllers"
	"github.com/kataras/iris"
	"LianFaPhone/lfp-marketing-api/api"
	apibackend "LianFaPhone/lfp-api/errdef"
	. "LianFaPhone/lfp-base/log/zap"
	"LianFaPhone/lfp-marketing-api/models"
)


type CardDateSheet struct{
	Controllers
}

func (this * CardDateSheet) BkList(ctx iris.Context) {
	param := new(api.BkCardDateSheetList)

	err := Tools.ShouldBindJSON(ctx, param)
	if err != nil {
		this.ExceptionSerive(ctx, apibackend.BASERR_INVALID_PARAMETER.Code(), apibackend.BASERR_INVALID_PARAMETER.Desc())
		ZapLog().Error( "param err", zap.Error(err))
		return
	}

	results, err := new(models.CardDatesheet).ListWithConds(param.Page, param.Size, nil ,nil )
	if err != nil {
		ZapLog().With(zap.Error(err)).Error("CardDatesheet ListWithConds err")
		this.ExceptionSerive(ctx, apibackend.BASERR_DATABASE_ERROR.Code(), apibackend.BASERR_DATABASE_ERROR.Desc())
		return
	}
	this.Response(ctx, results)
}