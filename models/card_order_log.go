package models

import (
	"LianFaPhone/lfp-marketing-api/api"
	"LianFaPhone/lfp-marketing-api/common"
	"LianFaPhone/lfp-marketing-api/db"
)

type CardOrderLog struct{
	Id         *int64  `json:"id,omitempty"       gorm:"column:id;primary_key;AUTO_INCREMENT:1;not null"`  //加上type:int(11)后AUTO_INCREMENT无效
	OrderId    *int64 `json:"order_id,omitempty"  gorm:"column:order_id;type:bigint(20);"` //订单号
	OrderNo    *string `json:"order_no,omitempty"   gorm:"column:order_no;type:varchar(30);"` //订单号
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
