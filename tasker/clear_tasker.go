package tasker

import (
	. "LianFaPhone/lfp-base/log/zap"
	"LianFaPhone/lfp-marketing-api/models"
	"go.uber.org/zap"
	"time"
)

//card_order_log //清理半年前的日志，半年前的日志基本没用了

func (this *Tasker) clearWork() {
	defer models.PanicPrint()

	ZapLog().Info("start to clearWork")

	//idRecord, err := new(models.IdRecorder).GetByName("card_order_log_clear")
	//if err != nil {
	//	ZapLog().Error("IdRecorder GetByName err", zap.Error(err))
	//	return
	//}
	//if idRecord == nil {
	//	return
	//}
	//
	//startId := int64(0)
	//if idRecord.IdTag != nil {
	//	startId = *idRecord.IdTag
	//}

	beginTime := time.Now().Unix() - 1 * 30 * 24* 3600
	conds := []*models.SqlPairCondition{
		//&models.SqlPairCondition{"id > ?", startId},
		&models.SqlPairCondition{"created_at <= ?", beginTime},
	}

	for i:=0; i<1000;i++ {

		new(models.CardOrderUrl).DelWithConds([]*models.SqlPairCondition{&models.SqlPairCondition{"created_at <= ?", time.Now().Unix() - 2* 3600}}, 10)

		delCount, err := new(models.CardOrderLog).DelWithConds(conds, 10)
		if err != nil {
			ZapLog().Error("CardOrderLog DelWithConds err", zap.Error(err))
			return
		}
		if  delCount <= 10 {
			break
		}

		time.Sleep(time.Second * 15)
	}

	return
}
