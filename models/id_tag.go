package models

import (
	"LianFaPhone/lfp-marketing-api/db"
	"github.com/jinzhu/gorm"
)

type IdRecorder struct {
	Id    int64   `json:"id,omitempty"        gorm:"column:id;primary_key;AUTO_INCREMENT:1;not null"` //加上type:int(11)后AUTO_INCREMENT无效
	IdTag *int64  `json:"id_tag,omitempty"     gorm:"column:id_tag;type:bigint(20)"`                  //订单号
	Name  *string `json:"name,omitempty"     gorm:"column:name;type:varchar(30);unique"`        //订单号
	Table
}

func (this *IdRecorder) TableName() string {
	return "id_recorder"
}

func (this *IdRecorder) Update(IdTag int64) error {
	err := db.GDbMgr.Get().Model(this).Where("id = ? ", this.Id).Update("id_tag", IdTag).Error
	if err != nil {
		return err
	}
	return nil
}

func (this *IdRecorder) TxUpdate(tx *gorm.DB, id, IdTag int64) error {
	err := tx.Model(this).Where("id = ? ", id).Update("id_tag", IdTag).Error
	if err != nil {
		return err
	}
	return nil
}

func (this *IdRecorder) Add(Name string, IdTag int64) (*IdRecorder, error) {
	p := &IdRecorder{
		IdTag: &IdTag,
		Name:  &Name,
	}
	p.Valid = new(int)
	*p.Valid = 1
	err := db.GDbMgr.Get().Model(this).Create(p).Error
	if err != nil {
		return p, err
	}
	return p, nil
}

func (this *IdRecorder) GetByName(name string) (*IdRecorder, error) {
	acty := new(IdRecorder)
	err := db.GDbMgr.Get().Where("name = ?", name).Last(acty).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return acty, err
}
