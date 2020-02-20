package models

import (
	"LianFaPhone/lfp-marketing-api/api"
	"github.com/jinzhu/gorm"
)

import (
	"LianFaPhone/lfp-marketing-api/common"
	"LianFaPhone/lfp-marketing-api/db"

)

type CardClassBigTp struct{
	Id    *int   `json:"id,omitempty"        gorm:"column:id;primary_key;AUTO_INCREMENT:1;not null"` //加上type:int(11)后AUTO_INCREMENT无效
	Isp *int `json:"isp,omitempty"     gorm:"column:isp;type:int(11)"` //拼音首字母缩写
	Name *string `json:"name,omitempty"     gorm:"column:name;type:varchar(20)"` //拼音首字母缩写

	Detail *string `json:"detail,omitempty"     gorm:"column:detail;type:varchar(50)"` //拼音首字母缩写

	Table
}

func (this *CardClassBigTp) TableName() string {
	return "card_class_big_tp"
}

func (this * CardClassBigTp) ParseAdd(p *api.BkCardClassBigTpAdd) *CardClassBigTp {
	cc := &CardClassBigTp{
		Isp: p.ISP,
		Name: p.Name,
		Detail: p.Detail,

	}
	cc.Valid = new(int)
	*cc.Valid = 1
	return  cc
}

func (this * CardClassBigTp) Parse(p *api.BkCardClassBigTp) *CardClassBigTp {
	d :=  &CardClassBigTp{
		Id: p.Id,

		Detail: p.Detail,

		Isp: p.ISP,
		Name: p.Name,

	}
	d.Valid = p.Valid
	return d
}

func (this * CardClassBigTp) ParseList(p *api.BkCardClassBigTpList) *CardClassBigTp {
	d:= &CardClassBigTp{
		Isp: p.ISP,
		Name: p.Name,
	}

	d.Valid = p.Valid
	return d
}

func (this *CardClassBigTp) Get() (*CardClassBigTp, error) {
	err := db.GDbMgr.Get().Model(this).Where("id = ? ", this.Id).Last(this).Error
	if err != nil {
		return nil,err
	}
	return this,nil
}


func (this *CardClassBigTp) Gets() ([]*CardClassBigTp, error) {
	var arr []*CardClassBigTp
	err := db.GDbMgr.Get().Model(this).Find(&arr).Error
	if err == gorm.ErrRecordNotFound {
		return nil,nil
	}
	if err != nil {
		return nil,err
	}
	return arr,nil
}



//批量导入
func (this *CardClassBigTp) Add() (*CardClassBigTp, error) {
	err := db.GDbMgr.Get().Create(this).Error
	if err != nil {
		return nil, err
	}
	err = db.GDbMgr.Get().Model(this).Where("id = ? ", this.Id).Last(this).Error
	if err != nil {
		return nil,err
	}
	return this,nil
}

func (this *CardClassBigTp) Update() (*CardClassBigTp, error) {
	err := db.GDbMgr.Get().Model(this).Where("id = ? ", this.Id).Updates(this).Error
	if err != nil {
		return nil,err
	}
	err = db.GDbMgr.Get().Model(this).Where("id = ? ", this.Id).Last(this).Error
	if err != nil {
		return nil,err
	}
	return this,nil
}

func (this *CardClassBigTp) ListWithConds(page, size int64, needFields []string, condPair []*SqlPairCondition) (*common.Result, error) {
	var list []*CardClassBigTp
	query := db.GDbMgr.Get().Where(this)

	for i := 0; i < len(condPair); i++ {
		if condPair[i] == nil {
			continue
		}
		query = query.Where(condPair[i].Key, condPair[i].Value)
	}

	query = query.Order("valid desc").Order("id desc")

	if len(needFields) > 0 {
		query = query.Select(needFields)
	}

	return new(common.Result).PageQuery(query, &CardClassBigTp{}, &list, page, size, nil, "")
}