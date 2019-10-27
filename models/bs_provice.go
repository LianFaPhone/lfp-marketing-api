package models

import (
	"LianFaPhone/lfp-marketing-api/db"
	"github.com/jinzhu/gorm"
)

type BsProvice struct {
	Id        *int    `json:"id,omitempty"        gorm:"column:id;primary_key;AUTO_INCREMENT:1;not null"`
	Code      *string `json:"code,omitempty"     gorm:"column:code;type:varchar(20)"`
	Name      *string `json:"name,omitempty"     gorm:"column:name;type:varchar(20)"`
	ShortName *string `json:"short_name,omitempty"     gorm:"column:short_name;type:varchar(20)"`
	Lng       *string `json:"lng,omitempty"     gorm:"column:lng;type:varchar(20)"`
	Lat       *string `json:"lat,omitempty"     gorm:"column:lat;type:varchar(20)"`
	Sort      *int    `json:"sort,omitempty"     gorm:"column:sort;type:int(11)"`
	Table
}

func (this *BsProvice) TableName() string {
	return "bs_province"
}

func (this *BsProvice) Gets() ([]*BsProvice, error) {
	var arr []*BsProvice
	err := db.GDbMgr.Get().Model(this).Where("valid = 1").Order("sort").Find(&arr).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return arr, err
}

func (this *BsProvice) GetByName(name string) (*BsProvice, error) {
	p := new(BsProvice)
	err := db.GDbMgr.Get().Model(this).Where("name = ? or short_name = ?", name, name).Last(p).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return p, err
}

func (this *BsProvice) GetByCode(code string) (*BsProvice, error) {
	p := new(BsProvice)
	err := db.GDbMgr.Get().Model(this).Where("code = ?", code).Last(p).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return p, err
}
