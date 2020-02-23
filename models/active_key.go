package models

import (
	"LianFaPhone/lfp-marketing-api/db"
	"github.com/jinzhu/gorm"
)

type ActiveCode struct {
	Id      *int64  `json:"id,omitempty"        gorm:"column:id;primary_key;AUTO_INCREMENT:1;not null"` //加上type:int(11)后AUTO_INCREMENT无效
	Code    *string `json:"code,omitempty"      gorm:"column:code;type:varchar(30);"`
	OrderNo *string `json:"order_no,omitempty"     gorm:"column:order_no;type:varchar(30)"`
	Phone   *string `json:"phone,omitempty"     gorm:"column:phone;type:varchar(20)"`
	Count   *int    `json:"count,omitempty"     gorm:"column:count;type:int(11);default 0"`
	Table
}

//每隔5小时清理一次，清理2天前的; count 最大5次
func (this *ActiveCode) TableName() string {
	return "active_code"
}

func (this *ActiveCode) Add() error {
	err := db.GDbMgr.Get().Create(this).Error
	if err != nil {
		return err
	}
	return nil
}

func (this *ActiveCode) GetBy(Code, OrderNo, Phone string) (*ActiveCode, error) {
	acty := new(ActiveCode)
	err := db.GDbMgr.Get().Where("code = ? and order_no = ? and phone = ?", Code, OrderNo, Phone).Last(acty).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return acty, err
}

func (this *ActiveCode) RecordCount(id int64) error {
	err := db.GDbMgr.Get().Where("id = ?", id).Update("count", gorm.Expr("count + 1")).Error
	return err
}

//得优化，分批删除
func (this *ActiveCode) Del(condPair []*SqlPairCondition) error {
	query := db.GDbMgr.Get()
	for i := 0; i < len(condPair); i++ {
		if condPair[i] == nil {
			continue
		}
		query = query.Where(condPair[i].Key, condPair[i].Value)
	}
	err := query.Delete(this).Error
	return err
}
