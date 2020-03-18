package models

import (
	"LianFaPhone/lfp-marketing-api/db"
	"github.com/jinzhu/gorm"
)

type DxnbhkProvice struct {
	Id        *int    `json:"id,omitempty"        gorm:"column:id;primary_key;AUTO_INCREMENT:1;not null"`
//	Code      *string `json:"code,omitempty"     gorm:"column:code;type:varchar(20)"`
	Name      *string `json:"name,omitempty"     gorm:"column:name;type:varchar(20)"`
	ShortName *string `json:"short_name,omitempty"     gorm:"column:short_name;type:varchar(20)"`
//	Lng       *string `json:"lng,omitempty"     gorm:"column:lng;type:varchar(20)"`
//	Lat       *string `json:"lat,omitempty"     gorm:"column:lat;type:varchar(20)"`
	Sort      *int    `json:"sort,omitempty"     gorm:"column:sort;type:int(11)"`
	Table
}

func (this *DxnbhkProvice) TableName() string {
	return "dxnbhk_province"
}

func (this *DxnbhkProvice) Gets() ([]*DxnbhkProvice, error) {
	var arr []*DxnbhkProvice
	err := db.GDbMgr.Get().Model(this).Where("valid = 1").Order("sort").Find(&arr).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return arr, err
}

func (this *DxnbhkProvice) Gets2() ([]*DxnbhkProvice, error) {
	var arr []*DxnbhkProvice
	err := db.GDbMgr.Get().Model(this).Where("valid = 1").Find(&arr).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return arr, err
}

func (this *DxnbhkProvice) GetByName(name string) (*DxnbhkProvice, error) {
	p := new(DxnbhkProvice)
	err := db.GDbMgr.Get().Model(this).Where("name = ? or short_name = ?", name, name).Last(p).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return p, err
}

func (this *DxnbhkProvice) GetByCode(code string) (*DxnbhkProvice, error) {
	p := new(DxnbhkProvice)
	err := db.GDbMgr.Get().Model(this).Where("code = ?", code).Last(p).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return p, err
}

