package models

import (
	"LianFaPhone/lfp-marketing-api/db"
)

const(
	CONST_ADTRACK_Tp_KuaiShou = 1
	CONST_ADTRACK_Tp_DouYin = 2
)

type AdTrack struct{
	Id         *int64  `json:"id,omitempty"       gorm:"column:id;primary_key;AUTO_INCREMENT:1;not null"`  //加上type:int(11)后AUTO_INCREMENT无效
	Tp         *int `json:"tp,omitempty"  gorm:"column:tp;type:int(11);"` //订单号
	OrderNo    *string `json:"order_no,omitempty"   gorm:"column:order_no;type:varchar(30);"` //订单号
	Extra      *string `json:"extra,omitempty"   gorm:"column:extra;type:varchar(3000);"`
	PushFlag   *int     `json:"push_flag,omitempty"  gorm:"column:push_flag;type:tinyint(3);"`
	SuccFlag   *int     `json:"succ_flag,omitempty"  gorm:"column:succ_flag;type:tinyint(3);"`
	Log        *string  `json:"log,omitempty"   gorm:"column:log;type:varchar(150);"`
	Table
}

func (this * AdTrack) TableName() string {
	return "ad_track"
}

func (this *AdTrack) FtParseAdd(orderNo, Extra, Log *string, tp int, pushFlag, succFlag *int) *AdTrack {
	this.Extra = Extra
	this.OrderNo = orderNo
	this.Tp = &tp
	this.Log = Log
	this.PushFlag = pushFlag
	this.SuccFlag = succFlag
	this.Valid = new(int)
	*this.Valid = 1
	return this
}



func (this * AdTrack) Add() error {
	err := db.GDbMgr.Get().Create(this).Error
	if err != nil {
		return err
	}
	return nil
}

//func (this * AdTrack) GetByOrderNo(orderNo string) (* AdTrack, error) {
//	err := db.GDbMgr.Get().Model(this).Where("order_no = ?", orderNo).Last(this).Error
//	if err == gorm.ErrRecordNotFound {
//		return nil,nil
//	}
//	return this, err
//}

func (this * AdTrack) DelWithConds(condPair []*SqlPairCondition, limit int) (int64,error) {
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
