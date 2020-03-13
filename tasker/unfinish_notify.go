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

func (this *Tasker) ydjthkFailNotify(idRecordName string) {
	defer models.PanicPrint()

	recoder, err := new(models.IdRecorder).GetByName(idRecordName)
	if err != nil {
		ZapLog().Error("IdRecorder GetByName err", zap.Error(err))
		return
	}

	if recoder == nil {
		if recoder, err = new(models.IdRecorder).Add(idRecordName, 0); err != nil {
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
			&models.SqlPairCondition{"third_order_at > ?", startId},
			&models.SqlPairCondition{"third_order_at <= ?", nowTime-5*3600},
			&models.SqlPairCondition{"third_order_at >= ?", nowTime-25*3600},
			&models.SqlPairCondition{"status = ?", models.CONST_OrderStatus_Fail_Retry},
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
			if *orderArr[i].ThirdOrderAt > startId {
				startId = *orderArr[i].ThirdOrderAt
			}

			//fail_retry 直接放过
			//if orderArr[i].Status != nil && *orderArr[i].Status ==  models.CONST_OrderStatus_Fail{
			//	logArr, err := new(models.CardOrderLog).GetsByOrderNoWithConds(*orderArr[i].OrderNo,10, []*models.SqlPairCondition{&models.SqlPairCondition{"created_at >= ?", nowTime - 2*24*3600}})
			//	if err !=nil {
			//		ZapLog().Error("GetsByOrderNoWithConds err", zap.Error(err))
			//	}
			//
			//	if len(logArr) <=0 {
			//		continue
			//	}
			//	oaoFlag := false
			//	for m:=0; m < len(logArr);m++ {
			//		orderlog := logArr[m]
			//		if orderlog.Log == nil || len(*orderlog.Log) <= 0{
			//			continue
			//		}
			//		if strings.Contains(*orderlog.Log, "OAO") {
			//			oaoFlag = true
			//			break
			//		}
			//	}
			//
			//	if !oaoFlag {
			//		continue
			//	}
			//}
			log := "短信发送：再次下单提醒, 成功"
			if err := sdk.GNotifySdk.SendSms(nil, *orderArr[i].Phone, "yd_jthk_fail",0, &platformTp); err != nil {
				ZapLog().Error("GNotifySdk.SendSms err", zap.Error(err))
				//continue
				log = "短信发送：再次下单提醒, 失败;"+err.Error()
			}

			new(models.CardOrderLog).FtParseAdd(nil, orderArr[i].OrderNo, &log).Add()


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
func (this *Tasker) ydjthkNewUnFinishNotify(idRecordName string) {
	defer models.PanicPrint()

	recoder, err := new(models.IdRecorder).GetByName(idRecordName)
	if err != nil {
		ZapLog().Error("IdRecorder GetByName err", zap.Error(err))
		return
	}

	if recoder == nil {
		if recoder, err = new(models.IdRecorder).Add(idRecordName, 0); err != nil {
			ZapLog().Error("IdRecorder Add  err", zap.Error(err))
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
			&models.SqlPairCondition{"third_order_at > ?", startId},
			&models.SqlPairCondition{"third_order_at <= ?", nowTime-16*60},
			&models.SqlPairCondition{"third_order_at >= ?", nowTime-35*60},
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
			if *orderArr[i].ThirdOrderAt > startId {
				startId = *orderArr[i].ThirdOrderAt
			}

			orderUrl,err := new(models.CardOrderUrl).GetByOrderNo(*orderArr[i].OrderNo)
			if err != nil {
				ZapLog().Error("CardOrderUrl GetByOrderNo err", zap.Error(err))
				log:= "短信发送：照片上传提醒，未找到Url信息"
				new(models.CardOrderLog).FtParseAdd(nil, orderArr[i].OrderNo, &log).Add()
				continue
			}
			if orderUrl == nil || orderUrl.Url == nil || len(*orderUrl.Url) < 3 {
				log:= "短信发送：照片上传提醒，未找到Url信息"
				new(models.CardOrderLog).FtParseAdd(nil, orderArr[i].OrderNo, &log).Add()
				continue
			}
			if !strings.HasPrefix(*orderUrl.Url, " ") {
				*orderUrl.Url = " "+*orderUrl.Url+" " //可生成短链
			}
			*orderUrl.Url = strings.Replace(*orderUrl.Url, "%", "%25", -1)

			log:= "短信发送：照片上传提醒，发送成功"

			if err := sdk.GNotifySdk.SendSms([]string{*orderUrl.Url}, *orderArr[i].Phone, "yd_jthk_new_unfinish",0, &platformTp); err != nil {
				ZapLog().Error("GNotifySdk.SendSms err", zap.Error(err))
				//continue
				log= "短信发送：照片上传提醒，发送失败;"+err.Error()
			}
			new(models.CardOrderLog).FtParseAdd(nil, orderArr[i].OrderNo, &log).Add()


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
