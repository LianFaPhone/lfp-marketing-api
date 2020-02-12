package models

import (
	"LianFaPhone/lfp-marketing-api/db"
	"github.com/jinzhu/gorm"
)

type BsStreet struct {
	STREET_ID   int    `json:"id,omitempty"        gorm:"column:STREET_ID;primary_key;AUTO_INCREMENT:1;not null"`
	STREET_CODE string `json:"street_code,omitempty"     gorm:"column:STREET_CODE;type:varchar(20)"`
	AREA_CODE   string `json:"area_code,omitempty"     gorm:"column:AREA_CODE;type:varchar(20);index"`
	STREET_NAME string `json:"street_name,omitempty"     gorm:"column:STREET_NAME;type:varchar(20)"`
	SHORT_NAME  string `json:"short_code,omitempty"     gorm:"column:SHORT_NAME;type:varchar(20)"`
	SORT        int    `json:"sort,omitempty"     gorm:"column:SORT;type:int(11)"`
	LNG         string `json:"lng,omitempty"     gorm:"column:LNG;type:varchar(20)"`
	LAT         string `json:"lat,omitempty"     gorm:"column:LAT;type:varchar(20)"`
	DATA_STATE  int    `json:"valid,omitempty"     gorm:"column:DATA_STATE;type:int(11)"`
}

func (this *BsStreet) TableName() string {
	return "bs_street"
}

func (this *BsStreet) GetsByArea(code string) ([]*BsStreet, error) {
	var arr []*BsStreet
	err := db.GDbMgr.Get().Model(this).Where("AREA_CODE = ?", code).Order("SORT").Find(&arr).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return arr, err
}
