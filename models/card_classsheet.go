package models

import (
	"LianFaPhone/lfp-marketing-api/api"
	"LianFaPhone/lfp-marketing-api/common"
	"LianFaPhone/lfp-marketing-api/db"
	"github.com/jinzhu/gorm"
)

type CardClasssheet struct {
	Id           *int64  `json:"id,omitempty"        gorm:"column:id;primary_key;AUTO_INCREMENT:1;not null"`                     //加上type:int(11)后AUTO_INCREMENT无效
	Date         *string `json:"date,omitempty"     gorm:"column:date;type:varchar(15)"` //订单号
	ClassBigTp      *int    `json:"class_big_tp,omitempty"     gorm:"column:class_big_tp;type:int(10);"`

	ClassTp      *int    `json:"class_tp,omitempty"     gorm:"column:class_tp;type:int(10)"`
	ClassISP     *int    `json:"isp,omitempty"    gorm:"column:isp;type:int(10)"`
	OrderCount   *int64  `json:"order_count,omitempty"     gorm:"column:order_count;type:bigint(20);"`
	ClassName    *string `json:"class_name,omitempty"     gorm:"-"`
	ClassIspName *string `json:"class_isp_name,omitempty"     gorm:"-"` //运营商
	DateTp       *int `json:"date_tp,omitempty"     gorm:"column:date_tp;type:tinyint(3);"`
	Table
}

func (this *CardClasssheet) TableName() string {
	return "card_order_classsheet"
}

//func (this *CardDatesheet) TxGetByDate(tx *gorm.DB, date string) (*CardDatesheet, error) {
//	m := new(CardDatesheet)
//	err := tx.Model(this).Last(m).Error
//	if err == gorm.ErrRecordNotFound {
//		return nil,nil
//	}
//
//}

func (this *CardClasssheet) ParseList(p *api.BkCardClassSheetList) *CardClasssheet {
	//this.Date  = p.Date
	//this.ClassBigTp  = p.ClassBigTp
	//this.ClassTp  = p.ClassTp
	//this.ClassISP = p.Isp
	//this.ClassName
	this.DateTp = p.DateTp

	this.Valid = p.Valid
	return this
}

func (this *CardClasssheet) ParseAdd(date string, OrderCount, LastAt int64) *CardClasssheet {
	this.Date = &date
	this.OrderCount = &OrderCount
//	this.ClassBigTp = ClassBigTp
	this.Valid = new(int)
	*this.Valid = 1
	return this
}

func (this *CardClasssheet) Parse(id *int64, OrderCount, LastAt int64) *CardClasssheet {
	this.OrderCount = &OrderCount
	this.Valid = new(int)
	*this.Valid = 0
	return this
}

func (this *CardClasssheet) Add() error {
	err := db.GDbMgr.Get().Create(this).Error
	if err != nil {
		return err
	}
	return nil
}

func (this *CardClasssheet) TxAdd(tx *gorm.DB) error {
	err := tx.Create(this).Error
	if err != nil {
		return err
	}
	return nil
}

func (this *CardClasssheet) Update() error {
	err := db.GDbMgr.Get().Model(this).Where("id = ? ", this.Id).Updates(this).Error
	if err != nil {
		return err
	}
	return nil
}

func (this *CardClasssheet) TxUpdate(tx *gorm.DB) error {
	err := tx.Model(this).Where("id = ? ", this.Id).Updates(this).Error
	if err != nil {
		return err
	}
	return nil
}

func (this *CardClasssheet) GetByDate(date string) (*CardClasssheet, error) {
	m := new(CardClasssheet)
	err := db.GDbMgr.Get().Model(this).Where("date = ?", date).Last(m).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return m, err
}

func (this *CardClasssheet) GetByDateAndTp(date string, class_tp int) (*CardClasssheet, error) {
	m := new(CardClasssheet)
	err := db.GDbMgr.Get().Model(this).Where("date = ? and class_tp = ?", date, class_tp).Last(m).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return m, err
}

func (this *CardClasssheet) TxGetByDateAndTp(tx *gorm.DB, date string, class_tp *int) (*CardClasssheet, error) {
	m := new(CardClasssheet)
	var err error
	if class_tp == nil {
		err = tx.Model(this).Where("date = ? and class_tp is null", date).Last(m).Error
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
	} else {
		err = tx.Model(this).Where("date = ? and class_tp = ?", date, class_tp).Last(m).Error
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
	}

	return m, err
}

func (this *CardClasssheet) ListWithConds(page, size int64, needFields []string, condPair []*SqlPairCondition, groupStr string) (*common.Result, error) {
	var list []*CardClasssheet
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
	if len(groupStr) > 0 {
		query = query.Group(groupStr)
	}

	return new(common.Result).PageQuery(query, &CardClasssheet{}, &list, page, size, nil, "")
}
