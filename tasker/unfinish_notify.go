package tasker

import (
	"LianFaPhone/lfp-marketing-api/models"
	"LianFaPhone/lfp-marketing-api/sdk"
	"go.uber.org/zap"
	"strings"
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
			&models.SqlPairCondition{"created_at >= ?", nowTime-10*3600},
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
			if orderArr[i] == nil || orderArr[i].OrderNo == nil || orderArr[i].Phone == nil{
				continue
			}
			if *orderArr[i].Id > startId {
				startId = *orderArr[i].Id
			}
			logArr, err := new(models.CardOrderLog).GetsByOrderNoWithConds(*orderArr[i].OrderNo,10, []*models.SqlPairCondition{&models.SqlPairCondition{"created_at >= ?", nowTime - 2*24*3600}})
			if err !=nil {
				ZapLog().Error("GetsByOrderNoWithConds err", zap.Error(err))
			}

			if len(logArr) <=0 {
				continue
			}
			oaoFlag := false
			for m:=0; m < len(logArr);m++ {
				orderlog := logArr[m]
				if orderlog.Log == nil || len(*orderlog.Log) <= 0{
					continue
				}
				if strings.Contains(*orderlog.Log, "OAO") {
					oaoFlag = true
					break
				}
			}

			if !oaoFlag {
				continue
			}

			if err := sdk.GNotifySdk.SendSms(nil, *orderArr[i].Phone, "yd_jthk_fail",0, &platformTp); err != nil {
				ZapLog().Error("GNotifySdk.SendSms err", zap.Error(err))
				continue
			}

			time.Sleep(time.Millisecond * 1000)
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
		time.Sleep(time.Second * 30)
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
	platformTp :=2 //chuanglan

	for true {
		conds := []*models.SqlPairCondition{
			&models.SqlPairCondition{"id > ?", startId},
			&models.SqlPairCondition{"created_at <= ?", nowTime-20*60},
			&models.SqlPairCondition{"created_at >= ?", nowTime-35*60},
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
			if orderArr[i] == nil || orderArr[i].OrderNo == nil || orderArr[i].Phone == nil{
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
			if !strings.HasPrefix(*orderUrl.Url, " ") {
				*orderUrl.Url = " "+*orderUrl.Url+" " //可生成短链
			}

			if err := sdk.GNotifySdk.SendSms([]string{*orderUrl.Url}, *orderArr[i].Phone, "yd_jthk_new_unfinish",0, &platformTp); err != nil {
				ZapLog().Error("GNotifySdk.SendSms err", zap.Error(err))
				continue
			}

			//time.Sleep(time.Millisecond * 1)
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
		//time.Sleep(time.Second * 1)
	}

}
