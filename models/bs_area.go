package models

import (
	"LianFaPhone/lfp-marketing-api/db"
	"github.com/jinzhu/gorm"
)

type BsArea struct {
	Id        *int    `json:"id,omitempty"        gorm:"column:id;primary_key;AUTO_INCREMENT:1;not null"`
	Code      *string `json:"code,omitempty"     gorm:"column:code;type:varchar(20)"`
	CityCode  *string `json:"city_code,omitempty"     gorm:"column:city_code;type:varchar(20);index"`
	Name      *string `json:"name,omitempty"     gorm:"column:name;type:varchar(20)"`
	ShortName string  `json:"short_name,omitempty"     gorm:"column:short_name;type:varchar(20)"`
	Lng       *string `json:"lng,omitempty"     gorm:"column:lng;type:varchar(20)"`
	Lat       *string `json:"lat,omitempty"     gorm:"column:lat;type:varchar(20)"`
	Sort      int     `json:"sort,omitempty"     gorm:"column:Sort;type:int(11)"`
	Table
}

func (this *BsArea) TableName() string {
	return "bs_area"
}

func (this *BsArea) GetsByCity(code string) ([]*BsArea, error) {
	var arr []*BsArea
	err := db.GDbMgr.Get().Model(this).Where("city_code = ? and valid = 1", code).Order("sort").Find(&arr).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return arr, err
}

func (this *BsArea) GetByName(name string) (*BsArea, error) {
	p := new(BsArea)
	err := db.GDbMgr.Get().Model(this).Where("name = ? or short_name = ?", name, name).Last(p).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return p, err
}

func (this *BsArea) GetByCode(code string) (*BsArea, error) {
	p := new(BsArea)
	err := db.GDbMgr.Get().Model(this).Where("code = ?", code).Last(p).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return p, err
}
