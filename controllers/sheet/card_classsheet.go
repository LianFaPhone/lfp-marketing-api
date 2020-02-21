package sheet

import "go.uber.org/zap"

import (
	apibackend "LianFaPhone/lfp-api/errdef"
	. "LianFaPhone/lfp-base/log/zap"
	"LianFaPhone/lfp-marketing-api/api"
	. "LianFaPhone/lfp-marketing-api/controllers"
	"LianFaPhone/lfp-marketing-api/models"
	"github.com/kataras/iris"
)

type CardClassSheet struct {
	Controllers
}

func (this *CardClassSheet) BkList(ctx iris.Context) {
	param := new(api.BkCardClassSheetList)

	err := Tools.ShouldBindJSON(ctx, param)
	if err != nil {
		this.ExceptionSerive(ctx, apibackend.BASERR_INVALID_PARAMETER.Code(), apibackend.BASERR_INVALID_PARAMETER.Desc())
		ZapLog().Error("param err", zap.Error(err))
		return
	}
	conds := make([]*models.SqlPairCondition, 0,3)

	if param.StartCreatedAt != nil {
		conds = append(conds, &models.SqlPairCondition{"start_created_at >= ?", param.StartCreatedAt})
	}
	if param.EndCreatedAt != nil {
		conds = append(conds, &models.SqlPairCondition{"end_created_at <= ?", param.EndCreatedAt})
	}
	selectFileds := []string{"date"}

	groupStr := "date"
	groupFlag := false
	if (param.GroupIspFlag != nil) && (*param.GroupIspFlag == 1){
		selectFileds = append(selectFileds, "isp")
		groupStr += ",isp"
		groupFlag = true
	}
	if (param.GroupClassBigTpFlag != nil)&&(*param.GroupClassBigTpFlag == 1) {
		selectFileds = append(selectFileds, "class_big_tp")
		groupStr += ",class_big_tp"
		groupFlag = true
	}
	if (param.GroupClassTpFlag != nil)&&(*param.GroupClassTpFlag == 1) {
		selectFileds = append(selectFileds, "class_tp")
		groupStr += ",class_tp"
		groupFlag = true
	}

	if groupFlag {
		selectFileds = append(selectFileds, "sum(order_count)")
	} else {
		selectFileds = append(selectFileds, "order_count")
		groupStr = ""
	}

	results, err := new(models.CardClasssheet).ParseList(param).ListWithConds(param.Page, param.Size, selectFileds, conds, groupStr)
	if err != nil {
		ZapLog().With(zap.Error(err)).Error("CardDatesheet ListWithConds err")
		this.ExceptionSerive(ctx, apibackend.BASERR_DATABASE_ERROR.Code(), apibackend.BASERR_DATABASE_ERROR.Desc())
		return
	}
	arr, ok := results.List.(*[]*models.CardClasssheet)
	if ok {

		for i := 0; i < len(*arr); i++ {
			if (*arr)[i].ClassTp != nil {
				ss, err :=new(models.CardClass).GetByIdFromCache(*(*arr)[i].ClassTp)
				if ss != nil && err == nil {
					(*arr)[i].ClassName = ss.Name
				}
			}
		}
	}
	this.Response(ctx, results)
}
