package tasker

import (
	. "LianFaPhone/lfp-base/log/zap"
	"LianFaPhone/lfp-marketing-api/models"
	"github.com/lieney/idCardParser-go"
	"go.uber.org/zap"
	"time"
)

func (this *Tasker) ipsWork(busyFlag bool) {
	defer models.PanicPrint()

	recoder, err := new(models.IdRecorder).GetByName("card_order_ips")
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
			if orderArr[i] == nil {
				continue
			}
			if *orderArr[i].Id > startId {
				startId = *orderArr[i].Id
			}

			mp := &models.CardOrder{
				Id: orderArr[i].Id,
			}
			this.setIps(orderArr[i], mp)
			this.setIdCardAudit(orderArr[i], mp)
			this.setAreaCode(orderArr[i], mp)

			if err = mp.Update(); err != nil {
				ZapLog().Error("CardOrder Update err", zap.Error(err))
				return
			}

			if busyFlag {
				time.Sleep(time.Second * 5)
			}else{
				time.Sleep(time.Second * 1)
			}
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
		time.Sleep(time.Second * 2)
	}

}

func (this *Tasker) setIps(inPut, outPut *models.CardOrder) {
	if inPut.IP == nil || len(*inPut.IP) == 0 || inPut.Id == nil || inPut.Ips != nil {
		return
	}

	oldOrder, err := new(models.CardOrder).GetByIp(*inPut.IP, []*models.SqlPairCondition{&models.SqlPairCondition{"id < ?", inPut.Id}})
	if err != nil {
		ZapLog().Error("CardOrder GetByIp err", zap.Error(err))
		return
	}

	ips := 0
	if oldOrder == nil {

	} else {
		if oldOrder.Ips != nil {
			ips = *oldOrder.Ips
		}
		ips = ips + 1
	}
	outPut.Ips = &ips
}

func (this *Tasker) setIdCardAudit(inPut, outPut *models.CardOrder) {
	if inPut.IdCardAudit == nil || *inPut.IdCardAudit == 0 {
		if inPut.IdCard != nil {
			if flag := idCardParser.IsValidate(*inPut.IdCard); !flag {
				outPut.IdCardAudit = new(int)
				*outPut.IdCardAudit = 2
			}
		}
	}
}

func (this *Tasker) setAreaCode(input, outPut *models.CardOrder) {
	if input == nil || outPut == nil {
		return
	}
	if input.ProvinceCode == nil && input.Province != nil {
		prv, err := new(models.BsProvice).GetByName(*input.Province)
		if err != nil {
			ZapLog().Error("BsProvice GetByName err", zap.Error(err))
		} else if prv != nil {
			outPut.ProvinceCode = prv.Code
		}
	}

	if input.CityCode == nil && input.City != nil {
		prv, err := new(models.BsCity).GetByName(*input.City)
		if err != nil {
			ZapLog().Error("BsCity GetByName err", zap.Error(err))
		} else if prv != nil {
			outPut.CityCode = prv.Code
		}
	}

	if input.AreaCode == nil && input.Area != nil {
		prv, err := new(models.BsArea).GetByName(*input.Area)
		if err != nil {
			ZapLog().Error("BsArea GetByName err", zap.Error(err))
		} else if prv != nil {
			outPut.AreaCode = prv.Code
		}
	}
	if input.Province == nil && input.ProvinceCode != nil {
		prv, err := new(models.BsProvice).GetByCode(*input.ProvinceCode)
		if err != nil {
			ZapLog().Error("BsProvice GetByCode err", zap.Error(err))
		} else if prv != nil {
			outPut.Province = prv.Name
		}
	}

	if input.City == nil && input.CityCode != nil {
		prv, err := new(models.BsCity).GetByCode(*input.CityCode)
		if err != nil {
			ZapLog().Error("BsCity GetByCode err", zap.Error(err))
		} else if prv != nil {
			outPut.City = prv.Name
		}
	}

	if input.Area == nil && input.AreaCode != nil {
		prv, err := new(models.BsArea).GetByCode(*input.AreaCode)
		if err != nil {
			ZapLog().Error("BsArea GetByName err", zap.Error(err))
		} else if prv != nil {
			outPut.Area = prv.Name
		}
	}
}
