package order

import (
	apibackend "LianFaPhone/lfp-api/errdef"
	. "LianFaPhone/lfp-base/log/zap"
	"LianFaPhone/lfp-marketing-api/api"
	"LianFaPhone/lfp-marketing-api/common"

	. "LianFaPhone/lfp-marketing-api/controllers"
	"LianFaPhone/lfp-marketing-api/db"
	"LianFaPhone/lfp-marketing-api/idcard-api"
	"LianFaPhone/lfp-marketing-api/models"
	"LianFaPhone/lfp-marketing-api/sdk"
	"LianFaPhone/lfp-marketing-api/tasker"
	"fmt"
	"github.com/kataras/iris"
	"go.uber.org/zap"
	"strings"
	"time"
)

type CardOrder struct {
	IdGener
	Controllers
}
//
//func (this *CardOrder) OldApply(ctx iris.Context) {
//	param := new(api.OldCardOrderApply)
//
//	err := Tools.ShouldBindQuery(ctx, param)
//	if err != nil {
//		this.ExceptionSerive(ctx, apibackend.BASERR_INVALID_PARAMETER.Code(), apibackend.BASERR_INVALID_PARAMETER.Desc())
//		ZapLog().Error("param err", zap.Error(err))
//		return
//	}
//	//countryCode := "0086"
//	//
//	////检测id是否有效
//	//recipient := countryCode + *param.Phone
//	//verification := common.NewVerification(&db.GRedis, "sim", "")
//	//flag,err := verification.Check(*param.VerifyId, 0,  recipient)
//	//if err != nil {
//	//	ZapLog().With(zap.Error(err)).Error("Verify err")
//	//	this.ExceptionSerive(ctx, apibackend.BASERR_DATABASE_ERROR.Code(), apibackend.BASERR_DATABASE_ERROR.Desc())
//	//	return
//	//}
//	//if !flag {
//	//	ZapLog().With(zap.Error(err)).Error("Sms Verify err")
//	//	//this.ExceptionSerive(ctx, apibackend.BASERR_ACTIVITY_FISSIONSHARE_SMS_INCORRECT_VERIFYCODE.Code(), apibackend.BASERR_ACTIVITY_FISSIONSHARE_SMS_INCORRECT_VERIFYCODE.OriginDesc())
//	//	//return
//	//}
//
//	IP := common.GetRealIp(ctx)
//
//	classTp, ok := models.ClassTpMap2[param.Channel]
//	if !ok {
//		ZapLog().Error("bug err")
//		this.ExceptionSerive(ctx, apibackend.BASERR_UNKNOWN_BUG.Code(), apibackend.BASERR_UNKNOWN_BUG.Desc())
//		return
//	}
//
//	//这个以后再想
//	orderNo := fmt.Sprintf("D%s%d", time.Now().Format("060102030405000"), this.Gen())
//
//	Status := models.CONST_OrderStatus_New
//
//	addrArr := strings.Split(param.Address, " ")
//	provice := ""
//	city := ""
//	area := ""
//	town := ""
//	if len(addrArr) == 3 {
//		provice = addrArr[0]
//		city = addrArr[0]
//		area = addrArr[1]
//		town = addrArr[2]
//	} else if len(addrArr) > 3 {
//		provice = addrArr[0]
//		city = addrArr[1]
//		area = addrArr[2]
//		town = addrArr[3]
//	} else if len(addrArr) >= 1 {
//		provice = addrArr[0]
//		city = addrArr[0]
//		area = addrArr[0]
//		town = addrArr[0]
//	}
//
//	modelParam := new(models.CardOrder).FtParseAdd2(param, orderNo, classTp.Tp, Status, provice, city, area, town, param.Address2, IP)
//
//	tx := db.GDbMgr.Get().Begin()
//	uniqueFlag, err := modelParam.TxLockUniqueByIdCardAndTime(tx, *param.IdCard, time.Now().Unix()-24*3600*30*6, classTp.Tp)
//	if err != nil {
//		tx.Rollback()
//		ZapLog().With(zap.Error(err)).Error("database err")
//		this.ExceptionSerive(ctx, apibackend.BASERR_DATABASE_ERROR.Code(), apibackend.BASERR_DATABASE_ERROR.Desc())
//		return
//	}
//	if !uniqueFlag {
//		tx.Rollback()
//		this.ExceptionSerive(ctx, apibackend.BASERR_OBJECT_EXISTS.Code(), apibackend.BASERR_OBJECT_EXISTS.Desc())
//		return
//	}
//
//	//插入数据库
//	if err := modelParam.TxAdd(tx); err != nil {
//		tx.Rollback()
//		ZapLog().With(zap.Error(err)).Error("database err")
//		this.ExceptionSerive(ctx, apibackend.BASERR_DATABASE_ERROR.Code(), apibackend.BASERR_DATABASE_ERROR.Desc())
//		return
//	}
//	tx.Commit()
//
//	go func() {
//		if !config.GConfig.BasNotify.SwitchFlag {
//			return
//		}
//		if err := sdk.GNotifySdk.SendSms(nil, *param.Phone, "wangka_complete", 0); err != nil {
//			ZapLog().With(zap.Error(err), zap.String("phone", *param.Phone)).Error("GNotifySdk.SendSms[wangka_complete] err")
//		}
//	}()
//	this.Response(ctx, nil)
//}



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
	orderNo := fmt.Sprintf("D%s%d", time.Now().Format("060102030405000"), this.Gen())
	param.Status = new(int)
	*param.Status = models.CONST_OrderStatus_New

	if (param.FinishFlag != nil) && (*param.FinishFlag == 0) {
		*param.Status = models.CONST_OrderStatus_New_UnFinish
	}

	modelParam := new(models.CardOrder).FtParseAdd(param, orderNo)

//	tx := db.GDbMgr.Get().Begin()
	upFlag, err := modelParam.LimitCheckByIdCardAndTime(*param.IdCard, time.Now().Unix()-24*3600*30*3, *param.ClassTp, 3)
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
		return
		if err := sdk.GNotifySdk.SendSms(nil, *param.Phone, "wangka_complete", 0); err != nil {
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

	err = new(models.CardOrder).UpdateStatusByOrderNo(param.OrderNo, models.CONST_OrderStatus_New_UnFinish)
	if err != nil {
		ZapLog().With(zap.Error(err)).Error("database err")
		this.ExceptionSerive(ctx, apibackend.BASERR_DATABASE_ERROR.Code(), apibackend.BASERR_DATABASE_ERROR.Desc())
		return
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

func (this *CardOrder) pathToOrderStatus(ctx iris.Context) *int {
	index := strings.LastIndex(ctx.Path(), "-")
	if index <= 0 {
		return nil
	}
	name := ctx.Path()[index:]

	status, _ := models.PathToOrderStatus[name]
	if status <= 0 {
		return nil
	}
	return &status
}

func (this *CardOrder) BkList(ctx iris.Context) {
	param := new(api.BkCardOrderList)

	err := Tools.ShouldBindJSON(ctx, param)
	if err != nil {
		this.ExceptionSerive(ctx, apibackend.BASERR_INVALID_PARAMETER.Code(), apibackend.BASERR_INVALID_PARAMETER.Desc())
		ZapLog().Error("param err", zap.Error(err))
		return
	}

	param.Status = this.pathToOrderStatus(ctx)

	condPair := make([]*models.SqlPairCondition, 0, 5)
	if param.LikeStr != nil && len(*param.LikeStr) > 0 {
		condPair = append(condPair, &models.SqlPairCondition{"true_name like ?", "%" + *param.LikeStr + "%"})
	}
	if param.StartCreatedAt != nil {
		condPair = append(condPair, &models.SqlPairCondition{"created_at >= ?", param.StartCreatedAt})
	}
	if param.EndCreatedAt != nil {
		condPair = append(condPair, &models.SqlPairCondition{"created_at <= ?", param.EndCreatedAt})
	}
	if param.StartDeliverAt != nil {
		condPair = append(condPair, &models.SqlPairCondition{"deliver_at >= ?", param.StartDeliverAt})
	}
	if param.EndDeliverAt != nil {
		condPair = append(condPair, &models.SqlPairCondition{"deliver_at <= ?", param.EndDeliverAt})
	}
	extend := ctx.Values().GetString("extend")
	if len(extend) > 0 {
		arr := strings.Split(extend, ",")
		condPair = append(condPair, &models.SqlPairCondition{"class_tp in (?)", arr})
	}
	if param.UploadFlag != nil {
		if *param.UploadFlag == 0 {
			condPair = append(condPair, &models.SqlPairCondition{"dataurl1 == ?", "null"})
		} else if *param.UploadFlag == 1 {
			condPair = append(condPair, &models.SqlPairCondition{"dataurl1 != ?", "null"})
		}
	}

	results, err := new(models.CardOrder).BkParseList(param).ListWithConds(param.Page, param.Size, nil, condPair)
	if err != nil {
		ZapLog().With(zap.Error(err)).Error("Verify err")
		this.ExceptionSerive(ctx, apibackend.BASERR_DATABASE_ERROR.Code(), apibackend.BASERR_DATABASE_ERROR.Desc())
		return
	}
	coArr := results.List.(*[]*models.CardOrder)
	var areaBlack *models.BlacklistAreaCacheStt
	areaBlackFlag := false
	for i := 0; i < len(*coArr); i++ {
		temp := (*coArr)[i]
		if temp.ClassTp != nil {
			m, _ := models.ClassTpMap[*temp.ClassTp]
			if m != nil {
				temp.ClassName = &m.Name
			}
		}
		if temp.PhoneOSTp != nil {
			m, _ := models.PhoneOsMap[*temp.PhoneOSTp]
			if m != nil {
				temp.PhoneOSName = &m.Name
			}
		}
		if temp.Status != nil {
			m := models.OrderStatusMap[*temp.Status]
			temp.StatusName = &m
		}
		if temp.BspStatus != nil {
			m := models.OrderStatusMap[*temp.BspStatus]
			temp.BspStatusName = &m
		}

		if param.BlackSwitch == nil || *param.BlackSwitch == 0 {
			continue
		}

		if temp.IsBacklist == nil && temp.Phone != nil {
			flag, err := new(models.BlacklistPhone).ExistByPhone(*temp.Phone)
			if err == nil {

			}
			if flag {
				temp.IsBacklist = new(int)
				*temp.IsBacklist = 1
				continue
			}
		}
		if temp.IsBacklist == nil && temp.IdCard != nil {
			flag, err := new(models.BlacklistIdcard).ExistByIdCard(*temp.IdCard)
			if err == nil {

			}
			if flag {
				temp.IsBacklist = new(int)
				*temp.IsBacklist = 1
				continue
			}
		}
		if !areaBlackFlag {
			areaBlackFlag = true
			areaBlack, err = new(models.BlacklistArea).GetFromCache("")
			if err != nil {

			}
		}

		if areaBlack == nil {
			continue
		}
		if temp.ProvinceCode != nil {
			_, ok := areaBlack.ProviceM[*temp.ProvinceCode]
			if ok {
				temp.IsBacklist = new(int)
				*temp.IsBacklist = 1
				continue
			}
		}
		if temp.CityCode != nil {
			_, ok := areaBlack.CityM[*temp.CityCode]
			if ok {
				temp.IsBacklist = new(int)
				*temp.IsBacklist = 1
				continue
			}
		}
		if temp.AreaCode != nil {
			_, ok := areaBlack.AreaM[*temp.AreaCode]
			if ok {
				temp.IsBacklist = new(int)
				*temp.IsBacklist = 1
				continue
			}
		}

	}
	this.Response(ctx, results)
}

func (this *CardOrder) BkUpdate(ctx iris.Context) {
	param := new(api.BkCardOrder)

	err := Tools.ShouldBindJSON(ctx, param)
	if err != nil {
		this.ExceptionSerive(ctx, apibackend.BASERR_INVALID_PARAMETER.Code(), apibackend.BASERR_INVALID_PARAMETER.Desc())
		ZapLog().Error("param err", zap.Error(err))
		return
	}

	err = new(models.CardOrder).BkParse(param).UpdateByOrderNo()
	if err != nil {
		ZapLog().With(zap.Error(err)).Error("Update err")
		this.ExceptionSerive(ctx, apibackend.BASERR_DATABASE_ERROR.Code(), apibackend.BASERR_DATABASE_ERROR.Desc())
		return
	}
	this.Response(ctx, nil)

}

func (this *CardOrder) BkUpdatesStatus(ctx iris.Context) {
	param := new(api.BkCardOrderStatusUpdates)

	err := Tools.ShouldBindJSON(ctx, param)
	if err != nil {
		this.ExceptionSerive(ctx, apibackend.BASERR_INVALID_PARAMETER.Code(), apibackend.BASERR_INVALID_PARAMETER.Desc())
		ZapLog().Error("param err", zap.Error(err))
		return
	}

	err = new(models.CardOrder).UpdatesStatusByOrderNo(param.OrderNo, *param.Status)
	if err != nil {
		ZapLog().With(zap.Error(err)).Error("Update err")
		this.ExceptionSerive(ctx, apibackend.BASERR_DATABASE_ERROR.Code(), apibackend.BASERR_DATABASE_ERROR.Desc())
		return
	}
	this.Response(ctx, nil)

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

//还缺个发送短信的功能
func (this *CardOrder) BkOrderExtraInport(ctx iris.Context) {
	params := make([]*api.BkCardOrderExtraImprot, 0)

	err := ctx.ReadJSON(&params)
	//err := Tools.ShouldBindJSON(ctx, param)
	if err != nil {
		this.ExceptionSerive(ctx, apibackend.BASERR_INVALID_PARAMETER.Code(), apibackend.BASERR_INVALID_PARAMETER.Desc())
		ZapLog().Error("param err", zap.Error(err))
		return
	}

	res := new(api.BkResCardOrderExtraImprot)
	notifyTp := models.CONST_OrderNotifyTp_Express
	for i := 0; i < len(params); i++ {
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
	if res.SuccCount > 0 {
		tasker.GNotifyTasker.Push()
	}

	res.FailCount = len(params) - res.SuccCount

	this.Response(ctx, res)
}

func (this *CardOrder) BkOrderActiveInport(ctx iris.Context) {
	params := make([]*api.BkCardOrderActiveImprot, 0)

	err := ctx.ReadJSON(&params)
	//err := Tools.ShouldBindJSON(ctx, param)
	if err != nil {
		this.ExceptionSerive(ctx, apibackend.BASERR_INVALID_PARAMETER.Code(), apibackend.BASERR_INVALID_PARAMETER.Desc())
		ZapLog().Error("param err", zap.Error(err))
		return
	}

	res := new(api.BkResCardOrderExtraImprot)
//	notifyTp := models.CONST_OrderNotifyTp_Express
	newPhones := make([]*string, 0, 50)
	for i := 0; i < len(params); i++ {
		newPhones = append(newPhones, params[i].NewPhone)
		if i !=0 && (i%50 == 0 || i == len(params) -1){
			aff,err := new(models.CardOrder).UpdatesStatusByNewphone(newPhones, models.CONST_OrderStatus_Already_Activated)
			if err != nil {
				ZapLog().With(zap.Error(err)).Error("Update err")
				this.ExceptionSerive(ctx, apibackend.BASERR_DATABASE_ERROR.Code(), apibackend.BASERR_DATABASE_ERROR.Desc())
				return
			}
			if aff > 0 {
				res.SuccCount += int(aff)
			}
			newPhones = make([]*string, 0, 50)
		}
	}
	if len(newPhones) > 0 {
		aff,err := new(models.CardOrder).UpdatesStatusByNewphone(newPhones, models.CONST_OrderStatus_Already_Activated)
		if err != nil {
			ZapLog().With(zap.Error(err)).Error("Update err")
			this.ExceptionSerive(ctx, apibackend.BASERR_DATABASE_ERROR.Code(), apibackend.BASERR_DATABASE_ERROR.Desc())
			return
		}
		if aff > 0 {
			res.SuccCount += int(aff)
		}
	}

	res.FailCount = len(params) - res.SuccCount

	this.Response(ctx, res)
}

func (this *CardOrder) BkOrderNewInport(ctx iris.Context) {
	params := make([]*api.BkCardOrderNewImport, 0)

	err := ctx.ReadJSON(&params)
	//err := Tools.ShouldBindJSON(ctx, param)
	if err != nil {
		this.ExceptionSerive(ctx, apibackend.BASERR_INVALID_PARAMETER.Code(), apibackend.BASERR_INVALID_PARAMETER.Desc())
		ZapLog().Error("param err", zap.Error(err))
		return
	}

	res := new(api.BkResCardOrderExtraImprot)
	//	notifyTp := models.CONST_OrderNotifyTp_Express
	for i := 0; i < len(params); i++ {
		//同一个套餐，同一个身份证三个月内出现过就用同一个
		modelParam := &models.CardOrder{
			IdCard: params[i].IdCard,
			Phone:  params[i].Phone,
		}
		order,err := modelParam.GetByIdcardAndPhone(nil)
		if err != nil {

		}
		if order != nil {
			continue
		}


	}

	res.FailCount = len(params) - res.SuccCount

	this.Response(ctx, res)
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
			cOrder, err := new(models.CardOrder).GetByOrderNo(param.OrderNo[i])
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
