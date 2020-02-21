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

type CardAreaSheet struct {
	Controllers
}

func (this *CardAreaSheet) BkList(ctx iris.Context) {
	param := new(api.BkCardAreaSheetList)

	err := Tools.ShouldBindJSON(ctx, param)
	if err != nil {
		this.ExceptionSerive(ctx, apibackend.BASERR_INVALID_PARAMETER.Code(), apibackend.BASERR_INVALID_PARAMETER.Desc())
		ZapLog().Error("param err", zap.Error(err))
		return
	}

	conds := make([]*models.SqlPairCondition, 0, 3)
	if param.LikeStr != nil {
		conds = append(conds, &models.SqlPairCondition{"city like ?", "%" + *param.LikeStr + "%"})
	}

	if param.StartCreatedAt != nil {
		conds = append(conds, &models.SqlPairCondition{"start_created_at >= ?", param.StartCreatedAt})
	}
	if param.EndCreatedAt != nil {
		conds = append(conds, &models.SqlPairCondition{"end_created_at <= ?", param.EndCreatedAt})
	}
	selectFileds := []string{"date, sum(order_count)"}

	groupStr := "date"

	if (param.GroupProvinceFlag != nil) && (*param.GroupProvinceFlag == 1){
		selectFileds = append(selectFileds, "province")
		groupStr += ",province"
	}
	if (param.GroupCityFlag != nil) && (*param.GroupCityFlag == 1) {
		selectFileds = append(selectFileds, "city")
		groupStr += ",city"
	}

	//results, err := new(models.CardOrder).BkParseList(param).ListAreaCountWithConds(param.Page, param.Size, cond )
	//if err != nil {
	//	ZapLog().With(zap.Error(err)).Error("Verify err")
	//	this.ExceptionSerive(ctx, apibackend.BASERR_DATABASE_ERROR.Code(), apibackend.BASERR_DATABASE_ERROR.Desc())
	//	return
	//}
	results, err := new(models.CardAreasheet).ParseList(param).ListWithConds(param.Page, param.Size, selectFileds, conds, groupStr)
	if err != nil {
		ZapLog().With(zap.Error(err)).Error("Verify err")
		this.ExceptionSerive(ctx, apibackend.BASERR_DATABASE_ERROR.Code(), apibackend.BASERR_DATABASE_ERROR.Desc())
		return
	}
	arr, ok := results.List.(*[]*models.CardAreasheet)
	if ok {
		for i := 0; i < len(*arr); i++ {
			/*if (*arr)[i].ClassTp != nil {
				ss, ok := models.ClassTpMap[*(*arr)[i].ClassTp]
				if ok {
					(*arr)[i].ClassName = &ss.Name
				}
			}
			if (*arr)[i].ClassISP != nil {
				(*arr)[i].ClassIspName = new(string)
				switch *(*arr)[i].ClassISP {
				case models.CONST_ISP_Dianxin:
					*(*arr)[i].ClassIspName = models.CONST_ISP_Dianxin_Name
				case models.CONST_ISP_YiDong:
					*(*arr)[i].ClassIspName = models.CONST_ISP_YiDong_Name
				case models.CONST_ISP_LiTong:
					*(*arr)[i].ClassIspName = models.CONST_ISP_LiTong_Name
				}
			}*/
		}
	}
	this.Response(ctx, results)
}
