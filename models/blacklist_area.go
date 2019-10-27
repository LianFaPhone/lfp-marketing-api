package models

import (
	"LianFaPhone/lfp-marketing-api/api"
	"LianFaPhone/lfp-marketing-api/common"
	"LianFaPhone/lfp-marketing-api/db"
	"github.com/jinzhu/gorm"
	"github.com/juju/errors"
	"math/rand"
	"time"
)

type BlacklistArea struct {
	Id           *int64  `json:"id,omitempty"        gorm:"column:id;primary_key;AUTO_INCREMENT:1;not null"` //加上type:int(11)后AUTO_INCREMENT无效
	Province     *string `json:"province,omitempty"     gorm:"column:province;type:varchar(15)"`             //省份
	ProvinceCode *string `json:"province_code,omitempty"     gorm:"column:province_code;type:varchar(10)"`   //省份
	City         *string `json:"city,omitempty"     gorm:"column:city;type:varchar(15)"`                     //城市
	CityCode     *string `json:"city_code,omitempty"     gorm:"column:city_code;type:varchar(10)"`           //城市
	Area         *string `json:"area,omitempty"     gorm:"column:area;type:varchar(20)"`                     //区
	AreaCode     *string `json:"area_code,omitempty"     gorm:"column:area_code;type:varchar(20)"`           //区
	Table
}

type BlacklistAreaCacheStt struct {
	ProviceM map[string]bool
	CityM    map[string]bool
	AreaM    map[string]bool
}

func (this *BlacklistArea) TableName() string {
	return "blacklist_area"
}

func (this *BlacklistArea) ParseAdd(p *api.BkBlacklistAreaAdd) *BlacklistArea {
	this.Province = p.Province
	this.ProvinceCode = p.ProvinceCode
	this.City = p.City
	this.CityCode = p.CityCode
	this.Area = p.Area
	this.AreaCode = p.AreaCode
	this.Valid = new(int)
	*this.Valid = 1
	return this
}

func (this *BlacklistArea) Parse(p *api.BkBlacklistArea) *BlacklistArea {
	this.Id = p.Id
	this.Province = p.Province
	this.ProvinceCode = p.ProvinceCode
	this.City = p.City
	this.CityCode = p.CityCode
	this.Area = p.Area
	this.AreaCode = p.AreaCode
	this.Valid = p.Valid
	return this
}

func (this *BlacklistArea) ParseList(p *api.BkBlacklistAreaList) *BlacklistArea {
	this.Id = p.Id
	this.Province = p.Province
	this.ProvinceCode = p.ProvinceCode
	this.City = p.City
	this.CityCode = p.CityCode
	this.Area = p.Area
	this.AreaCode = p.AreaCode
	this.Valid = p.Valid
	return this
}

func (this *BlacklistArea) Add() error {
	err := db.GDbMgr.Get().Create(this).Error
	if err != nil {
		return err
	}
	return nil
}

func (this *BlacklistArea) Update() error {
	err := db.GDbMgr.Get().Model(this).Where("id = ? ", this.Id).Updates(this).Error
	if err != nil {
		return err
	}
	return nil
}

func (this *BlacklistArea) GetFromCache(uuid string) (*BlacklistAreaCacheStt, error) {
	data, err := db.GCache.GetBlacklistArea(uuid)
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, nil
	}

	acty, ok := data.(*BlacklistAreaCacheStt)
	if !ok {
		return nil, errors.Annotate(err, "type err")
	}
	return acty, nil
}

func (this *BlacklistArea) InnerGetBy(input interface{}) (interface{}, *time.Duration, error) {
	expire := time.Second * time.Duration(3600*5+rand.Intn(600))

	var list []BlacklistArea
	err := db.GDbMgr.Get().Model(this).Where("valid = 1").Find(&list).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil, err
	}
	if err != nil {
		return nil, nil, err
	}
	if list == nil {
		return nil, nil, gorm.ErrRecordNotFound // nil,nil,nil可能将是永远不超时
	}
	res := &BlacklistAreaCacheStt{
		ProviceM: make(map[string]bool),
		CityM:    make(map[string]bool),
		AreaM:    make(map[string]bool),
	}

	for i := 0; i < len(list); i++ {
		if list[i].AreaCode != nil {
			res.AreaM[*list[i].AreaCode] = true
			continue
		}
		if list[i].CityCode != nil {
			res.CityM[*list[i].CityCode] = true
			continue
		}
		if list[i].ProvinceCode != nil {
			res.ProviceM[*list[i].ProvinceCode] = true
			continue
		}
	}

	return res, &expire, nil
}

//func (this * BlacklistArea) CheckExistProvice(name string) (bool,error){
//	count := 0
//	err := db.GDbMgr.Get().Model(this).Where("province = ? ", name).Count(&count).Error
//	if err != nil {
//		return  false,err
//	}
//	return count > 0,nil
//}

func (this *BlacklistArea) Del() error {
	err := db.GDbMgr.Get().Model(this).Where("id = ? ", this.Id).Delete(this).Error
	if err != nil {
		return err
	}
	return nil
}

func (this *BlacklistArea) ListWithConds(page, size int64, needFields []string, condPair []*SqlPairCondition) (*common.Result, error) {
	var list []*BlacklistArea
	query := db.GDbMgr.Get().Where(this)

	for i := 0; i < len(condPair); i++ {
		if condPair[i] == nil {
			continue
		}
		query = query.Where(condPair[i].Key, condPair[i].Value)
	}
	if len(needFields) > 0 {
		query = query.Select(needFields)
	}

	query = query.Order("valid desc").Order("id desc")

	return new(common.Result).PageQuery(query, &BlacklistArea{}, &list, page, size, nil, "")
}
