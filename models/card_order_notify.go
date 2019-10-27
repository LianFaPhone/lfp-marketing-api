package models

import (
	"LianFaPhone/lfp-marketing-api/db"
	"github.com/jinzhu/gorm"
	"time"
)

const (
	CONST_OrderNotifyTp_Express = 1
)

type CardOrderNotify struct {
	Id *int64 `json:"id,omitempty"        gorm:"column:id;primary_key;AUTO_INCREMENT:1;not null"` //加上type:int(11)后AUTO_INCREMENT无效

	OrderNo  *string `json:"order_no,omitempty"     gorm:"column:order_no;type:varchar(30);index"` //订单号
	OrderId  *int64  `json:"order_id,omitempty"      gorm:"column:order_id;type:bigint(20);"`
	Tp       *int    `json:"tp,omitempty"           gorm:"column:tp;type:int(10)"` // 1快递短信，
	PushFlag *int    `json:"push_flag,omitempty"    gorm:"column:push_flag;type:tinyint(2)"`
	TryCount *int    `json:"try_count,omitempty"     gorm:"column:try_count;type:int(10);"`
	LastAt   *int64  `json:"last_at,omitempty"     gorm:"column:last_at;type:bigint(20);"`
	Table
}

func (this *CardOrderNotify) TableName() string {
	return "card_order_notify"
}

func (this *CardOrderNotify) Add(OrderNo *string, OrderId *int64, Tp *int) error {
	p := &CardOrderNotify{
		OrderNo:  OrderNo,
		OrderId:  OrderId,
		Tp:       Tp,
		TryCount: new(int),
	}
	p.Valid = new(int)
	*p.Valid = 1
	err := db.GDbMgr.Get().Create(this).Error
	if err != nil {
		return err
	}
	return nil
}

func (this *CardOrderNotify) GetsByLimit(conds []*SqlPairCondition, limit int) ([]*CardOrderNotify, error) {
	var acty []*CardOrderNotify
	query := db.GDbMgr.Get().Where(this)
	for i := 0; i < len(conds); i++ {
		if conds[i] == nil {
			continue
		}
		query = query.Where(conds[i].Key, conds[i].Value)
	}
	err := query.Limit(limit).Order("id desc").Find(&acty).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return acty, err
}

func (this *CardOrderNotify) IncrTryCount(id int64) error {
	err := db.GDbMgr.Get().Model(this).Where("id = ?", id).Updates(map[string]interface{}{"try_count": gorm.Expr("try_count + 1"), "last_at": time.Now().Unix()}).Error
	if err == gorm.ErrRecordNotFound {
		return nil
	}
	return err
}

func (this *CardOrderNotify) SetPushFlag(id int64) error {
	err := db.GDbMgr.Get().Model(this).Where("id = ?", id).Updates(map[string]interface{}{"try_count": gorm.Expr("try_count + 1"), "last_at": time.Now().Unix(), "push_flag": 1}).Error
	if err == gorm.ErrRecordNotFound {
		return nil
	}
	return err
}

func (this *CardOrderNotify) Del(conds []*SqlPairCondition) error {
	query := db.GDbMgr.Get()
	for i := 0; i < len(conds); i++ {
		if conds[i] == nil {
			continue
		}
		query = query.Where(conds[i].Key, conds[i].Value)
	}
	return query.Delete(&CardOrderNotify{}, this).Error
}
