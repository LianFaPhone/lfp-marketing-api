package order

import(
	"github.com/kataras/iris"
	"go.uber.org/zap"
	. "LianFaPhone/lfp-marketing-api/controllers"
	"LianFaPhone/lfp-marketing-api/api"
	apibackend "LianFaPhone/lfp-api/errdef"
	. "LianFaPhone/lfp-base/log/zap"
	"LianFaPhone/lfp-marketing-api/common"
	"LianFaPhone/lfp-marketing-api/db"
	"LianFaPhone/lfp-marketing-api/models"
	"time"
	"fmt"
)

type CardOrder struct{
	IdGener
	Controllers
}

//申请订单
func (this * CardOrder) Apply(ctx iris.Context) {
	param := new(api.CardOrderApply)

	err := Tools.ShouldBindJSON(ctx, param)
	if err != nil {
		this.ExceptionSerive(ctx, apibackend.BASERR_INVALID_PARAMETER.Code(), apibackend.BASERR_INVALID_PARAMETER.Desc())
		ZapLog().Error( "param err", zap.Error(err))
		return
	}
	countryCode := "0086"
	if param.CountryCode !=nil && len(*param.CountryCode) > 0 {
		countryCode = *param.CountryCode
	}

	//检测id是否有效
	recipient := countryCode + *param.Phone
	verification := common.NewVerification(&db.GRedis, "sim", "")
	flag,err := verification.Check(*param.VerifyId, 0,  recipient)
	if err != nil {
		ZapLog().With(zap.Error(err)).Error("Verify err")
		this.ExceptionSerive(ctx, apibackend.BASERR_DATABASE_ERROR.Code(), apibackend.BASERR_DATABASE_ERROR.Desc())
		return
	}
	if !flag {
		ZapLog().With(zap.Error(err)).Error("Sms Verify err")
		//this.ExceptionSerive(ctx, apibackend.BASERR_ACTIVITY_FISSIONSHARE_SMS_INCORRECT_VERIFYCODE.Code(), apibackend.BASERR_ACTIVITY_FISSIONSHARE_SMS_INCORRECT_VERIFYCODE.OriginDesc())
		//return
	}

	param.IP = common.GetRealIp(ctx)

	//这个以后再想
	orderNo := fmt.Sprintf("D%s%d", time.Now().Format("060102030405000"), this.Gen())
	param.Status = new(int)
	*param.Status = models.CONST_OrderStatus_New

	modelParam := new(models.CardOrder).FtParseAdd(param, orderNo)


	tx := db.GDbMgr.Get().Begin()
	uniqueFlag,err := modelParam.TxLockUniqueByIdCardAndTime(tx, *param.IdCard, time.Now().Unix() - 24*3600*30*6, *param.ClassTp)
	if err != nil {
		tx.Rollback()
		ZapLog().With(zap.Error(err)).Error("database err")
		this.ExceptionSerive(ctx, apibackend.BASERR_DATABASE_ERROR.Code(), apibackend.BASERR_DATABASE_ERROR.Desc())
		return
	}
	if !uniqueFlag {
		tx.Rollback()
		this.ExceptionSerive(ctx, apibackend.BASERR_OBJECT_EXISTS.Code(), apibackend.BASERR_OBJECT_EXISTS.Desc())
		return
	}
	if param.Number != nil {
		succFlag, err := new(models.PhoneNumberPool).TxUseNumber(tx, *param.Number)
		if err != nil {
			tx.Rollback()
			ZapLog().With(zap.Error(err)).Error("database err")
			this.ExceptionSerive(ctx, apibackend.BASERR_DATABASE_ERROR.Code(), apibackend.BASERR_DATABASE_ERROR.Desc())
			return
		}
		if !succFlag{
			tx.Rollback()
			ZapLog().Error("phone Number be use err")
			this.ExceptionSerive(ctx, 1, "be use")
			return
		}
	}

	//插入数据库
	if err := modelParam.TxAdd(tx); err != nil {
		tx.Rollback()
		ZapLog().With(zap.Error(err)).Error("database err")
		this.ExceptionSerive(ctx, apibackend.BASERR_DATABASE_ERROR.Code(), apibackend.BASERR_DATABASE_ERROR.Desc())
		return
	}
	tx.Commit()
	this.Response(ctx, nil)
}

func (this * CardOrder) BkList(ctx iris.Context) {
	param := new(api.BkCardOrderList)

	err := Tools.ShouldBindJSON(ctx, param)
	if err != nil {
		this.ExceptionSerive(ctx, apibackend.BASERR_INVALID_PARAMETER.Code(), apibackend.BASERR_INVALID_PARAMETER.Desc())
		ZapLog().Error( "param err", zap.Error(err))
		return
	}
	condPair := make([]*models.SqlPairCondition, 0, 5)
	if param.LikeStr != nil && len(*param.LikeStr) > 0 {
		condPair= append(condPair, &models.SqlPairCondition{"true_name like ?", "%"+*param.LikeStr+"%"})
	}
	if param.StartCreatedAt != nil {
		condPair= append(condPair, &models.SqlPairCondition{"created_at >= ?", param.StartCreatedAt})
	}
	if param.EndCreatedAt != nil {
		condPair= append(condPair, &models.SqlPairCondition{"created_at <= ?", param.EndCreatedAt})
	}
	if param.StartDeliverAt != nil {
		condPair= append(condPair, &models.SqlPairCondition{"deliver_at >= ?", param.StartDeliverAt})
	}
	if param.EndDeliverAt != nil {
		condPair= append(condPair, &models.SqlPairCondition{"deliver_at <= ?", param.EndDeliverAt})
	}

	results, err := new(models.CardOrder).BkParseList(param).ListWithConds(param.Page, param.Size, nil ,condPair )
	if err != nil {
		ZapLog().With(zap.Error(err)).Error("Verify err")
		this.ExceptionSerive(ctx, apibackend.BASERR_DATABASE_ERROR.Code(), apibackend.BASERR_DATABASE_ERROR.Desc())
		return
	}
	coArr := results.List.(*[]*models.CardOrder)
	for i:=0; i < len(*coArr);i++ {
		temp := (*coArr)[i]
		if temp.ClassTp != nil {
			m,_ := models.ClassTpMap[*temp.ClassTp]
			if m != nil {
				temp.ClassName = &m.Name
			}
		}
		if temp.PhoneOSTp != nil {
			m,_ := models.PhoneOsMap[*temp.PhoneOSTp]
			if m != nil {
				temp.PhoneOSName = &m.Name
			}
		}
		if temp.Status != nil {
			m:= models.OrderStatusMap[*temp.Status]
			temp.StatusName = &m
		}

	}
	this.Response(ctx, results)
}

func (this * CardOrder) BkUpdate(ctx iris.Context) {
	param := new(api.BkCardOrder)

	err := Tools.ShouldBindJSON(ctx, param)
	if err != nil {
		this.ExceptionSerive(ctx, apibackend.BASERR_INVALID_PARAMETER.Code(), apibackend.BASERR_INVALID_PARAMETER.Desc())
		ZapLog().Error( "param err", zap.Error(err))
		return
	}

	err = new(models.CardOrder).BkParse(param).Update()
	if err != nil {
		ZapLog().With(zap.Error(err)).Error("Update err")
		this.ExceptionSerive(ctx, apibackend.BASERR_DATABASE_ERROR.Code(), apibackend.BASERR_DATABASE_ERROR.Desc())
		return
	}
	this.Response(ctx, nil)


}

func (this * CardOrder) BkUpdatesStatus(ctx iris.Context) {
	param := new(api.BkCardOrderStatusUpdates)

	err := Tools.ShouldBindJSON(ctx, param)
	if err != nil {
		this.ExceptionSerive(ctx, apibackend.BASERR_INVALID_PARAMETER.Code(), apibackend.BASERR_INVALID_PARAMETER.Desc())
		ZapLog().Error( "param err", zap.Error(err))
		return
	}

	err = new(models.CardOrder).UpdatesStatus(param.Id, *param.Status)
	if err != nil {
		ZapLog().With(zap.Error(err)).Error("Update err")
		this.ExceptionSerive(ctx, apibackend.BASERR_DATABASE_ERROR.Code(), apibackend.BASERR_DATABASE_ERROR.Desc())
		return
	}
	this.Response(ctx, nil)


}


func (this * CardOrder) FtUpdatePhotoUrls(ctx iris.Context) {
	param := new(api.FtCardOrderPhotoUrlUpdates)

	err := Tools.ShouldBindJSON(ctx, param)
	if err != nil {
		this.ExceptionSerive(ctx, apibackend.BASERR_INVALID_PARAMETER.Code(), apibackend.BASERR_INVALID_PARAMETER.Desc())
		ZapLog().Error( "param err", zap.Error(err))
		return
	}

	activeCode,err := new(models.ActiveCode).GetBy(*param.ActiveKey, *param.OrderNo, *param.Phone)
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

	// 检测请求有效性
	err = new(models.CardOrder).UpdatesPhotosByOrderNo(*param.OrderNo, *param.Dataurl1, *param.Dataurl2, *param.Dataurl3)
	if err != nil {
		ZapLog().With(zap.Error(err)).Error("Update err")
		this.ExceptionSerive(ctx, apibackend.BASERR_DATABASE_ERROR.Code(), apibackend.BASERR_DATABASE_ERROR.Desc())
		return
	}

	new(models.ActiveCode).RecordCount(*activeCode.Id)

	this.Response(ctx, nil)
}

//还缺个发送短信的功能
func (this * CardOrder) BkOrderExtraInport(ctx iris.Context) {
	params := make([]*api.BkCardOrderExtraImprot, 0)

	err := ctx.ReadJSON(&params)
	//err := Tools.ShouldBindJSON(ctx, param)
	if err != nil {
		this.ExceptionSerive(ctx, apibackend.BASERR_INVALID_PARAMETER.Code(), apibackend.BASERR_INVALID_PARAMETER.Desc())
		ZapLog().Error( "param err", zap.Error(err))
		return
	}

	res := new(api.BkResCardOrderExtraImprot)
	notifyTp := models.CONST_OrderNotifyTp_Express
	for i:=0; i < len(params); i++ {
		aff, err := new(models.CardOrder).BkParseExtraImport(params[i]).UpdatesByOrderNo()
		if err != nil {
			ZapLog().With(zap.Error(err)).Error("Update err")
			this.ExceptionSerive(ctx, apibackend.BASERR_DATABASE_ERROR.Code(), apibackend.BASERR_DATABASE_ERROR.Desc())
			return
		}
		if aff > 0 {
			res.SuccCount += 1
		}
		if err = new(models.CardOrderNotify).Add(params[i].OrderNo, nil, &notifyTp); err != nil {
			ZapLog().With(zap.Error(err)).Error("CardOrderNotify.Add err")
		}
	}
	res.FailCount = len(params) - res.SuccCount

	this.Response(ctx, res)
}