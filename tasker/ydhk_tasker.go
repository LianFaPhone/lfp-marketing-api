package tasker

import (
	. "LianFaPhone/lfp-base/log/zap"
	"LianFaPhone/lfp-marketing-api/models"
	"LianFaPhone/lfp-marketing-api/thirdcard-api/ydhk"
	"fmt"
	"go.uber.org/zap"
	"time"
)

func (this *Tasker) ydhkWork() {
	defer models.PanicPrint()

	recoder, err := new(models.IdRecorder).GetByName("ydhk_oao")
	if err != nil {
		ZapLog().Error("IdRecorder GetByName err", zap.Error(err))
		return
	}

	if recoder == nil {
		if recoder, err = new(models.IdRecorder).Add("card_order_ips", 0); err != nil {
			ZapLog().Error("IdRecorder Add card_order_ips err", zap.Error(err))
			return
		}
	}

	startId := int64(0)
	if recoder.IdTag != nil {
		startId = *recoder.IdTag
	}

	for true {
		conds := []*models.SqlPairCondition{
			&models.SqlPairCondition{"id > ?", startId},
			&models.SqlPairCondition{"class_big_tp = ?", 5},
			&models.SqlPairCondition{"created_at >= ?", time.Now().Unix() - 15*60},
			//条件还得处理下
		}

		orderArr, err := new(models.CardOrder).GetLimitByCond(10, conds)
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
			if orderArr[i] == nil {
				continue
			}
			if *orderArr[i].Id > startId {
				startId = *orderArr[i].Id
			}
			if orderArr[i].Phone == nil || orderArr[i].NewPhone == nil || orderArr[i].IdCard == nil {
				continue
			}

			mp := &models.CardOrder{
				Id: orderArr[i].Id,
			}

			yidongArr,err := new(ydhk.ReOrderSerach).Send(*orderArr[i].Phone, *orderArr[i].IdCard)
			if err != nil {
				
			}

			oaoFlag := false
			for j:=0; j< len(yidongArr); j++ {
				if yidongArr[j].Number == nil {
					continue
				}
				if * yidongArr[j].Number == *orderArr[i].NewPhone {
					oaoFlag = true
					break
				}
			}

			fmt.Println(oaoFlag)

			if err = mp.Update(); err != nil {
				ZapLog().Error("CardOrder Update err", zap.Error(err))
				return
			}

			time.Sleep(time.Second * 1)
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


