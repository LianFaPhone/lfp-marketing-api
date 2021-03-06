package models

import (
	"LianFaPhone/lfp-marketing-api/api"
	"LianFaPhone/lfp-marketing-api/common"
	"LianFaPhone/lfp-marketing-api/db"
	"github.com/jinzhu/gorm"
)

type CardDatesheet struct {
	Id           *int64  `json:"id,omitempty"        gorm:"column:id;primary_key;AUTO_INCREMENT:1;not null"` //加上type:int(11)后AUTO_INCREMENT无效
	DateAt       *int64  `json:"date_at,omitempty"     gorm:"column:date_at;type:bigint(20);"`
	Province     *string `json:"province,omitempty"     gorm:"column:province;type:varchar(20)"`
	ProvinceCode *string `json:"province_code,omitempty"     gorm:"column:province_code;type:varchar(20)"`
	City         *string `json:"city,omitempty"     gorm:"column:city;type:varchar(20)"`
	CityCode     *string `json:"city_code,omitempty"     gorm:"column:city_code;type:varchar(20)"`
	Isp   *int    `json:"isp,omitempty"     gorm:"column:isp;type:int(11);"`               //运营商

	PartnerId    *int64    `json:"partner_id,omitempty"     gorm:"column:partner_id;type:bigint(20);"`                 //手机卡套餐类型
	PartnerGoodsCode    *string    `json:"partner_goods_code,omitempty"     gorm:"column:partner_goods_code;type:varchar(10);"`                 //手机卡套餐类型
	PartnerGoodsName    *string `json:"partner_goods_name,omitempty"     gorm:"-"`
	IspName *string `json:"isp_name,omitempty"     gorm:"-"` //运营商

	OrderCount *int64 `json:"order_count,omitempty"     gorm:"column:order_count;type:bigint(20);"`
	Table
}

func (this *CardDatesheet) TableName() string {
	return "card_order_datesheet"
}

//
func (this *CardDatesheet) ParseList(p *api.BkCardDateSheetList) *CardDatesheet {
	ss := &CardDatesheet{
		//CityCode: p.CityCode,
		Province:p.Province,
		//ProvinceCode:p.ProvinceCode ,
		City  :p.City,

		PartnerId : p.PartnerId,
		PartnerGoodsCode  :p.PartnerGoodsCode,
		Isp  :p.ClassISP,

	}
	ss.Valid = p.Valid
	return ss
}

func (this *CardDatesheet) TxAdd(tx *gorm.DB) error {
	err := tx.Create(this).Error
	if err != nil {
		return err
	}
	return nil
}

func (this *CardDatesheet) TxUpdate(tx *gorm.DB) error {
	err := tx.Model(this).Where("id = ? ", this.Id).Updates(this).Error
	if err != nil {
		return err
	}
	return nil
}

func (this *CardDatesheet) TxGetByConds(tx *gorm.DB, date int64, Province *string, City *string, PartnerGoodsCode *string, ClassISP *int) (*CardDatesheet, error) {
	m := &CardDatesheet{
		DateAt:   &date,
		Province: Province,
		City:     City,
		PartnerGoodsCode  :PartnerGoodsCode,
		Isp  :ClassISP,
	}
	var err error
	err = tx.Model(this).Where(m).Last(m).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}

	return m, err
}

func (this *CardDatesheet) ListWithConds(page, size int64, needFields []string, condPair []*SqlPairCondition) (*common.Result, error) {
	var list []*CardDatesheet
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

	return new(common.Result).PageQuery(query, &CardDatesheet{}, &list, page, size, nil, "")
}
