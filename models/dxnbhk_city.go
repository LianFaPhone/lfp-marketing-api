package models

import (
	"LianFaPhone/lfp-marketing-api/db"
	"github.com/jinzhu/gorm"
)

type DxnbhkCity struct {
	Id           *int    `json:"id,omitempty"        gorm:"column:id;primary_key;AUTO_INCREMENT:1;not null"`
	ProvinceId   *int    `json:"province_id,omitempty"        gorm:"column:province_id;type:int(11);index"`
	//	Code         *string `json:"code,omitempty"     gorm:"column:code;type:varchar(20)"`
//	ProvinceCode *string `json:"province_code,omitempty"     gorm:"column:province_code;type:varchar(20);index"`
	Name         *string `json:"name,omitempty"     gorm:"column:name;type:varchar(20)"`
	ShortName    *string `json:"short_name,omitempty"     gorm:"column:short_name;type:varchar(20)"`
//	Lng          *string `json:"lng,omitempty"     gorm:"column:lng;type:varchar(20)"`
//	Lat          *string `json:"lat,omitempty"     gorm:"column:lat;type:varchar(20)"`
	Sort         *int    `json:"sort,omitempty"     gorm:"column:sort;type:int(11)"`
	Table
}

func (this *DxnbhkCity) TableName() string {
	return "dxnbhk_city"
}

func (this *DxnbhkCity) GetByProvinceIdAndName(pid int, name string) (*DxnbhkCity, error) {
	p := new(DxnbhkCity)
	err := db.GDbMgr.Get().Model(this).Where("province_id = ?", pid).Where("name = ? or short_name = ?", name, name).Last(p).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return p, err
}

func (this *DxnbhkCity) GetsByProvince(code string) ([]*DxnbhkCity, error) {
	var arr []*DxnbhkCity
	err := db.GDbMgr.Get().Model(this).Where("province_code = ? and valid = 1", code).Order("sort").Find(&arr).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return arr, err
}

func (this *DxnbhkCity) GetByName(name string) (*DxnbhkCity, error) {
	p := new(DxnbhkCity)
	err := db.GDbMgr.Get().Model(this).Where("name = ? or short_name = ?", name, name).Last(p).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return p, err
}

func (this *DxnbhkCity) GetByCode(code string) (*DxnbhkCity, error) {
	p := new(DxnbhkCity)
	err := db.GDbMgr.Get().Model(this).Where("code = ?", code).Last(p).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return p, err
}

