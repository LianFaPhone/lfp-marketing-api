package models

import(
	"LianFaPhone/lfp-marketing-api/db"
	"LianFaPhone/lfp-marketing-api/common"
	"github.com/jinzhu/gorm"
	"LianFaPhone/lfp-marketing-api/api"
)

type PhoneNumberLevel struct{
	Id         		*int64      `json:"id,omitempty"        gorm:"column:id;primary_key;AUTO_INCREMENT:1;not null"` //加上type:int(11)后AUTO_INCREMENT无效
	Level           *int         `json:"level,omitempty"     gorm:"column:level;type:int(10)"`
	Desc            *string      `json:"desc,omitempty"     gorm:"column:desc;type:varchar(200)"` //描述
	Table
}

func (this * PhoneNumberLevel) TableName() string {
	return "phone_number_level"
}

func (this * PhoneNumberLevel) ParseAdd(p * api.BkPhoneNumberLevelAdd) *PhoneNumberLevel {
	np :=  &PhoneNumberLevel{
		Desc: p.Desc,
		Level: p.Level,
	}
	np.Valid = p.Valid
	if np.Valid == nil {
		np.Valid = new(int)
		*np.Valid = 1
	}

	return np
}

func (this * PhoneNumberLevel) ParseGets(p * api.BkPhoneNumberLevelGets) *PhoneNumberLevel {
	np :=  &PhoneNumberLevel{
		Desc: p.Desc,
		Level: p.Level,
	}
	np.Valid = p.Valid

	return np
}


func (this * PhoneNumberLevel) ParseList(p * api.BkPhoneNumberLevelList) *PhoneNumberLevel {
	np :=  &PhoneNumberLevel{
		Id: p.Id,
		Desc: p.Desc,
		Level: p.Level,
	}
	np.Valid = p.Valid
	return np
}

func (this * PhoneNumberLevel) ParseUpdate(p * api.BkPhoneNumberLevelUpdate) *PhoneNumberLevel {
	np :=  &PhoneNumberLevel{
		Id: p.Id,
		Desc: p.Desc,
		Level: p.Level,
	}
	np.Valid = p.Valid

	return np
}

func (this * PhoneNumberLevel) Add() error {
	err := db.GDbMgr.Get().Create(this).Error
	if err != nil  {
		return err
	}
	return nil
}

func (this * PhoneNumberLevel) UniqueByLevel(level int) (bool, error) {
	cc := 0
	err := db.GDbMgr.Get().Model(this).Where(" level = ?", level).Count(&cc).Error
	if err != nil  {
		return false,err
	}
	return cc == 0, nil
}

func (this * PhoneNumberLevel) Update() error {
	err := db.GDbMgr.Get().Model(this).Where("id = ? ", this.Id).Updates(this).Error
	if err != nil {
		return  err
	}
	return nil
}

func (this *PhoneNumberLevel) ListWithConds(page, size int64, needFields []string,  condPair []*SqlPairCondition) (*common.Result, error) {
	var list []*PhoneNumberLevel
	query := db.GDbMgr.Get().Where(this)

	for i:=0; i < len(condPair);i++ {
		if condPair[i] == nil {
			continue
		}
		query = query.Where(condPair[i].Key, condPair[i].Value)
	}

	query = query.Order("valid desc").Order("id desc")

	if len(needFields) > 0 {
		query = query.Select(needFields)
	}

	return new(common.Result).PageQuery(query, &PhoneNumberLevel{}, &list, page, size, nil, "")
}

func (this * PhoneNumberLevel) Gets() ([]*PhoneNumberLevel,error) {
	p := make([]*PhoneNumberLevel, 0)
	err := db.GDbMgr.Get().Model(this).Where(this).Find(&p).Error
	if err == gorm.ErrRecordNotFound {
		return nil,nil
	}
	if err != nil {
		return  nil,err
	}
	return p,nil
}
