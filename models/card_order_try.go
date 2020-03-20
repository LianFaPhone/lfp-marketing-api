package models

import (
	"LianFaPhone/lfp-marketing-api/common"
	"LianFaPhone/lfp-marketing-api/db"
	"fmt"
	"github.com/jinzhu/gorm"
)

type CardOrderRetry struct{
	Id         *int64  `json:"id,omitempty"       gorm:"column:id;primary_key;AUTO_INCREMENT:1;not null"`  //加上type:int(11)后AUTO_INCREMENT无效
	OrderId    *int64 `json:"order_id,omitempty"  gorm:"column:order_id;type:bigint(20);"` //订单号
	OrderNo    *string `json:"order_no,omitempty"   gorm:"column:order_no;type:varchar(30);"` //订单号
	PushFlag   *int      `json:"push_flag,omitempty"   gorm:"column:push_flag;type:tinyint(4)"`
	TryCount   *int      `json:"try_count,omitempty"   gorm:"column:try_count;type:int(11); default 0"`
	LastPushAt *int64 	 `json:"last_push_at,omitempty" gorm:"column:last_push_at;type:bigint(11)"`
	Table
}

func (this * CardOrderRetry) TableName() string {
	return "card_order_retry"
}

func (this *CardOrderRetry) FtParseAdd(orderId *int64, orderNo *string, log *string) *CardOrderRetry {
	this.OrderId = orderId
	this.OrderNo = orderNo
	this.Valid = new(int)
	*this.Valid = 1
	return this
}

func (this *CardOrderRetry) FtParseAdd2(orderId *int64, orderNo *string, log string) *CardOrderRetry {
	this.OrderId = orderId
	this.OrderNo = orderNo
	this.Valid = new(int)
	*this.Valid = 1
	return this
}

func (this * CardOrderRetry) Add() error {
	err := db.GDbMgr.Get().Create(this).Error
	if err != nil {
		return err
	}
	return nil
}

func (this * CardOrderRetry) Incr(orderNo string) (int, error) {
	this.OrderNo = &orderNo
	err := db.GDbMgr.Get().Where(this).Attrs(CardOrderRetry{TryCount: new(int)}).FirstOrCreate(&this).Update("try_count", gorm.Expr("try_count+1")).Error
	if err != nil {
		return 0,err
	}
	if this.TryCount == nil {
		return 0, fmt.Errorf("nil err")
	}
	return *this.TryCount,nil
}

func (this * CardOrderRetry) DelWithConds(condPair []*SqlPairCondition, limit int) (int64,error) {
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

func (this *CardOrderRetry) ListWithConds(page, size int64, needFields []string, condPair []*SqlPairCondition) (*common.Result, error) {
	var list []*CardOrderRetry

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

	return new(common.Result).PageQuery(query, &CardOrderRetry{}, &list, page, size, nil, "")
}

func (this *CardOrderRetry) GetsByOrderNoWithConds(orderNo string,limit int, condPair []*SqlPairCondition) ([]*CardOrderRetry, error) {
	var list []*CardOrderRetry

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

func (this *CardOrderRetry) GetsByOrderNo(orderNo string) ([]*CardOrderRetry, error) {
	var list []*CardOrderRetry

	query := db.GDbMgr.Get().Where("order_no = ?", orderNo)



	query = query.Order("id desc")

	err := query.Find(&list).Error
	if err == gorm.ErrRecordNotFound {
		return nil,nil
	}
	return list,err
}

