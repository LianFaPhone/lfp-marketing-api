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


type CardAreaSheet struct{
	Controllers
}

func (this * CardAreaSheet) BkList(ctx iris.Context) {
	param := new(api.BkCardAreaSheetList)

	err := Tools.ShouldBindJSON(ctx, param)
	if err != nil {
		this.ExceptionSerive(ctx, apibackend.BASERR_INVALID_PARAMETER.Code(), apibackend.BASERR_INVALID_PARAMETER.Desc())
		ZapLog().Error( "param err", zap.Error(err))
		return
	}

	cond := make([]*models.SqlPairCondition, 0)
	if param.LikeStr != nil {
		cond = append(cond, &models.SqlPairCondition{"city like ?", "%"+*param.LikeStr+"%"})
	}

	//results, err := new(models.CardOrder).BkParseList(param).ListAreaCountWithConds(param.Page, param.Size, cond )
	//if err != nil {
	//	ZapLog().With(zap.Error(err)).Error("Verify err")
	//	this.ExceptionSerive(ctx, apibackend.BASERR_DATABASE_ERROR.Code(), apibackend.BASERR_DATABASE_ERROR.Desc())
	//	return
	//}
	results, err := new(models.CardAreasheet).ParseList(param).ListWithConds(param.Page, param.Size, nil, cond)
	if err != nil {
		ZapLog().With(zap.Error(err)).Error("Verify err")
		this.ExceptionSerive(ctx, apibackend.BASERR_DATABASE_ERROR.Code(), apibackend.BASERR_DATABASE_ERROR.Desc())
		return
	}
	this.Response(ctx, results)
}