package models

import (
	"LianFaPhone/lfp-marketing-api/db"
	"github.com/jinzhu/gorm"
)

type BsCity struct {
	Id           *int    `json:"id,omitempty"        gorm:"column:id;primary_key;AUTO_INCREMENT:1;not null"`
	Code         *string `json:"code,omitempty"     gorm:"column:code;type:varchar(20)"`
	ProvinceCode *string `json:"province_code,omitempty"     gorm:"column:province_code;type:varchar(20);index"`
	Name         *string `json:"name,omitempty"     gorm:"column:name;type:varchar(20)"`
	ShortName    *string `json:"short_name,omitempty"     gorm:"column:short_name;type:varchar(20)"`
	Lng          *string `json:"lng,omitempty"     gorm:"column:lng;type:varchar(20)"`
	Lat          *string `json:"lat,omitempty"     gorm:"column:lat;type:varchar(20)"`
	Sort         *int    `json:"sort,omitempty"     gorm:"column:sort;type:int(11)"`
	Table
}

func (this *BsCity) TableName() string {
	return "bs_city"
}

func (this *BsCity) GetsByProvince(code string) ([]*BsCity, error) {
	var arr []*BsCity
	err := db.GDbMgr.Get().Model(this).Where("province_code = ? and valid = 1", code).Order("sort").Find(&arr).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return arr, err
}

func (this *BsCity) GetByName(name string) (*BsCity, error) {
	p := new(BsCity)
	err := db.GDbMgr.Get().Model(this).Where("name = ? or short_name = ?", name, name).Last(p).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return p, err
}

func (this *BsCity) GetByCode(code string) (*BsCity, error) {
	p := new(BsCity)
	err := db.GDbMgr.Get().Model(this).Where("code = ?", code).Last(p).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return p, err
}
