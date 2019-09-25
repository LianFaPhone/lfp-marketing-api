package models

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"LianFaPhone/lfp-base/log/zap"
	"LianFaPhone/lfp-marketing-api/config"
	"go.uber.org/zap"
	"LianFaPhone/lfp-marketing-api/db"
	"fmt"
)

func InitDbTable() {
	log.ZapLog().Info("start InitDbTable")
	if !config.GConfig.Db.Debug {
		log.ZapLog().Info("end InitDbTable" )
		fmt.Println(config.GConfig.Db.Debug)
		return
	}
	err := db.GDbMgr.Get().Set("gorm:table_options", "ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 AUTO_INCREMENT=1;").AutoMigrate(&CardOrder{}, &CardDatesheet{}).Error
	if err != nil {
		log.ZapLog().Error("AutoMigrate err", zap.Error(err))
	}
	err = db.GDbMgr.Get().Set("gorm:table_options", "ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 AUTO_INCREMENT=1;").AutoMigrate( &IdRecorder{}, &ActiveCode{}, &BlacklistIdcard{}).Error
	if err != nil {
		log.ZapLog().Error("AutoMigrate err", zap.Error(err))
	}
	err = db.GDbMgr.Get().Set("gorm:table_options", "ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 AUTO_INCREMENT=1;").AutoMigrate( &BlacklistPhone{}, &BlacklistArea{}).Error
	if err != nil {
		log.ZapLog().Error("AutoMigrate err", zap.Error(err))
	}
	err = db.GDbMgr.Get().Set("gorm:table_options", "ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 AUTO_INCREMENT=1;").AutoMigrate( &BsStreet{}).Error
	if err != nil {
		log.ZapLog().Error("AutoMigrate err", zap.Error(err))
	}
	err = db.GDbMgr.Get().Set("gorm:table_options", "ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 AUTO_INCREMENT=1;").AutoMigrate(&PhoneNumberLevel{}, &PhoneNumberPool{}).Error
	if err != nil {
		log.ZapLog().Error("AutoMigrate err", zap.Error(err))
	}


	log.ZapLog().Info("end InitDbTable")
}
