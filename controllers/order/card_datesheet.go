package order

import "go.uber.org/zap"

import (
	apibackend "LianFaPhone/lfp-api/errdef"
	. "LianFaPhone/lfp-base/log/zap"
	"LianFaPhone/lfp-marketing-api/api"
	. "LianFaPhone/lfp-marketing-api/controllers"
	"LianFaPhone/lfp-marketing-api/models"
	"github.com/kataras/iris"
)

type CardDateSheet struct {
	Controllers
}

func (this *CardDateSheet) BkList(ctx iris.Context) {
	param := new(api.BkCardDateSheetList)

	err := Tools.ShouldBindJSON(ctx, param)
	if err != nil {
		this.ExceptionSerive(ctx, apibackend.BASERR_INVALID_PARAMETER.Code(), apibackend.BASERR_INVALID_PARAMETER.Desc())
		ZapLog().Error("param err", zap.Error(err))
		return
	}

	results, err := new(models.CardDatesheet).ListWithConds(param.Page, param.Size, nil, nil)
	if err != nil {
		ZapLog().With(zap.Error(err)).Error("CardDatesheet ListWithConds err")
		this.ExceptionSerive(ctx, apibackend.BASERR_DATABASE_ERROR.Code(), apibackend.BASERR_DATABASE_ERROR.Desc())
		return
	}
	arr, ok := results.List.(*[]*models.CardDatesheet)
	if ok {
		for i := 0; i < len(*arr); i++ {
			if (*arr)[i].ClassTp != nil {
				ss, ok := models.ClassTpMap[*(*arr)[i].ClassTp]
				if ok {
					(*arr)[i].ClassName = &ss.Name
				}
			}
		}
	}
	this.Response(ctx, results)
}
