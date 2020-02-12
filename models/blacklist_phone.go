package models

import (
	"LianFaPhone/lfp-marketing-api/api"
	"LianFaPhone/lfp-marketing-api/common"
	"LianFaPhone/lfp-marketing-api/db"
)

type BlacklistPhone struct {
	Id    *int64  `json:"id,omitempty"        gorm:"column:id;primary_key;AUTO_INCREMENT:1;not null"` //加上type:int(11)后AUTO_INCREMENT无效
	Phone *string `json:"phone,omitempty"     gorm:"column:phone;type:varchar(15)"`                   //订单号
	Table
}

func (this *BlacklistPhone) TableName() string {
	return "blacklist_phone"
}

func (this *BlacklistPhone) ParseAdd(p *api.BkBlacklistPhoneAdd) *BlacklistPhone {
	this.Phone = p.Phone
	this.Valid = new(int)
	*this.Valid = 1
	return this
}

func (this *BlacklistPhone) Parse(p *api.BkBlacklistPhone) *BlacklistPhone {
	this.Id = p.Id
	this.Phone = p.Phone
	this.Valid = p.Valid
	return this
}

func (this *BlacklistPhone) ParseList(p *api.BkBlacklistPhoneList) *BlacklistPhone {
	this.Id = p.Id
	this.Phone = p.Phone
	this.Valid = p.Valid
	return this
}

func (this *BlacklistPhone) Add() error {
	err := db.GDbMgr.Get().Create(this).Error
	if err != nil {
		return err
	}
	return nil
}

func (this *BlacklistPhone) Update() error {
	err := db.GDbMgr.Get().Model(this).Where("id = ? ", this.Id).Updates(this).Error
	if err != nil {
		return err
	}
	return nil
}

func (this *BlacklistPhone) Del() error {
	err := db.GDbMgr.Get().Model(this).Where("id = ? ", this.Id).Delete(this).Error
	if err != nil {
		return err
	}
	return nil
}

func (this *BlacklistPhone) ExistByPhone(phone string) (bool, error) {
	count := 0
	err := db.GDbMgr.Get().Model(this).Where("phone = ? and valid = 1", phone).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (this *BlacklistPhone) ListWithConds(page, size int64, needFields []string, condPair []*SqlPairCondition) (*common.Result, error) {
	var list []*BlacklistPhone
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

	return new(common.Result).PageQuery(query, &BlacklistPhone{}, &list, page, size, nil, "")
}
