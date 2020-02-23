package order

import (
	apibackend "LianFaPhone/lfp-api/errdef"
	. "LianFaPhone/lfp-base/log/zap"
	"LianFaPhone/lfp-marketing-api/api"
	"LianFaPhone/lfp-marketing-api/common"
	"LianFaPhone/lfp-marketing-api/config"
	"github.com/go-redis/redis"

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
	//IdGener
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

	limit := 3
	if param.ClassName != nil {
		cc, err := new(models.CardClass).GetByNameFromCache(*param.ClassName)
		if err == nil && cc != nil && cc.MaxLimit != nil {
			limit = *cc.MaxLimit
		}
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
	orderNo := fmt.Sprintf("D%s%s%d", config.GConfig.Server.DevId, time.Now().Format("060102030405000"), GIdGener.Gen())
	param.Status = new(int)
	*param.Status = models.CONST_OrderStatus_New

	if (param.FinishFlag != nil) && (*param.FinishFlag == 0) {
		*param.Status = models.CONST_OrderStatus_New_UnFinish
	}

	modelParam := new(models.CardOrder).FtParseAdd(param, orderNo)

//	tx := db.GDbMgr.Get().Begin()
	upFlag, err := modelParam.LimitCheckByIdCardAndTime(*param.IdCard, time.Now().Unix()-24*3600*30*3, *param.ClassTp, limit)
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

func (this *CardOrder) pathToOrderStatus(ctx iris.Context) *int {
	index := strings.LastIndex(ctx.Path(), "-")
	if index <= 0 {
		return nil
	}
	name := ctx.Path()[index+1:]

	status, _ := models.PathToOrderStatus[name]
	if status <= 0 {
		return nil
	}
	return &status
}

func (this *CardOrder) bkSubList(ctx iris.Context, status *int) {
	param := new(api.BkCardOrderList)

	err := Tools.ShouldBindJSON(ctx, param)
	if err != nil {
		this.ExceptionSerive(ctx, apibackend.BASERR_INVALID_PARAMETER.Code(), apibackend.BASERR_INVALID_PARAMETER.Desc())
		ZapLog().Error("param err", zap.Error(err))
		return
	}
	if param.Status == nil {
		param.Status = status
	}

	condPair := make([]*models.SqlPairCondition, 0, 10)
	if param.LikeStr != nil && len(*param.LikeStr) > 0 {
		condPair = append(condPair, &models.SqlPairCondition{"true_name like ?", "%" + *param.LikeStr + "%"})
		condPair = append(condPair, &models.SqlPairCondition{"order_no like ?", "%" + *param.LikeStr + "%"})
		condPair = append(condPair, &models.SqlPairCondition{"phone like ?", "%" + *param.LikeStr + "%"})
		condPair = append(condPair, &models.SqlPairCondition{"new_phone like ?", "%" + *param.LikeStr + "%"})
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
	//extend := ctx.Values().GetString("extend")
	//if len(extend) > 0 {
	//	arr := strings.Split(extend, ",")
	//	condPair = append(condPair, &models.SqlPairCondition{"class_tp in (?)", arr})
	//}
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
			//ZapLog().Info("cardclass 1")
			cc,err := new(models.CardClass).GetByIdFromCache(*temp.ClassTp)
			//ZapLog().Info("cardclass 2", zap.Error(err), zap.Any("cc",cc))
			if err == nil && cc != nil{
				temp.ClassName = cc.Detail
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
		//if temp.BspStatus != nil {
		//	m := models.OrderStatusMap[*temp.BspStatus]
		//	temp.BspStatusName = &m
		//}

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

func (this *CardOrder) BkList(ctx iris.Context) {
	this.bkSubList(ctx, nil)
}

func (this *CardOrder) BkListAll(ctx iris.Context) {
	this.bkSubList(ctx, nil)
}

func (this *CardOrder) BkListNew(ctx iris.Context) {
	status := models.CONST_OrderStatus_New
	this.bkSubList(ctx, &status)
}

func (this *CardOrder) BkListNewunfinish(ctx iris.Context) {
	status := models.CONST_OrderStatus_New_UnFinish
	this.bkSubList(ctx, &status)
}

func (this *CardOrder) BkListExport(ctx iris.Context) {
	status := models.CONST_OrderStatus_Already_Export
	this.bkSubList(ctx, &status)
}

func (this *CardOrder) BkListDeliver(ctx iris.Context) {
	status := models.CONST_OrderStatus_Already_Delivered
	this.bkSubList(ctx, &status)
}

func (this *CardOrder) BkListWaitdone(ctx iris.Context) {
	status := models.CONST_OrderStatus_Wait_Done
	this.bkSubList(ctx, &status)
}

func (this *CardOrder) BkListAlreadydone(ctx iris.Context) {
	status := models.CONST_OrderStatus_Already_Done
	this.bkSubList(ctx, &status)
}

func (this *CardOrder) BkListRecyclebin(ctx iris.Context) {
	status := models.CONST_OrderStatus_Recyclebin
	this.bkSubList(ctx, &status)
}

func (this *CardOrder) BkListUnmatch(ctx iris.Context) {
	status := models.CONST_OrderStatus_UnMatch
	this.bkSubList(ctx, &status)
}

func (this *CardOrder) BkListActivated(ctx iris.Context) {
	status := models.CONST_OrderStatus_Already_Activated
	this.bkSubList(ctx, &status)
}
//func (this *CardOrder) BkList(ctx iris.Context) {
//	param := new(api.BkCardOrderList)
//
//	err := Tools.ShouldBindJSON(ctx, param)
//	if err != nil {
//		this.ExceptionSerive(ctx, apibackend.BASERR_INVALID_PARAMETER.Code(), apibackend.BASERR_INVALID_PARAMETER.Desc())
//		ZapLog().Error("param err", zap.Error(err))
//		return
//	}
//	if param.Status == nil {
//		param.Status = this.pathToOrderStatus(ctx)
//	}
//
//	condPair := make([]*models.SqlPairCondition, 0, 10)
//	if param.LikeStr != nil && len(*param.LikeStr) > 0 {
//		condPair = append(condPair, &models.SqlPairCondition{"true_name like ?", "%" + *param.LikeStr + "%"})
//		condPair = append(condPair, &models.SqlPairCondition{"order_no like ?", "%" + *param.LikeStr + "%"})
//		condPair = append(condPair, &models.SqlPairCondition{"phone like ?", "%" + *param.LikeStr + "%"})
//		condPair = append(condPair, &models.SqlPairCondition{"new_phone like ?", "%" + *param.LikeStr + "%"})
//	}
//	if param.StartCreatedAt != nil {
//		condPair = append(condPair, &models.SqlPairCondition{"created_at >= ?", param.StartCreatedAt})
//	}
//	if param.EndCreatedAt != nil {
//		condPair = append(condPair, &models.SqlPairCondition{"created_at <= ?", param.EndCreatedAt})
//	}
//	if param.StartDeliverAt != nil {
//		condPair = append(condPair, &models.SqlPairCondition{"deliver_at >= ?", param.StartDeliverAt})
//	}
//	if param.EndDeliverAt != nil {
//		condPair = append(condPair, &models.SqlPairCondition{"deliver_at <= ?", param.EndDeliverAt})
//	}
//	//extend := ctx.Values().GetString("extend")
//	//if len(extend) > 0 {
//	//	arr := strings.Split(extend, ",")
//	//	condPair = append(condPair, &models.SqlPairCondition{"class_tp in (?)", arr})
//	//}
//	if param.UploadFlag != nil {
//		if *param.UploadFlag == 0 {
//			condPair = append(condPair, &models.SqlPairCondition{"dataurl1 == ?", "null"})
//		} else if *param.UploadFlag == 1 {
//			condPair = append(condPair, &models.SqlPairCondition{"dataurl1 != ?", "null"})
//		}
//	}
//
//	results, err := new(models.CardOrder).BkParseList(param).ListWithConds(param.Page, param.Size, nil, condPair)
//	if err != nil {
//		ZapLog().With(zap.Error(err)).Error("Verify err")
//		this.ExceptionSerive(ctx, apibackend.BASERR_DATABASE_ERROR.Code(), apibackend.BASERR_DATABASE_ERROR.Desc())
//		return
//	}
//	coArr := results.List.(*[]*models.CardOrder)
//	var areaBlack *models.BlacklistAreaCacheStt
//	areaBlackFlag := false
//	for i := 0; i < len(*coArr); i++ {
//		temp := (*coArr)[i]
//		if temp.ClassTp != nil {
//			//ZapLog().Info("cardclass 1")
//			cc,err := new(models.CardClass).GetByIdFromCache(*temp.ClassTp)
//			//ZapLog().Info("cardclass 2", zap.Error(err), zap.Any("cc",cc))
//			if err == nil && cc != nil{
//				temp.ClassName = cc.Detail
//			}
//		}
//		if temp.PhoneOSTp != nil {
//			m, _ := models.PhoneOsMap[*temp.PhoneOSTp]
//			if m != nil {
//				temp.PhoneOSName = &m.Name
//			}
//		}
//		if temp.Status != nil {
//			m := models.OrderStatusMap[*temp.Status]
//			temp.StatusName = &m
//		}
//		//if temp.BspStatus != nil {
//		//	m := models.OrderStatusMap[*temp.BspStatus]
//		//	temp.BspStatusName = &m
//		//}
//
//		if param.BlackSwitch == nil || *param.BlackSwitch == 0 {
//			continue
//		}
//
//		if temp.IsBacklist == nil && temp.Phone != nil {
//			flag, err := new(models.BlacklistPhone).ExistByPhone(*temp.Phone)
//			if err == nil {
//
//			}
//			if flag {
//				temp.IsBacklist = new(int)
//				*temp.IsBacklist = 1
//				continue
//			}
//		}
//		if temp.IsBacklist == nil && temp.IdCard != nil {
//			flag, err := new(models.BlacklistIdcard).ExistByIdCard(*temp.IdCard)
//			if err == nil {
//
//			}
//			if flag {
//				temp.IsBacklist = new(int)
//				*temp.IsBacklist = 1
//				continue
//			}
//		}
//		if !areaBlackFlag {
//			areaBlackFlag = true
//			areaBlack, err = new(models.BlacklistArea).GetFromCache("")
//			if err != nil {
//
//			}
//		}
//
//		if areaBlack == nil {
//			continue
//		}
//		if temp.ProvinceCode != nil {
//			_, ok := areaBlack.ProviceM[*temp.ProvinceCode]
//			if ok {
//				temp.IsBacklist = new(int)
//				*temp.IsBacklist = 1
//				continue
//			}
//		}
//		if temp.CityCode != nil {
//			_, ok := areaBlack.CityM[*temp.CityCode]
//			if ok {
//				temp.IsBacklist = new(int)
//				*temp.IsBacklist = 1
//				continue
//			}
//		}
//		if temp.AreaCode != nil {
//			_, ok := areaBlack.AreaM[*temp.AreaCode]
//			if ok {
//				temp.IsBacklist = new(int)
//				*temp.IsBacklist = 1
//				continue
//			}
//		}
//
//	}
//	this.Response(ctx, results)
//}

func (this *CardOrder) BkExport(ctx iris.Context) {
	param := new(api.BkCardOrderList)

	err := Tools.ShouldBindJSON(ctx, param)
	if err != nil {
		this.ExceptionSerive(ctx, apibackend.BASERR_INVALID_PARAMETER.Code(), apibackend.BASERR_INVALID_PARAMETER.Desc())
		ZapLog().Error("param err", zap.Error(err))
		return
	}

	//param.Status = this.pathToOrderStatus(ctx)

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
	//extend := ctx.Values().GetString("extend")
	//if len(extend) > 0 {
	//	arr := strings.Split(extend, ",")
	//	condPair = append(condPair, &models.SqlPairCondition{"class_tp in (?)", arr})
	//}
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
			//ZapLog().Info("cardclass 1")
			cc,err := new(models.CardClass).GetByIdFromCache(*temp.ClassTp)
			//ZapLog().Info("cardclass 2", zap.Error(err), zap.Any("cc",cc))
			if err == nil && cc != nil{
				temp.ClassName = cc.Detail
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
		//if temp.BspStatus != nil {
		//	m := models.OrderStatusMap[*temp.BspStatus]
		//	temp.BspStatusName = &m
		//}

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

func (this *CardOrder) BkFileCreate(ctx iris.Context) {
	param := new(api.BkCardOrderList)

	err := Tools.ShouldBindJSON(ctx, param)
	if err != nil {
		this.ExceptionSerive(ctx, apibackend.BASERR_INVALID_PARAMETER.Code(), apibackend.BASERR_INVALID_PARAMETER.Desc())
		ZapLog().Error("param err", zap.Error(err))
		return
	}

	rd := common.RandomDigit(5)
	timeStr := time.Now().Format("2006-01-02-15-04-05")
	fileName := timeStr+"-"+rd +".xlsx"
	excel,err := common.NewExcel("sheet1", config.GConfig.Server.FilePath + "/" +fileName)
	if err != nil {
		this.ExceptionSerive(ctx, apibackend.BASERR_DATABASE_ERROR.Code(), apibackend.BASERR_DATABASE_ERROR.Desc())
		ZapLog().Error("create file err", zap.Error(err))
		return
	}
	_,err = db.GRedis.GetConn().Set("filecreate_"+fileName, "", time.Duration(time.Second*1800)).Result()
	if err != nil {
		ZapLog().With(zap.Error(err)).Error("redis set err")
		this.ExceptionSerive(ctx, apibackend.BASERR_DATABASE_ERROR.Code(), apibackend.BASERR_DATABASE_ERROR.Desc())

		return
	}
	this.Response(ctx, fileName)

	go func() {
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

		headers :=[] string{
			"序号","订单号","套餐名","订单状态","姓名","身份证","手机号","省","市","区县","镇街道","详细地址","新手机号","ICCID","归属地","快递","快递单号","发货时间","照片上传","黑名单","下单时间",
		}

		if err := excel.AddHeader(headers); err != nil {
			ZapLog().Error("excel AddHeader err", zap.Error(err))
			return
		}
		total := param.Size
		newSize := int64(10)
		AllPage := total/newSize
		condStr := ""
		for page:= int64(0); page < AllPage ;page ++ {
			results, err := new(models.CardOrder).BkParseList(param).GetsWithConds(newSize, nil, condPair, condStr)
			if err != nil {
				ZapLog().With(zap.Error(err)).Error("Verify err")
				this.ExceptionSerive(ctx, apibackend.BASERR_DATABASE_ERROR.Code(), apibackend.BASERR_DATABASE_ERROR.Desc())
				return
			}

			coArr := &results
			minId := int64(0)
			for i := 0; i < len(*coArr); i++ {
				temp := (*coArr)[i]
				if i == len(*coArr) - 1 {
					minId = *temp.Id
				}
				appendData := make([]string, len(headers), len(headers))
				if temp.Id != nil {
					appendData[0] = fmt.Sprintf("%d", *temp.Id)
				}
				if temp.OrderNo != nil {
					appendData[1] = *temp.OrderNo
				}

				if temp.ClassTp != nil {
					//ZapLog().Info("cardclass 1")
					cc,err := new(models.CardClass).GetByIdFromCache(*temp.ClassTp)
					//ZapLog().Info("cardclass 2", zap.Error(err), zap.Any("cc",cc))
					if err == nil && cc != nil{
						temp.ClassName = cc.Detail
					}
				}
				if temp.ClassName != nil {
					appendData[2] = *temp.ClassName
				}

				if temp.Status != nil {
					m := models.OrderStatusMap[*temp.Status]
					temp.StatusName = &m
					if temp.StatusName != nil {
						appendData[3] = *temp.StatusName
					}
				}
				if temp.TrueName != nil {
					appendData[4] = *temp.TrueName
				}
				if temp.IdCard != nil {
					appendData[5] = *temp.IdCard
				}
				if temp.Phone != nil {
					appendData[6] = *temp.Phone
				}
				if temp.Province != nil {
					appendData[7] = *temp.Province
				}
				if temp.City != nil {
					appendData[8] = *temp.City
				}
				if temp.Area != nil {
					appendData[9] = *temp.Area
				}
				if temp.Town != nil {
					appendData[10] = *temp.Town
				}
				if temp.Address != nil {
					appendData[11] = *temp.Address
				}
				if temp.NewPhone != nil {
					appendData[12] = *temp.NewPhone
				}
				if temp.ICCID != nil {
					appendData[13] = *temp.ICCID
				}
				if temp.Guishudi != nil {
					appendData[14] = *temp.Guishudi
				}

				if temp.Express != nil {
					appendData[15] = *temp.Express
				}
				if temp.ExpressNo != nil {
					appendData[16] = *temp.ExpressNo
				}
				if temp.DeliverAt != nil {
					str := time.Unix(*temp.DeliverAt, 0).Format("2006/01/02_15-04-05")
					appendData[17] = str
				}
				if temp.IdCardPicFlag != nil {
					str := "否"
					if *temp.IdCardPicFlag == 1 {
						str = "是"
					}
					appendData[18] = str
				}
				if temp.CreatedAt != nil {
					str := time.Unix(*temp.CreatedAt, 0).Format("2006/01/02_15-04-05")
					appendData[20] = str
				}

				if param.BlackSwitch != nil && *param.BlackSwitch == 1 {
					this.CheckBlack(temp)
				}
				if temp.IsBacklist != nil && *temp.IsBacklist == 1 {
					appendData[19] = "是"
				}
				if err := excel.AppendCache(appendData); err != nil {
					ZapLog().With(zap.Error(err)).Error("excel.AppendCache err")
					return
				}
				if page %10 == 9 {
					if err:= excel.Sync(); err != nil {
						ZapLog().With(zap.Error(err)).Error("excel.Sync err")
						return
					}
				}
			}
			if err:= excel.Sync(); err != nil {
				ZapLog().With(zap.Error(err)).Error("excel.Sync err")
				return
			}
			condStr = fmt.Sprintf("id < %d", minId)
		}
		url := "http://file.lfcxwifi.com" + "/" + fileName
		_,err := db.GRedis.GetConn().Set("filecreate_"+fileName, url, time.Duration(time.Second*1800)).Result()
		if err != nil {
			ZapLog().With(zap.Error(err)).Error("redis set err")
		}
	}()
}

func (this *CardOrder) CheckBlack(temp *models.CardOrder)  {
	var areaBlack *models.BlacklistAreaCacheStt
	areaBlackFlag := false
	if temp.IsBacklist == nil && temp.Phone != nil {
		flag, err := new(models.BlacklistPhone).ExistByPhone(*temp.Phone)
		if err == nil {

		}
		if flag {
			temp.IsBacklist = new(int)
			*temp.IsBacklist = 1
			return
		}
	}
	if temp.IsBacklist == nil && temp.IdCard != nil {
		flag, err := new(models.BlacklistIdcard).ExistByIdCard(*temp.IdCard)
		if err == nil {

		}
		if flag {
			temp.IsBacklist = new(int)
			*temp.IsBacklist = 1
			return
		}
	}
	if !areaBlackFlag {
		areaBlackFlag = true
		var err error
		areaBlack, err = new(models.BlacklistArea).GetFromCache("")
		if err != nil {

		}
	}

	if areaBlack == nil {
		return
	}
	if temp.ProvinceCode != nil {
		_, ok := areaBlack.ProviceM[*temp.ProvinceCode]
		if ok {
			temp.IsBacklist = new(int)
			*temp.IsBacklist = 1
			return
		}
	}
	if temp.CityCode != nil {
		_, ok := areaBlack.CityM[*temp.CityCode]
		if ok {
			temp.IsBacklist = new(int)
			*temp.IsBacklist = 1
			return
		}
	}
	if temp.AreaCode != nil {
		_, ok := areaBlack.AreaM[*temp.AreaCode]
		if ok {
			temp.IsBacklist = new(int)
			*temp.IsBacklist = 1
			return
		}
	}

}

func (this *CardOrder) BkFileGet(ctx iris.Context) {
	fileName := ctx.URLParam("filename")

	if len(fileName) <= 0 {
		this.ExceptionSerive(ctx, apibackend.BASERR_INVALID_PARAMETER.Code(), apibackend.BASERR_INVALID_PARAMETER.Desc())
		ZapLog().Error("param err")
		return
	}

	url,err := db.GRedis.GetConn().Get("filecreate_"+fileName).Result()
	if err == redis.Nil {
		this.ExceptionSerive(ctx, apibackend.BASERR_OBJECT_NOT_FOUND.Code(), apibackend.BASERR_OBJECT_NOT_FOUND.Desc())
		ZapLog().Error("nofind")
		return
	}
	if err != nil {
		ZapLog().With(zap.Error(err)).Error("redis set err")
		this.ExceptionSerive(ctx, apibackend.BASERR_DATABASE_ERROR.Code(), apibackend.BASERR_DATABASE_ERROR.Desc())
		return
	}

	this.Response(ctx, url)
}