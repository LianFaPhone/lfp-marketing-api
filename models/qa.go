package models

import (
	"LianFaPhone/lfp-marketing-api/api"
	"LianFaPhone/lfp-marketing-api/common"
	"LianFaPhone/lfp-marketing-api/db"
)

type Qa struct{
	Id           *int64  `json:"id,omitempty"        gorm:"column:id;primary_key;AUTO_INCREMENT:1;not null"` //加上type:int(11)后AUTO_INCREMENT无效
	Title *string   `json:"title,omitempty"     gorm:"column:title;type:varchar(50)"`
	Content *string `json:"content,omitempty"     gorm:"column:content;type:text"`
	Table
}

func (this *Qa) TableName() string{
	return "qa"
}

func (this *Qa) ParseAdd(p *api.BkQaAdd) *Qa {
	this.Title = p.Title
	this.Content = p.Content
	this.Valid = new(int)
	*this.Valid = 1
	return this
}

func (this *Qa) Parse(p *api.BkQa) *Qa {
	this.Id = p.Id
	this.Title = p.Title
	this.Content = p.Content
	this.Valid = new(int)
	*this.Valid = 1
	return this
}

func (this *Qa) Add() error {
	err := db.GDbMgr.Get().Create(this).Error
	if err != nil {
		return err
	}
	return nil
}

func (this *Qa) Update()  error {
	err := db.GDbMgr.Get().Model(this).Where("id = ? ", this.Id).Updates(this).Error
	if err != nil {
		return err
	}
	return nil
}


func (this *Qa) ListWithConds(page, size int64,likeStr *string, needFields []string, condPair []*SqlPairCondition) (*common.Result, error) {
	var list []*Qa
	query := db.GDbMgr.Get().Where(this)

	if likeStr != nil {
		query = query.Where("title like ?", "%"+*likeStr+"%")
	}

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

	return new(common.Result).PageQuery(query, &Qa{}, &list, page, size, nil, "")
}