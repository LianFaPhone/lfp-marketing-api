package order

import (
	apibackend "LianFaPhone/lfp-api/errdef"
	. "LianFaPhone/lfp-base/log/zap"
	"LianFaPhone/lfp-marketing-api/api"
	. "LianFaPhone/lfp-marketing-api/controllers"
	"LianFaPhone/lfp-marketing-api/idcard-api"
	"LianFaPhone/lfp-marketing-api/models"
	"github.com/kataras/iris"
	"go.uber.org/zap"
)

func (this *CardOrder) BkRetryApply(ctx iris.Context) {
	param := new(api.BkCardOrderThirdApply)

	err := Tools.ShouldBindJSON(ctx, param)
	if err != nil {
		this.ExceptionSerive(ctx, apibackend.BASERR_INVALID_PARAMETER.Code(), apibackend.BASERR_INVALID_PARAMETER.Desc())
		ZapLog().Error("param err", zap.Error(err))
		return
	}

	err = new(models.CardOrder).UpdatesStatusByOrderNo(param.OrderNo, models.CONST_OrderStatus_Retry_Apply_Doing)
	if err != nil {
		ZapLog().With(zap.Error(err)).Error("UpdatesStatusByOrderNo err")
		this.ExceptionSerive(ctx, apibackend.BASERR_DATABASE_ERROR.Code(), apibackend.BASERR_DATABASE_ERROR.Desc())
		return
	}
	go func(){
		for i:=0; i< len(param.OrderNo);i++{
			log := "管理员|重试下单开始"
			new(models.CardOrderLog).FtParseAdd(nil, param.OrderNo[i], &log)
		}
	}()
	this.Response(ctx, nil)
}

func (this *CardOrder) BkIdCardCheck(ctx iris.Context) {
	param := new(api.BkCardOrderIdCardCheck)

	err := Tools.ShouldBindJSON(ctx, param)
	if err != nil {
		this.ExceptionSerive(ctx, apibackend.BASERR_INVALID_PARAMETER.Code(), apibackend.BASERR_INVALID_PARAMETER.Desc())
		ZapLog().Error("param err", zap.Error(err))
		return
	}

	go func() {
		for i := 0; i < len(param.OrderNo); i++ {
			cOrder, err := new(models.CardOrder).GetByOrderNo(param.OrderNo[i], nil)
			if err != nil {
				ZapLog().With(zap.Error(err)).Error("GetById err")
				return
			}
			if cOrder == nil || cOrder.TrueName == nil || cOrder.IdCard == nil || len(*cOrder.TrueName) == 0 || len(*cOrder.IdCard) == 0 {
				continue
			}
			if cOrder.IdCardAudit != nil && *cOrder.IdCardAudit != 0 {
				continue
			}
			res, err := new(idcard_api.ReIdCardCheck).Send(*cOrder.TrueName, *cOrder.IdCard)
			if err != nil {
				ZapLog().With(zap.Error(err)).Error("ReIdCardCheck err")
				return
			}
			status := 0 //未检测
			switch res.Result {
			case idcard_api.CONNST_LANCHUANG_IDCARDCHECK_Same:
				status = 1 //匹配
			case idcard_api.CONNST_LANCHUANG_IDCARDCHECK_NotSame:
				status = 2 //不匹配
			case idcard_api.CONNST_LANCHUANG_IDCARDCHECK_NotSure:
				status = 3 //未知
			default:
				status = 3
			}
			modelsParam := &models.CardOrder{
				OrderNo:          &param.OrderNo[i],
				IdCardAudit: &status,
			}

			if err := modelsParam.UpdateByOrderNo(); err != nil {
				ZapLog().With(zap.Error(err)).Error("Update err")
			}
		}
	}()

	this.Response(ctx, nil)
}