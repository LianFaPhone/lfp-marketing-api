package tasker

import (
	"fmt"
)

import (
	. "LianFaPhone/lfp-base/log/zap"
	"LianFaPhone/lfp-marketing-api/models"
	"LianFaPhone/lfp-marketing-api/db"
	"go.uber.org/zap"
	"time"
)

//多节点的时候考虑并发问题，或者加个开关，只让一个服务计算
func (this *Tasker) sheetWork(busyFlag bool) {
	defer models.PanicPrint()

	ZapLog().Info("start to sheetWork")

	idRecord, err := new(models.IdRecorder).GetByName("card_order_sheet")
	if err != nil {
		ZapLog().Error("IdRecorder GetByName err", zap.Error(err))
		return
	}
	if idRecord == nil {
		return
	}

	startId := int64(0)
	if idRecord.IdTag != nil {
		startId = *idRecord.IdTag
	}

	//todayUt := time.Now()
	//dateStr := todayUt.In(time.FixedZone("UTC", 8*3600)).Format("2006/01/02")
	//
	//lastUt := time.Unix(todayUt.Unix() -14*3600, 0)
	//lastDateStr := lastUt.In(time.FixedZone("UTC", 8*3600)).Format("2006/01/02")

	needFileds := []string{"id", "order_no", "class_isp", "class_big_tp", "class_tp", "province", "city", "created_at"}
	for true {
		conds := []*models.SqlPairCondition{
			&models.SqlPairCondition{"id > ?", startId},
		}

		orders, err := new(models.CardOrder).GetLimitByCond(10, conds, needFileds)
		if err != nil {
			ZapLog().Error("CardOrder GetLimitByCond err", zap.Error(err))
			return
		}
		if orders == nil || len(orders) <= 0 {
			return
		}

		dateMap, AreaMap, maxId := this.genSheetMap(orders)

		if maxId > 0 {
			flag := this.StoreSheet(dateMap, AreaMap, maxId, idRecord.Id)
			if !flag {
				return
			}
			startId = maxId
		}

		if len(orders) < 10 {
			break
		}
		if busyFlag {
			time.Sleep(time.Second * 10)
		}else{
			time.Sleep(time.Second * 5)
		}

	}

	return
}

func (this *Tasker) StoreSheet(dateMap map[string]*models.CardClasssheet, AreaMap map[string]*models.CardAreasheet, maxId, idRecordId int64) bool {
	tx := db.GDbMgr.Get().Begin()
	if err := new(models.IdRecorder).TxUpdate(tx, idRecordId, maxId); err != nil {
		tx.Rollback()
		ZapLog().Error("IdRecorder TxUpdate err", zap.Error(err))
		return false
	}
	for _, v := range dateMap {
		sheet, err := new(models.CardClasssheet).TxGetByDateAndTp(tx, *v.Date, v.ClassTp)
		if err != nil {
			ZapLog().Error("CardDatesheet GetByDate err", zap.Error(err))
			tx.Rollback()
			return false
		}
		if sheet == nil || sheet.Id == nil {
			err = v.TxAdd(tx)
			if err != nil {
				ZapLog().Error("saveSheet err", zap.Error(err))
				tx.Rollback()
				return false
			}
		} else {
			v.Id = sheet.Id
			if sheet.OrderCount == nil {
				sheet.OrderCount = new(int64)
				*sheet.OrderCount = 0
			}
			if v.OrderCount == nil {
				v.OrderCount = new(int64)
				*v.OrderCount = 0
			}
			*v.OrderCount += *sheet.OrderCount
			err = v.TxUpdate(tx)
			if err != nil {
				ZapLog().Error("saveSheet err", zap.Error(err))
				tx.Rollback()
				return false
			}
		}
	}

	for _, v := range AreaMap {
		sheet, err := new(models.CardAreasheet).TxGetByConds(tx, *v.DateAt, v.Province, v.City, new(int), new(int))
		if err != nil {
			ZapLog().Error("CardDatesheet GetByDate err", zap.Error(err))
			tx.Rollback()
			return false
		}
		if sheet == nil || sheet.Id == nil {
			err = v.TxAdd(tx)
			if err != nil {
				ZapLog().Error("saveSheet err", zap.Error(err))
				tx.Rollback()
				return false
			}
		} else {
			v.Id = sheet.Id
			if sheet.OrderCount == nil {
				sheet.OrderCount = new(int64)
				*sheet.OrderCount = 0
			}
			if v.OrderCount == nil {
				v.OrderCount = new(int64)
				*v.OrderCount = 0
			}
			*v.OrderCount += *sheet.OrderCount
			err = v.TxUpdate(tx)
			if err != nil {
				ZapLog().Error("saveSheet err", zap.Error(err))
				tx.Rollback()
				return false
			}
		}
	}

	//
	tx.Commit()
	return true
}

func (this *Tasker) genSheetMap(orders []*models.CardOrder) (map[string]*models.CardClasssheet, map[string]*models.CardAreasheet, int64) {
	maxId := int64(0)
	classMap := make(map[string]*models.CardClasssheet)
	AreaMap := make(map[string]*models.CardAreasheet)
	for j := 0; j < len(orders); j++ {
		if *orders[j].Id > maxId {
			maxId = *orders[j].Id
		}
		if orders[j].CreatedAt == nil {
			continue
		}
		dateStr := time.Unix(*orders[j].CreatedAt, 0).In(time.FixedZone("UTC", 8*3600)).Format("2006/01/02")
		dateInt := GenDay2(*orders[j].CreatedAt)
		//dateInt := GenDay2(*orders[j].CreatedAt)

		classsheet, ok := classMap[dateStr]
		if !ok {
			classsheet = &models.CardClasssheet{
				Date:       &dateStr,
				OrderCount: new(int64),
			}
			*classsheet.OrderCount = 0
			classMap[dateStr] = classsheet
		}
		if classsheet.OrderCount == nil {
			classsheet.OrderCount = new(int64)
			*classsheet.OrderCount = 0
		}
		*classsheet.OrderCount = *classsheet.OrderCount + 1

		if orders[j].ClassTp != nil {
			classsheet, ok = classMap[fmt.Sprintf("%s_%d", dateStr, *orders[j].ClassTp)]
			if !ok {
				classsheet = &models.CardClasssheet{
					Date:       &dateStr,
					ClassTp:    orders[j].ClassTp,
					OrderCount: new(int64),
				}
				*classsheet.OrderCount = 0
				classMap[fmt.Sprintf("%s_%d", dateStr, *orders[j].ClassTp)] = classsheet
			}
			if classsheet.OrderCount == nil {
				classsheet.OrderCount = new(int64)
				*classsheet.OrderCount = 0
			}
			*classsheet.OrderCount = *classsheet.OrderCount + 1
			//tps, ok := models.ClassTpMap[*orders[j].ClassTp]
			//if ok {
			//	classsheet.ClassISP = &tps.ISP
			//}
		}

		///////////////////////

		if orders[j].Province == nil {
			continue
		}

		//省，总
		key := ""
		//key := fmt.Sprintf("%d_%s", dateInt, *orders[j].Province)
		//areasheet,ok  := AreaMap[key]
		//if !ok {
		//	areasheet = &models.CardAreasheet{
		//		DateAt: &dateInt,
		//		Province: orders[j].Province,
		//		City: new(string),
		//		ClassTp: new(int),
		//		ClassISP: new(int),
		//		OrderCount: new(int64),
		//	}
		//	*areasheet.ClassTp = 0
		//	*areasheet.ClassISP = 0
		//	*areasheet.OrderCount = 0
		//	AreaMap[key] = areasheet
		//}
		//if areasheet.OrderCount == nil {
		//	areasheet.OrderCount = new(int64)
		//	*areasheet.OrderCount = 0
		//}
		//*areasheet.OrderCount = *areasheet.OrderCount + 1

		//省，运营商
		//if orders[j].ClassIsp != nil {
		//	key = fmt.Sprintf("%d_%s_%d", dateInt, *orders[j].Province, *orders[j].ClassIsp)
		//	areasheet,ok  = AreaMap[key]
		//	if !ok {
		//		areasheet = &models.CardAreasheet{
		//			DateAt: &dateInt,
		//			Province: orders[j].Province,
		//			City: new(string),
		//			ClassTp: new(int),
		//			ClassISP: orders[j].ClassIsp,
		//			OrderCount: new(int64),
		//		}
		//		*areasheet.ClassTp = 0
		//		*areasheet.OrderCount = 0
		//		AreaMap[key] = areasheet
		//	}
		//	if areasheet.OrderCount == nil {
		//		areasheet.OrderCount = new(int64)
		//		*areasheet.OrderCount = 0
		//	}
		//	*areasheet.OrderCount = *areasheet.OrderCount + 1
		//}

		//省，运营商,套餐
		//if orders[j].ClassIsp != nil && orders[j].ClassTp != nil {
		//	key = fmt.Sprintf("%d_%s_%d_%d", dateInt, *orders[j].Province, *orders[j].ClassIsp, *orders[j].ClassTp)
		//	areasheet,ok  := AreaMap[key]
		//	if !ok {
		//		areasheet = &models.CardAreasheet{
		//			DateAt: &dateInt,
		//			Province: orders[j].Province,
		//			City: new(string),
		//			ClassTp: orders[j].ClassTp,
		//			ClassISP: orders[j].ClassIsp,
		//			OrderCount: new(int64),
		//		}
		//		*areasheet.OrderCount = 0
		//		AreaMap[key] = areasheet
		//	}
		//	if areasheet.OrderCount == nil {
		//		areasheet.OrderCount = new(int64)
		//		*areasheet.OrderCount = 0
		//	}
		//	*areasheet.OrderCount = *areasheet.OrderCount + 1
		//}

		if orders[j].City == nil {
			continue
		}

		//省,市， 运营商
		//if orders[j].ClassIsp != nil {
		//	key = fmt.Sprintf("%d_%s_%s_%d", dateInt, *orders[j].Province,*orders[j].City, *orders[j].ClassIsp)
		//	areasheet,ok  := AreaMap[key]
		//	if !ok {
		//		areasheet = &models.CardAreasheet{
		//			DateAt: &dateInt,
		//			Province: orders[j].Province,
		//			City: orders[j].City,
		//			ClassTp: new(int),
		//			ClassISP: orders[j].ClassIsp,
		//			OrderCount: new(int64),
		//		}
		//		*areasheet.ClassTp = 0
		//		*areasheet.OrderCount = 0
		//		AreaMap[key] = areasheet
		//	}
		//	if areasheet.OrderCount == nil {
		//		areasheet.OrderCount = new(int64)
		//		*areasheet.OrderCount = 0
		//	}
		//	*areasheet.OrderCount = *areasheet.OrderCount + 1
		//}

		//省，市，运营商,套餐
		if orders[j].ClassIsp != nil && orders[j].ClassTp != nil {
			key = fmt.Sprintf("%d_%s_%s_%d_%d", dateInt, *orders[j].Province, *orders[j].City, *orders[j].ClassIsp, *orders[j].ClassTp)
			areasheet, ok := AreaMap[key]
			if !ok {
				areasheet = &models.CardAreasheet{
					DateAt:     &dateInt,
					Province:   orders[j].Province,
					City:       orders[j].City,
					//ClassTp:    orders[j].ClassTp,
					//ClassISP:   orders[j].ClassIsp,
					OrderCount: new(int64),
				}
				*areasheet.OrderCount = 0
				AreaMap[key] = areasheet
			}
			if areasheet.OrderCount == nil {
				areasheet.OrderCount = new(int64)
				*areasheet.OrderCount = 0
			}
			*areasheet.OrderCount = *areasheet.OrderCount + 1
		}
	}

	return classMap, AreaMap, maxId
}


