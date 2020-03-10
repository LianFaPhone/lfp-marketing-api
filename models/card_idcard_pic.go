package models

import (
	"LianFaPhone/lfp-marketing-api/common"
	"LianFaPhone/lfp-marketing-api/db"
	"github.com/jinzhu/gorm"
)

type CardIdcardPic struct {
	Id       *int64  `json:"id,omitempty"        gorm:"column:id;primary_key;AUTO_INCREMENT:1;not null"`  //加上type:int(11)后AUTO_INCREMENT无效
	OrderNo  *string `json:"order_no,omitempty"     gorm:"column:order_no;type:varchar(30)"` //订单号
	PicUrl1  *string `json:"pic_url1,omitempty"     gorm:"column:pic_url1;type:varchar(300)"` //订单号
	PicUrl2  *string `json:"pic_url2,omitempty"     gorm:"column:pic_url2;type:varchar(300)"` //订单号
	PicUrl3  *string `json:"pic_url3,omitempty"     gorm:"column:pic_url3;type:varchar(300)"` //订单号

	Table
}

func (this *CardIdcardPic)TableName() string {
	return "card_idcard_pic"
}

func (this *CardIdcardPic) Add(orderNo, pic1, pic2, pic3 *string) error {
	this.OrderNo = orderNo
	this.PicUrl1 = pic1
	this.PicUrl2 = pic2
	this.PicUrl3 = pic3
	this.Valid = new(int)
	*this.Valid = 1

	err := db.GDbMgr.Get().Create(this).Error
	if err != nil {
		return  err
	}
	return nil
}

func (this *CardIdcardPic) UpdateByOrderNo( orderNo, pic1, pic2, pic3 *string) error {
	this.OrderNo = orderNo
	this.PicUrl1 = pic1
	this.PicUrl2 = pic2
	this.PicUrl3 = pic3

	err := db.GDbMgr.Get().Model(this).Where("order_no = ? ", orderNo).Updates(this).Error
	if err != nil {
		return err
	}
	return nil
}

func (this *CardIdcardPic) GetByOrderNo(order_no *string) (*CardIdcardPic, error) {
	err := db.GDbMgr.Get().Model(this).Where("order_no = ? ", order_no).Last(this).Error
	if err == gorm.ErrRecordNotFound{
		return nil,nil
	}
	if err != nil {
		return nil,err
	}
	return this,nil
}

func (this *CardIdcardPic) DelByOrderNo(order_no *string) error {
	err := db.GDbMgr.Get().Model(this).Where("order_no = ? ", order_no).Delete(this).Error
	if err != nil {
		return err
	}
	return nil
}

func (this *CardIdcardPic) ListWithConds(page, size int64, needFields []string, condPair []*SqlPairCondition) (*common.Result, error) {
	var list []*CardIdcardPic
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

	return new(common.Result).PageQuery(query, &CardIdcardPic{}, &list, page, size, nil, "")
}