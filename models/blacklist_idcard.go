package models

import(
	"LianFaPhone/lfp-marketing-api/db"
	"LianFaPhone/lfp-marketing-api/common"
	"LianFaPhone/lfp-marketing-api/api"
)

type BlacklistIdcard struct{
	Id         		*int64      `json:"id,omitempty"        gorm:"column:id;primary_key;AUTO_INCREMENT:1;not null"` //加上type:int(11)后AUTO_INCREMENT无效
	IdCard    	    *string     `json:"idcard,omitempty"     gorm:"column:idcard;type:varchar(30)"` //订单号
	Table
}


func (this * BlacklistIdcard) TableName() string{
	return "blacklist_idcard"
}

func (this * BlacklistIdcard) ParseAdd(p * api.BkBlacklistIdcardAdd) *BlacklistIdcard {
	this.IdCard = p.IdCard
	this.Valid = new(int)
	*this.Valid = 1
	return this
}

func (this * BlacklistIdcard) Parse(p * api.BkBlacklistIdcard) *BlacklistIdcard {
	this.Id = p.Id
	this.IdCard = p.IdCard
	this.Valid = p.Valid
	return this
}

func (this * BlacklistIdcard) ParseList(p * api.BkBlacklistIdcardList) *BlacklistIdcard {
	this.Id = p.Id
	this.IdCard = p.IdCard
	this.Valid = p.Valid
	return this
}

func (this * BlacklistIdcard) Add() error {
	err := db.GDbMgr.Get().Create(this).Error
	if err != nil {
		return  err
	}
	return nil
}

func (this * BlacklistIdcard) Update() error {
	err := db.GDbMgr.Get().Model(this).Where("id = ? ", this.Id).Updates(this).Error
	if err != nil {
		return  err
	}
	return nil
}

func (this * BlacklistIdcard) Del() error {
	err := db.GDbMgr.Get().Model(this).Where("id = ? ", this.Id).Delete(this).Error
	if err != nil {
		return  err
	}
	return nil
}

func (this *BlacklistIdcard) ListWithConds(page, size int64, needFields []string,  condPair []*SqlPairCondition) (*common.Result, error) {
	var list []*BlacklistIdcard
	query := db.GDbMgr.Get().Where(this)

	for i:=0; i < len(condPair);i++ {
		if condPair[i] == nil {
			continue
		}
		query = query.Where(condPair[i].Key, condPair[i].Value)
	}
	if len(needFields) > 0 {
		query = query.Select(needFields)
	}

	query = query.Order("valid desc").Order("id desc")

	return new(common.Result).PageQuery(query, &BlacklistIdcard{}, &list, page, size, nil, "")
}