package tasker

import (
	"LianFaPhone/lfp-marketing-api/models"
)

func (this *Tasker) orderHistory() {
	defer models.PanicPrint()

	//tt := time.Now().Unix() - 3600*24*30*12
	//condPair := []*models.SqlPairCondition{
	//	&models.SqlPairCondition{"created_at <= ?", tt},
	//}
	//
	//for j := 0; j < 200; j++ {
	//	orders, err := new(models.CardOrder).GetLimitByCond2(5, condPair)
	//	if err != nil {
	//		ZapLog().Error("CardOrder GetLimitByCond2 err", zap.Error(err))
	//		return
	//	}
	//
	//	for i := 0; i < len(orders); i++ {
	//		p := new(models.CardOrderHistory).BkParseAdd(orders[i])
	//		uFlag, err := p.Unique()
	//		if err != nil {
	//			ZapLog().Error("CardOrderHistory Unique err", zap.Error(err))
	//			return
	//		}
	//		if uFlag {
	//			if err = p.Add(); err != nil {
	//				ZapLog().Error("CardOrderHistory Add err", zap.Error(err))
	//				return
	//			}
	//		}
	//		if err := orders[i].Del(); err != nil {
	//			ZapLog().Error("CardOrder Del err", zap.Error(err))
	//			return
	//		}
	//		time.Sleep(time.Second * 1)
	//	}
	//
	//	if len(orders) < 10 {
	//		break
	//	}
	//
	//	time.Sleep(time.Second * 10)
	//}

}
