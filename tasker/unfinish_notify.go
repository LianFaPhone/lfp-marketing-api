package tasker

import (
	"LianFaPhone/lfp-marketing-api/models"
	"LianFaPhone/lfp-marketing-api/sdk"
	"go.uber.org/zap"
	"time"
)

import (
	. "LianFaPhone/lfp-base/log/zap"
)

func (this *Tasker) jthkFailNotify() {
	defer models.PanicPrint()

	recoder, err := new(models.IdRecorder).GetByName("JthkFailNotify")
	if err != nil {
		ZapLog().Error("IdRecorder GetByName err", zap.Error(err))
		return
	}

	if recoder == nil {
		if recoder, err = new(models.IdRecorder).Add("JthkFailNotify", 0); err != nil {
			ZapLog().Error("IdRecorder Add JthkFailNotify err", zap.Error(err))
			return
		}
	}

	partnerIds := GetJtHkPartnerIds()
	if len(partnerIds) <= 0 {
		return
	}

	startId := int64(0)
	if recoder.IdTag != nil {
		startId = *recoder.IdTag
	}

	nowTime := time.Now().Unix()
	platformTp :=2

	for true {
		conds := []*models.SqlPairCondition{
			&models.SqlPairCondition{"id > ?", startId},
			&models.SqlPairCondition{"created_at <= ?", nowTime-5*3600},
			&models.SqlPairCondition{"created_at >= ?", nowTime-7*3600},
			&models.SqlPairCondition{"status = ?", models.CONST_OrderStatus_Fail},
		}

		if len(partnerIds) > 0 {
			conds = append(conds, &models.SqlPairCondition{"partner_id in (?)", partnerIds})
		}

		orderArr, err := new(models.CardOrder).GetLimitByCond(10, conds, nil)
		if err != nil {
			ZapLog().Error("CardOrder GetLimitByCond err", zap.Error(err))
			return
		}
		if orderArr == nil || len(orderArr) <= 0 {
			//fmt.Println("nil data")
			return
		}

		//记录id, 倒叙
		for i := len(orderArr) - 1; i >= 0; i-- {
			if orderArr[i] == nil || orderArr[i].OrderNo == nil{
				continue
			}
			if *orderArr[i].Id > startId {
				startId = *orderArr[i].Id
			}

			if err := sdk.GNotifySdk.SendSms(nil, *orderArr[i].Phone, "yd_jthk_fail",0, &platformTp); err != nil {
				ZapLog().Error("GNotifySdk.SendSms err", zap.Error(err))
				continue
			}

			time.Sleep(time.Millisecond * 1)
		}

		err = recoder.Update(startId)
		if err != nil {
			ZapLog().Error("IdRecorder Update err", zap.Error(err))
			return
		}
		//fmt.Println("startId= ", startId)
		if len(orderArr) < 10 {
			break
		}
	}

}

//新未完成订单短信通知
func (this *Tasker) jthkNewUnFinishNotify() {
	defer models.PanicPrint()

	recoder, err := new(models.IdRecorder).GetByName("JthkNewUnfinishNotify")
	if err != nil {
		ZapLog().Error("IdRecorder GetByName err", zap.Error(err))
		return
	}

	if recoder == nil {
		if recoder, err = new(models.IdRecorder).Add("JthkNewUnfinishNotify", 0); err != nil {
			ZapLog().Error("IdRecorder Add JthkNewUnfinishNotify err", zap.Error(err))
			return
		}
	}

	startId := int64(0)
	if recoder.IdTag != nil {
		startId = *recoder.IdTag
	}

	partnerIds := GetJtHkPartnerIds()
	if len(partnerIds) <= 0 {
		return
	}

	nowTime := time.Now().Unix()
	platformTp :=2

	for true {
		conds := []*models.SqlPairCondition{
			&models.SqlPairCondition{"id > ?", startId},
			&models.SqlPairCondition{"created_at <= ?", nowTime-15*60},
			&models.SqlPairCondition{"created_at >= ?", nowTime-25*60},
			&models.SqlPairCondition{"status = ?", models.CONST_OrderStatus_New_UnFinish},
		}
		if len(partnerIds) > 0 {
			conds = append(conds, &models.SqlPairCondition{"partner_id in (?)", partnerIds})
		}

		orderArr, err := new(models.CardOrder).GetLimitByCond(10, conds, nil)
		if err != nil {
			ZapLog().Error("CardOrder GetLimitByCond err", zap.Error(err))
			return
		}
		if orderArr == nil || len(orderArr) <= 0 {
			//fmt.Println("nil data")
			return
		}

		//记录id, 倒叙
		for i := len(orderArr) - 1; i >= 0; i-- {
			if orderArr[i] == nil || orderArr[i].OrderNo == nil {
				continue
			}
			if *orderArr[i].Id > startId {
				startId = *orderArr[i].Id
			}

			orderUrl,err := new(models.CardOrderUrl).GetByOrderNo(*orderArr[i].OrderNo)
			if err != nil {
				ZapLog().Error("CardOrderUrl GetByOrderNo err", zap.Error(err))
				continue
			}
			if orderUrl == nil || orderUrl.Url == nil || len(*orderUrl.Url) < 3 {
				continue
			}

			if err := sdk.GNotifySdk.SendSms([]string{*orderUrl.Url}, *orderArr[i].Phone, "yd_jthk_new_unfinish",0, &platformTp); err != nil {
				ZapLog().Error("GNotifySdk.SendSms err", zap.Error(err))
				continue
			}

			time.Sleep(time.Millisecond * 1)
		}

		err = recoder.Update(startId)
		if err != nil {
			ZapLog().Error("IdRecorder Update err", zap.Error(err))
			return
		}
		//fmt.Println("startId= ", startId)
		if len(orderArr) < 10 {
			break
		}
	}

}
