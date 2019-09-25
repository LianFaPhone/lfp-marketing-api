package models

import(
	"LianFaPhone/lfp-marketing-api/db"
	"LianFaPhone/lfp-marketing-api/common"
	"github.com/jinzhu/gorm"
)

type CardDatesheet struct {
	Id         		*int64       `json:"id,omitempty"        gorm:"column:id;primary_key;AUTO_INCREMENT:1;not null"` //加上type:int(11)后AUTO_INCREMENT无效
	Date              *string     `json:"date,omitempty"     gorm:"column:date;type:varchar(15);unique_index:carddatesheet_date_classtp"` //订单号
	ClassTp           *int        `json:"class_tp,omitempty"     gorm:"column:class_tp;type:int(10);unique_index:carddatesheet_date_classtp"`
	OrderCount        *int64      `json:"order_count,omitempty"     gorm:"column:order_count;type:bigint(20);"`
	Tp                *int        `json:"tp,omitempty"     gorm:"column:tp;type:int(10)"`
	EndFlag           *int       `json:"end_flag,omitempty"     gorm:"column:end_flag;type:tinyint(2)"`
	LastAt            *int64     `json:"last_at,omitempty"     gorm:"column:last_at;type:bigint(20);"`
	Table
}

func (this* CardDatesheet) TableName() string {
	return "sim_card_datesheet"
}

//func (this *CardDatesheet) TxGetByDate(tx *gorm.DB, date string) (*CardDatesheet, error) {
//	m := new(CardDatesheet)
//	err := tx.Model(this).Last(m).Error
//	if err == gorm.ErrRecordNotFound {
//		return nil,nil
//	}
//
//}

func (this *CardDatesheet) ParseAdd(date string, OrderCount, LastAt int64) *CardDatesheet {
	this.Date = &date
	this.OrderCount = &OrderCount
	this.EndFlag = new(int)
	*this.EndFlag = 0
	this.LastAt = new(int64)
	*this.LastAt = LastAt
	this.Valid = new(int)
	*this.Valid = 0
	return this
}

func (this *CardDatesheet) Parse(id *int64, OrderCount, LastAt int64) *CardDatesheet {
	this.OrderCount = &OrderCount
	this.EndFlag = new(int)
	*this.EndFlag = 0
	this.LastAt = new(int64)
	*this.LastAt = LastAt
	this.Valid = new(int)
	*this.Valid = 0
	return this
}

func (this *CardDatesheet) Add() error {
	err := db.GDbMgr.Get().Create(this).Error
	if err != nil {
		return  err
	}
	return nil
}

func (this *CardDatesheet) Update() error {
	err := db.GDbMgr.Get().Model(this).Where("id = ? ", this.Id).Updates(this).Error
	if err != nil {
		return  err
	}
	return nil
}

func (this *CardDatesheet) GetByDate(date string) (*CardDatesheet, error) {
	m := new(CardDatesheet)
	err := db.GDbMgr.Get().Model(this).Where("date = ?", date).Last(m).Error
	if err == gorm.ErrRecordNotFound {
		return nil,nil
	}
	return m, err
}

func (this *CardDatesheet) GetByDateAndTp(date string, class_tp int) (*CardDatesheet, error) {
	m := new(CardDatesheet)
	err := db.GDbMgr.Get().Model(this).Where("date = ? and class_tp = ?", date, class_tp).Last(m).Error
	if err == gorm.ErrRecordNotFound {
		return nil,nil
	}
	return m, err
}

func (this *CardDatesheet) ListWithConds(page, size int64, needFields []string,  condPair []*SqlPairCondition) (*common.Result, error) {
	var list []*CardDatesheet
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

	return new(common.Result).PageQuery(query, &CardDatesheet{}, &list, page, size, nil, "")
}