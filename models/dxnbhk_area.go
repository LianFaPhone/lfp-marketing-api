package models

import (
	"LianFaPhone/lfp-marketing-api/db"
	"github.com/jinzhu/gorm"
)

type DxnbhkArea struct {
	Id        *int    `json:"id,omitempty"        gorm:"column:id;primary_key;AUTO_INCREMENT:1;not null"`
	CityId   *int    `json:"city_id,omitempty"        gorm:"column:city_id;type:int(11);index"`

	//	Code      *string `json:"code,omitempty"     gorm:"column:code;type:varchar(20)"`
//	CityCode  *string `json:"city_code,omitempty"     gorm:"column:city_code;type:varchar(20);index"`
	Name      *string `json:"name,omitempty"     gorm:"column:name;type:varchar(20)"`
	ShortName string  `json:"short_name,omitempty"     gorm:"column:short_name;type:varchar(20)"`
//	Lng       *string `json:"lng,omitempty"     gorm:"column:lng;type:varchar(20)"`
//	Lat       *string `json:"lat,omitempty"     gorm:"column:lat;type:varchar(20)"`
	Sort      int     `json:"sort,omitempty"     gorm:"column:Sort;type:int(11)"`
	Table
}

func (this *DxnbhkArea) TableName() string {
	return "dxnbhk_area"
}

func (this *DxnbhkArea) GetsByCity(code string) ([]*DxnbhkArea, error) {
	var arr []*DxnbhkArea
	err := db.GDbMgr.Get().Model(this).Where("city_code = ? and valid = 1", code).Order("sort").Find(&arr).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return arr, err
}

func (this *DxnbhkArea) GetByCityIdAndLikeName(cid int, name string) (*DxnbhkArea, error) {
	p := new(DxnbhkArea)
	err := db.GDbMgr.Get().Model(this).Where("city_id = ?", cid).Where("name like ? or short_name like ?", name+"%", name+"%").Last(p).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return p, err
}

func (this *DxnbhkArea) GetByCityIdAndName(cid int, name string) (*DxnbhkArea, error) {
	p := new(DxnbhkArea)
	err := db.GDbMgr.Get().Model(this).Where("city_id = ?", cid).Where("name = ? or short_name = ?", name, name).Last(p).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return p, err
}

func (this *DxnbhkArea) GetByName(name string) (*DxnbhkArea, error) {
	p := new(DxnbhkArea)
	err := db.GDbMgr.Get().Model(this).Where("name = ? or short_name = ?", name, name).Last(p).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return p, err
}

func (this *DxnbhkArea) GetByCode(code string) (*DxnbhkArea, error) {
	p := new(DxnbhkArea)
	err := db.GDbMgr.Get().Model(this).Where("code = ?", code).Last(p).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return p, err
}
