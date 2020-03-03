package models

import (
	"LianFaPhone/lfp-base/log/zap"
	"LianFaPhone/lfp-marketing-api/config"
	"LianFaPhone/lfp-marketing-api/db"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go.uber.org/zap"
)

func InitDbTable() {
	log.ZapLog().Info("start InitDbTable")
	if !config.GConfig.Db.Debug {
		log.ZapLog().Info("end InitDbTable")
		fmt.Println(config.GConfig.Db.Debug)
		return
	}
	err := db.GDbMgr.Get().Set("gorm:table_options", "ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 AUTO_INCREMENT=1;").AutoMigrate(&CardOrder{}, &CardDatesheet{}, &CardAreasheet{}, &PdPartnerGoods{}, &CardIdcardPic{}, &CardOrderLog{}).Error
	if err != nil {
		log.ZapLog().Error("AutoMigrate err", zap.Error(err))
	}
	err = db.GDbMgr.Get().Set("gorm:table_options", "ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 AUTO_INCREMENT=1;").AutoMigrate(&IdRecorder{}, &ActiveCode{}, &BlacklistIdcard{}, &CardOrderUrl{}).Error
	if err != nil {
		log.ZapLog().Error("AutoMigrate err", zap.Error(err))
	}
	err = db.GDbMgr.Get().Set("gorm:table_options", "ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 AUTO_INCREMENT=1;").AutoMigrate(&BlacklistPhone{}, &BlacklistArea{}, &BlacklistIdcard{}, &Qa{}).Error
	if err != nil {
		log.ZapLog().Error("AutoMigrate err", zap.Error(err))
	}
	err = db.GDbMgr.Get().Set("gorm:table_options", "ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 AUTO_INCREMENT=1;").AutoMigrate(&BsStreet{}, &CardOrderNotify{}, &IdRecorder{}).Error
	if err != nil {
		log.ZapLog().Error("AutoMigrate err", zap.Error(err))
	}
	err = db.GDbMgr.Get().Set("gorm:table_options", "ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 AUTO_INCREMENT=1;").AutoMigrate(&PhoneNumberLevel{}, &PhoneNumberPool{}).Error
	if err != nil {
		log.ZapLog().Error("AutoMigrate err", zap.Error(err))
	}
	err = db.GDbMgr.Get().Set("gorm:table_options", "ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 AUTO_INCREMENT=1;").AutoMigrate(&BsProvice{}, &BsCity{}, &BsArea{}, &PdPartner{}).Error
	if err != nil {
		log.ZapLog().Error("AutoMigrate err", zap.Error(err))
	}

	log.ZapLog().Info("end InitDbTable")
}
