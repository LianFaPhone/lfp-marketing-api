package models

import (
	"LianFaPhone/lfp-marketing-api/api"
	"LianFaPhone/lfp-marketing-api/common"
	"LianFaPhone/lfp-marketing-api/db"
	"github.com/jinzhu/gorm"
)

type CardAreasheet struct {
	Id           *int64  `json:"id,omitempty"        gorm:"column:id;primary_key;AUTO_INCREMENT:1;not null"` //加上type:int(11)后AUTO_INCREMENT无效
	DateAt       *int64  `json:"date_at,omitempty"     gorm:"column:date_at;type:bigint(20);"`
	Province     *string `json:"province,omitempty"     gorm:"column:province;type:varchar(20)"`
	ProvinceCode *string `json:"province_code,omitempty"     gorm:"column:province_code;type:varchar(20)"`
	City         *string `json:"city,omitempty"     gorm:"column:city;type:varchar(20)"`
	CityCode     *string `json:"city_code,omitempty"     gorm:"column:city_code;type:varchar(20)"`

	OrderCount *int64 `json:"order_count,omitempty"     gorm:"column:order_count;type:bigint(20);"`

	DateTp       *int `json:"date_tp,omitempty"     gorm:"column:date_tp;type:tinyint(3);"`

	Table
}

func (this *CardAreasheet) TableName() string {
	return "card_order_areasheet"
}

//
func (this *CardAreasheet) ParseList(p *api.BkCardAreaSheetList) *CardAreasheet {
	ss := &CardAreasheet{
		CityCode: p.CityCode,
		//DateAt:
		Province: p.Province,
		ProvinceCode: p.ProvinceCode,
		City: p.City,
		//
		//OrderCount *int64 `json:"order_count,omitempty"     gorm:"column:order_count;type:bigint(20);"`
		DateTp: p.DateTp,

	}
	ss.Valid = p.Valid
	return ss
}

func (this *CardAreasheet) TxAdd(tx *gorm.DB) error {
	err := tx.Create(this).Error
	if err != nil {
		return err
	}
	return nil
}

func (this *CardAreasheet) TxUpdate(tx *gorm.DB) error {
	err := tx.Model(this).Where("id = ? ", this.Id).Updates(this).Error
	if err != nil {
		return err
	}
	return nil
}

func (this *CardAreasheet) TxGetByConds(tx *gorm.DB, date int64, Province *string, City *string, class_tp, ClassISP *int) (*CardAreasheet, error) {
	m := &CardAreasheet{
		DateAt:   &date,
		Province: Province,
		City:     City,
	}
	var err error
	err = tx.Model(this).Where(m).Last(m).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}

	return m, err
}

func (this *CardAreasheet) ListWithConds(page, size int64, needFields []string, condPair []*SqlPairCondition) (*common.Result, error) {
	var list []*CardAreasheet
	query := db.GDbMgr.Get().Model(this).Where(this)

	for i := 0; i < len(condPair); i++ {
		if condPair[i] == nil {
			continue
		}
		query = query.Where(condPair[i].Key, condPair[i].Value)
	}

	query = query.Order("valid desc").Order("created_at desc")
	if len(needFields) > 0 {
		query = query.Select(needFields)
	}

	return new(common.Result).PageQuery(query, &CardAreasheet{}, &list, page, size, nil, "")
}
