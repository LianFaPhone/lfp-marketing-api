package models

import (
	"LianFaPhone/lfp-marketing-api/api"
	"LianFaPhone/lfp-marketing-api/common"
	"LianFaPhone/lfp-marketing-api/db"
	"github.com/jinzhu/gorm"
)

type CardOrderLog struct{
	Id         *int64  `json:"id,omitempty"       gorm:"column:id;primary_key;AUTO_INCREMENT:1;not null"`  //加上type:int(11)后AUTO_INCREMENT无效
	OrderId    *int64 `json:"order_id,omitempty"  gorm:"column:order_id;type:bigint(20);"` //订单号
	OrderNo    *string `json:"order_no,omitempty"   gorm:"column:order_no;type:varchar(30);"` //订单号
	Tp         *int   `json:"tp,omitempty"   gorm:"column:tp;type:tinyint(6);default 0"` //类型
	Log        *string `json:"log,omitempty"   gorm:"column:log;type:varchar(100);"`
	Table
}

func (this * CardOrderLog) TableName() string {
	return "card_order_log"
}

func (this *CardOrderLog) FtParseAdd(orderId *int64, orderNo *string, log *string) *CardOrderLog {
	this.OrderId = orderId
	this.OrderNo = orderNo
	this.Log = log
	if (this.Log != nil) {
		temp := []rune(*this.Log)
		if len(temp) >=100 {
			*this.Log = string(temp[0:100])
		}
	}
	this.Valid = new(int)
	*this.Valid = 1
	return this
}

func (this *CardOrderLog) FtParseAdd2(orderId *int64, orderNo *string, log string) *CardOrderLog {
	this.OrderId = orderId
	this.OrderNo = orderNo
	this.Log = &log
	temp := []rune(log)
	if len(temp) >=100 {
		*this.Log = string(temp[0:100])
	}

	this.Valid = new(int)
	*this.Valid = 1
	return this
}

func (this *CardOrderLog) BkParseList(p *api.BkCardOrderLogList) *CardOrderLog {
	this.OrderId = p.OrderId
	this.OrderNo = p.OrderNo
	return this
}

func (this * CardOrderLog) Add() error {
	err := db.GDbMgr.Get().Create(this).Error
	if err != nil {
		return err
	}
	return nil
}

func (this * CardOrderLog) DelWithConds(condPair []*SqlPairCondition, limit int) (int64,error) {
	query := db.GDbMgr.Get()

	for i := 0; i < len(condPair); i++ {
		if condPair[i] == nil {
			continue
		}
		query = query.Where(condPair[i].Key, condPair[i].Value)
	}
	query = query.Limit(limit)
	query = query.Delete(this)
	return query.RowsAffected, query.Error
}

func (this *CardOrderLog) ListWithConds(page, size int64, needFields []string, condPair []*SqlPairCondition) (*common.Result, error) {
	var list []*CardOrderLog

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

	query = query.Order("valid desc").Order("created_at desc")

	return new(common.Result).PageQuery(query, &CardOrderLog{}, &list, page, size, nil, "")
}

func (this *CardOrderLog) GetsByOrderNoWithConds(orderNo string,limit int, condPair []*SqlPairCondition) ([]*CardOrderLog, error) {
	var list []*CardOrderLog

	query := db.GDbMgr.Get().Where("order_no = ?", orderNo)

	for i := 0; i < len(condPair); i++ {
		if condPair[i] == nil {
			continue
		}
		query = query.Where(condPair[i].Key, condPair[i].Value)
	}


	query = query.Order("id desc").Limit(limit)

	err := query.Find(&list).Error
	if err == gorm.ErrRecordNotFound {
		return nil,nil
	}
	return list,err
}

func (this *CardOrderLog) GetsByOrderNo(orderNo string) ([]*CardOrderLog, error) {
	var list []*CardOrderLog

	query := db.GDbMgr.Get().Where("order_no = ?", orderNo)



	query = query.Order("id desc")

	err := query.Find(&list).Error
	if err == gorm.ErrRecordNotFound {
		return nil,nil
	}
	return list,err
}
