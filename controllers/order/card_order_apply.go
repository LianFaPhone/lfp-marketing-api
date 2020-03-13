package order

import (
	"LianFaPhone/lfp-marketing-api/api"
	"LianFaPhone/lfp-marketing-api/common"
	"LianFaPhone/lfp-marketing-api/config"
	"LianFaPhone/lfp-marketing-api/db"
	"LianFaPhone/lfp-marketing-api/models"
	"LianFaPhone/lfp-marketing-api/sdk"
	"fmt"
	"go.uber.org/zap"
	"time"
)

import (
	apibackend "LianFaPhone/lfp-api/errdef"
	. "LianFaPhone/lfp-base/log/zap"
	. "LianFaPhone/lfp-marketing-api/controllers"
	"github.com/kataras/iris"
)

//申请订单
func (this *CardOrder) Apply(ctx iris.Context) {
	param := new(api.CardOrderApply)

	err := Tools.ShouldBindJSON(ctx, param)
	if err != nil {
		this.ExceptionSerive(ctx, apibackend.BASERR_INVALID_PARAMETER.Code(), apibackend.BASERR_INVALID_PARAMETER.Desc())
		ZapLog().Error("param err", zap.Error(err))
		return
	}
	countryCode := "0086"
	if param.CountryCode != nil && len(*param.CountryCode) > 0 {
		countryCode = *param.CountryCode
	}

	/*********薅羊毛控制*********************/
	limitCardCount := 3
	limitCardPeriod := int64(2592000)
	if param.PartnerId != nil {
		cc, _ := new(models.PdPartner).GetByIdFromCache(*param.PartnerId)
		if cc != nil {
			if cc.LimitCardCount != nil {
				limitCardCount = *cc.LimitCardCount
			}
			if cc.LimitCardPeriod != nil {
				limitCardPeriod = *cc.LimitCardPeriod
			}
		}
	}
	/******************************/

	//检测id是否有效
	recipient := countryCode + *param.Phone
	if param.VerifyId != nil {
		verification := common.NewVerification(&db.GRedis, "sim", "")
		flag, err := verification.Check(*param.VerifyId, 0, recipient)
		if err != nil {
			ZapLog().With(zap.Error(err)).Error("Verify err")
			this.ExceptionSerive(ctx, apibackend.BASERR_DATABASE_ERROR.Code(), apibackend.BASERR_DATABASE_ERROR.Desc())
			return
		}
		if !flag {
			ZapLog().With(zap.Error(err)).Error("Sms Verify err")
			this.ExceptionSerive(ctx, apibackend.BASERR_ACTIVITY_FISSIONSHARE_SMS_INCORRECT_VERIFYCODE.Code(), apibackend.BASERR_ACTIVITY_FISSIONSHARE_SMS_INCORRECT_VERIFYCODE.OriginDesc())
			return
		}
	}

	param.IP = common.GetRealIp(ctx)

	//这个以后再想
	orderNo := fmt.Sprintf("D%s%s%03d", config.GConfig.Server.DevId, time.Now().Format("060102030405000"), GIdGener.Gen())
	param.Status = new(int)
	*param.Status = models.CONST_OrderStatus_New

	if (param.FinishFlag != nil) && (*param.FinishFlag == 0) {
		*param.Status = models.CONST_OrderStatus_New_UnFinish
	}

	modelParam := new(models.CardOrder).FtParseAdd(param, orderNo)

	//	tx := db.GDbMgr.Get().Begin()
	upFlag, err := modelParam.LimitCheckByIdCardAndTime(*param.IdCard, time.Now().Unix()-limitCardPeriod, *param.PartnerGoodsCode, limitCardCount)
	if err != nil {
		//		tx.Rollback()
		ZapLog().With(zap.Error(err)).Error("database err")
		this.ExceptionSerive(ctx, apibackend.BASERR_DATABASE_ERROR.Code(), apibackend.BASERR_DATABASE_ERROR.Desc())
		return
	}
	if upFlag {
		//		tx.Rollback()
		this.ExceptionSerive(ctx, apibackend.BASERR_OBJECT_EXISTS.Code(), apibackend.BASERR_OBJECT_EXISTS.Desc())
		return
	}
	if param.Number != nil && param.NumberPoolFlag != nil && *param.NumberPoolFlag == 1 {
		succFlag, err := new(models.PhoneNumberPool).UseNumberByNumber(*param.Number, time.Now().Unix(), param.SessionId, orderNo, *param.TrueName, *param.Phone)
		if err != nil {
			//			tx.Rollback()
			ZapLog().With(zap.Error(err)).Error("database err")
			this.ExceptionSerive(ctx, apibackend.BASERR_DATABASE_ERROR.Code(), apibackend.BASERR_DATABASE_ERROR.Desc())
			return
		}
		if !succFlag {
			ZapLog().Error("phone Number be use err")
			this.ExceptionSerive(ctx, apibackend.BASERR_CARDMARKET_PHONEPOOL_USE_FAIL.Code(), apibackend.BASERR_CARDMARKET_PHONEPOOL_USE_FAIL.Desc())
			return
		}
	}

	//插入数据库
	if err := modelParam.Add(); err != nil {
		//		tx.Rollback()
		ZapLog().With(zap.Error(err)).Error("database err")
		this.ExceptionSerive(ctx, apibackend.BASERR_DATABASE_ERROR.Code(), apibackend.BASERR_DATABASE_ERROR.Desc())
		return
	}
	//	tx.Commit()

	go func() {
		if param.Log != nil && len(*param.Log) > 1 {
			new(models.CardOrderLog).FtParseAdd(nil, &orderNo, param.Log).Add()
		}

		return
		if err := sdk.GNotifySdk.SendSms(nil, *param.Phone, "wangka_complete", 0, nil); err != nil {
			ZapLog().With(zap.Error(err), zap.String("phone", *param.Phone)).Error("GNotifySdk.SendSms[wangka_complete] err")
		}
	}()
	this.Response(ctx, &api.ResCardOrderApply{orderNo})
}

func (this *CardOrder) FtConfirm(ctx iris.Context) {
	param := new(api.FtCardOrderStatus)

	err := Tools.ShouldBindJSON(ctx, param)
	if err != nil {
		this.ExceptionSerive(ctx, apibackend.BASERR_INVALID_PARAMETER.Code(), apibackend.BASERR_INVALID_PARAMETER.Desc())
		ZapLog().Error("param err", zap.Error(err))
		return
	}
	if (param.SuccFlag != nil) && (*param.SuccFlag !=0) {
		err = new(models.CardOrder).UpdateStatusByOrderNo(param.OrderNo, models.CONST_OrderStatus_New)
		if err != nil {
			ZapLog().With(zap.Error(err)).Error("database err")
			this.ExceptionSerive(ctx, apibackend.BASERR_DATABASE_ERROR.Code(), apibackend.BASERR_DATABASE_ERROR.Desc())
			return
		}
	}


	if param.Log != nil && len(*param.Log) > 1 {
		go func (){
			new(models.CardOrderLog).FtParseAdd(nil, &param.OrderNo, param.Log).Add()
		}()
	}
	this.Response(ctx, nil)
}

//申请订单
func (this *CardOrder) ApplyFulfil(ctx iris.Context) {
	ctx.Params().Get("order_no")
	ctx.Params().Get("phone")
	ctx.Params().Get("code")
	//检测是否合法
	//更新信息
}


func (this *CardOrder) FtUpdatePhotoUrls(ctx iris.Context) {
	param := new(api.FtCardOrderPhotoUrlUpdates)

	err := Tools.ShouldBindJSON(ctx, param)
	if err != nil {
		this.ExceptionSerive(ctx, apibackend.BASERR_INVALID_PARAMETER.Code(), apibackend.BASERR_INVALID_PARAMETER.Desc())
		ZapLog().Error("param err", zap.Error(err))
		return
	}

	if param.ActiveKey != nil {
		activeCode, err := new(models.ActiveCode).GetBy(*param.ActiveKey, *param.OrderNo, *param.Phone)
		if err != nil {
			ZapLog().With(zap.Error(err)).Error("Update err")
			this.ExceptionSerive(ctx, apibackend.BASERR_DATABASE_ERROR.Code(), apibackend.BASERR_DATABASE_ERROR.Desc())
			return
		}
		if activeCode == nil {
			this.ExceptionSerive(ctx, apibackend.BASERR_TOKEN_EXPIRED.Code(), apibackend.BASERR_TOKEN_EXPIRED.Desc())
			return
		}
		if activeCode.Count != nil && *activeCode.Count >= 10 {
			this.ExceptionSerive(ctx, apibackend.BASERR_TOKEN_EXPIRED.Code(), apibackend.BASERR_TOKEN_EXPIRED.Desc())
			return
		}
		new(models.ActiveCode).RecordCount(*activeCode.Id)
	}


	// 检测请求有效性
	err = new(models.CardOrder).UpdatesPhotosByOrderNo(*param.OrderNo, *param.Dataurl1, *param.Dataurl2, *param.Dataurl3)
	if err != nil {
		ZapLog().With(zap.Error(err)).Error("Update err")
		this.ExceptionSerive(ctx, apibackend.BASERR_DATABASE_ERROR.Code(), apibackend.BASERR_DATABASE_ERROR.Desc())
		return
	}

	this.Response(ctx, nil)
}