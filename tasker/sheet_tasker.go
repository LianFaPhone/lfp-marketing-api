package tasker

import (
	"time"
	"LianFaPhone/lfp-marketing-api/models"
	. "LianFaPhone/lfp-base/log/zap"
	"go.uber.org/zap"
)

/*
日期统计 每1小时一次，每次查下昨天的是否计算完，然后再算今天的.先全量计算，以后增量计算
"2006-01-02 15:04:05"
 */

var GTasker Tasker
type Tasker struct{

}

func (this *Tasker) Init() error {

	return nil
}

func (this *Tasker) Start() {
//	go this.run()
}

func (this *Tasker) run() {
	dateSheetTicker := time.NewTicker(time.Hour *2)
	ipsTicker := time.NewTicker(time.Minute *10)
	activeCodeTicker := time.NewTicker(time.Hour *5)


	for {
		select {
			case <-dateSheetTicker.C:
				this.datesheetWork()
			case <-ipsTicker.C:
				this.ipsWork()
			case <- activeCodeTicker.C:
				this.activeCodeWork()

		}
	}
}

func (this *Tasker) activeCodeWork() {
	t := time.Now().Unix() - 3600 * 24 * 7
	conds := []*models.SqlPairCondition{
		&models.SqlPairCondition{"created_at <= ?", t},
	}

	if err := new(models.ActiveCode).Del(conds); err != nil {
		return
	}
}

func (this *Tasker) ipsWork() {
	recoder,err := new(models.IdRecorder).GetByName("IPS")
	if err != nil {

		return
	}

	if recoder == nil {
		if recoder,err = new(models.IdRecorder).Add("IPS", 0); err != nil{
			return
		}
	}

	startId := recoder.IdTag

	conds := []*models.SqlPairCondition{
		&models.SqlPairCondition{"id > ?", startId},
	}
	orderArr,err := new(models.CardOrder).GetLimitByCond(10, conds)
	if err != nil {
		return
	}
	if orderArr == nil || len(orderArr) <=0 {
		return
	}

	//记录id, 倒叙
	for i:= len(orderArr)-1; i >= 0; i++ {
		if orderArr[i] == nil {
			continue
		}
		if orderArr[i].IP == nil || len(*orderArr[i].IP) == 0 {
			continue
		}
		oldOrder, err := new(models.CardOrder).GetByIp(*orderArr[i].IP,  []*models.SqlPairCondition{&models.SqlPairCondition{"id <= ?", startId}})
		if err != nil {
			return
		}
		if oldOrder == nil {
			continue
		}
		ips := 0
		if oldOrder.Ips != nil {
			ips = *oldOrder.Ips
		}
		ips = ips + 1

		mp := models.CardOrder{
			Id: orderArr[i].Id,
			Ips: &ips,
		}
		if err = mp.Update(); err != nil {
			return
		}
		startId = *orderArr[i].Id
	}

	err = new(models.IdRecorder).Update(recoder.Id, startId)
	if err != nil {
		return
	}

}

//多节点的时候考虑并发问题，或者加个开关，只让一个服务计算
func (this *Tasker) datesheetWork() {
	ZapLog().Info("start to datesheetWork")
	todayUt := time.Now()
	dateStr := todayUt.In(time.FixedZone("UTC", 8*3600)).Format("2006/01/02")

	lastUt := time.Unix(todayUt.Unix() -14*3600, 0)
	lastDateStr := lastUt.In(time.FixedZone("UTC", 8*3600)).Format("2006/01/02")

	if dateStr != lastDateStr {
		//计算你昨日的
		for i:=0; i < len(models.ClassTpArr); i++ {
			sheet,err :=  new(models.CardDatesheet).GetByDateAndTp(lastDateStr, models.ClassTpArr[i].Tp)
			if err != nil {
				ZapLog().Error("CardDatesheet GetByDate err", zap.Error(err))
			}else{
				if sheet == nil {
					err := this.saveSheet(GenDay(lastUt), GenDay(todayUt), todayUt.Unix(), lastDateStr, 1, nil, &models.ClassTpArr[i].Tp)
					if err != nil {
						ZapLog().Error("saveSheet err", zap.Error(err))
					}
				}else if  (sheet.EndFlag != nil && *sheet.EndFlag == 0){
					err := this.saveSheet(GenDay(lastUt), GenDay(todayUt), todayUt.Unix(), lastDateStr, 1, sheet.Id, &models.ClassTpArr[i].Tp)
					if err != nil {
						ZapLog().Error("saveSheet err", zap.Error(err))
					}
					//计算
					//设置标志
				}
			}
		}

		sheet,err :=  new(models.CardDatesheet).GetByDate(lastDateStr)
		if err != nil {
			ZapLog().Error("CardDatesheet GetByDate err", zap.Error(err))
		}else{
			if sheet == nil {
				err := this.saveSheet(GenDay(lastUt), GenDay(todayUt), todayUt.Unix(), lastDateStr, 1, nil, nil)
				if err != nil {
					ZapLog().Error("saveSheet err", zap.Error(err))
				}
			}else if  (sheet.EndFlag != nil && *sheet.EndFlag == 0){
				err := this.saveSheet(GenDay(lastUt), GenDay(todayUt), todayUt.Unix(), lastDateStr, 1, sheet.Id, nil)
				if err != nil {
					ZapLog().Error("saveSheet err", zap.Error(err))
				}
				//计算
				//设置标志
			}
		}

	}

	//计算当天的
	for i:=0; i < len(models.ClassTpArr); i++ {
		sheet,err :=  new(models.CardDatesheet).GetByDateAndTp(lastDateStr, models.ClassTpArr[i].Tp)
		if err != nil {
			ZapLog().Error("CardDatesheet GetByDate err", zap.Error(err))
		}else{
			if sheet == nil {
				err := this.saveSheet(GenDay(todayUt), todayUt.Unix(), todayUt.Unix(), dateStr, 0, nil, &models.ClassTpArr[i].Tp)
				if err != nil {
					ZapLog().Error("saveSheet err", zap.Error(err))
				}
			}else{
				err := this.saveSheet(GenDay(todayUt), todayUt.Unix(), todayUt.Unix(), dateStr, 0, sheet.Id, &models.ClassTpArr[i].Tp)
				if err != nil {
					ZapLog().Error("saveSheet err", zap.Error(err))
				}
			}
		}
	}

	sheet,err :=  new(models.CardDatesheet).GetByDate(lastDateStr)
	if err != nil {
		ZapLog().Error("CardDatesheet GetByDate err", zap.Error(err))
	}else{
		if sheet == nil {
			err := this.saveSheet(GenDay(todayUt), todayUt.Unix(), todayUt.Unix(), dateStr, 0, nil, nil)
			if err != nil {
				ZapLog().Error("saveSheet err", zap.Error(err))
			}
		}else{
			err := this.saveSheet(GenDay(todayUt), todayUt.Unix(), todayUt.Unix(), dateStr, 0, sheet.Id, nil)
			if err != nil {
				ZapLog().Error("saveSheet err", zap.Error(err))
			}
		}
	}


	return
}

func (this *Tasker) saveSheet(startOrder, endOrder, LastAt int64, dateStr string, EndFlag int, id* int64, classTp *int) error {
	conds := []*models.SqlPairCondition{
		&models.SqlPairCondition{ "created_at >= ?", startOrder},
		&models.SqlPairCondition{ "created_at < ?", endOrder},
	}
	if classTp != nil {
		conds = append(conds, &models.SqlPairCondition{ "class_tp = ?", classTp})
	}
	count,err := new(models.CardOrder).CountByConds(conds)
	if err != nil {
		return err
	}
	p := &models.CardDatesheet{
		Id : id,
		Date: &dateStr,
		OrderCount:&count,
		ClassTp: classTp,
		LastAt: &LastAt,
		EndFlag: &EndFlag,
	}
	p.Valid = new(int)
	*p.Valid = 1
	if p.Id == nil {
		return p.Add()
	}else{
		p.Date = nil //不是主键更新会报错
		return p.Update()
	}
	return nil
}

func GenDay(ut time.Time) int64{
	t1:=ut.Year()        //年
	t2:=ut.Month()       //月
	t3:=ut.Day()         //日
	loc := time.FixedZone("UTC", 8*3600)
	currentTimeData:=time.Date(t1,t2,t3,0,0,0,0,loc)
	return currentTimeData.Unix()
}

