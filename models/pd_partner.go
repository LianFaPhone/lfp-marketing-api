package models

import (
	"LianFaPhone/lfp-marketing-api/api"
	"github.com/jinzhu/gorm"
)

import (
	"LianFaPhone/lfp-marketing-api/common"
	"LianFaPhone/lfp-marketing-api/db"

)

type PdPartner struct{
	Id    *int64   `json:"id,omitempty"        gorm:"column:id;primary_key;AUTO_INCREMENT:1;not null"` //加上type:int(11)后AUTO_INCREMENT无效
	Isp *int `json:"isp,omitempty"     gorm:"column:isp;type:int(11)"` //拼音首字母缩写
	Name *string `json:"name,omitempty"     gorm:"column:name;type:varchar(20)"` //拼音首字母缩写
	Code *string `json:"code,omitempty"     gorm:"column:code;type:varchar(20)"`
	Detail *string `json:"detail,omitempty"     gorm:"column:detail;type:varchar(50)"` //拼音首字母缩写
	GsdProvince *string `json:"gsd_province,omitempty"     gorm:"column:gsd_province;type:varchar(15)"`
	GsdCity *string     `json:"gsd_city,omitempty"     gorm:"column:gsd_city;type:varchar(15)"`
	GsdProvinceCode *string   `json:"gsd_province_code,omitempty"     gorm:"column:gsd_province_code;type:varchar(15)"`
	GsdCityCode *string       `json:"gsd_city_code,omitempty"     gorm:"column:gsd_city_code;type:varchar(15)"`
	MadeIn      *string       `json:"made_in,omitempty"     gorm:"column:made_in;type:varchar(15)"`
	NoExpAddr   *string      `json:"no_exp_addr,omitempty"     gorm:"column:no_exp_addr;type:varchar(200)"`
	MinAge    *int    `json:"min_age,omitempty"     gorm:"column:min_age;type:int(11)"`
	MaxAge    *int    `json:"max_age,omitempty"     gorm:"column:max_age;type:int(11)"`
	LimitCardCount  *int         `json:"limit_card_count,omitempty"     gorm:"column:limit_card_count;type:int(11)"`
	LimitCardPeriod *int64     `json:"limit_card_period,omitempty"     gorm:"column:limit_card_period;type:bigint(20)"`
	IdcardFiveFlag  *int      `json:"idcard_five_flag,omitempty"     gorm:"column:idcard_five_flag;type:tinyint(3)"`
	IdcardFivePeriod  *int64  `json:"idcard_five_period,omitempty"     gorm:"column:idcard_five_period;type:bigint(20)"`
	RepeatExpAddrCount *int   `json:"repeat_exp_addr_count,omitempty"     gorm:"column:repeat_exp_addr_count;type:int(11)"`
	RepeatExpAddrPeriod *int  `json:"repeat_exp_addr_period,omitempty"     gorm:"column:repeat_exp_addr_period;type:bigint(20)"`
	RepeatPhoneCount  *int   `json:"repeat_phone_count,omitempty"     gorm:"column:repeat_phone_count;type:int(11)"`
	RepeatPhonePeriod  *int  `json:"repeat_phone_period,omitempty"     gorm:"column:repeat_phone_period;type:bigint(20)"`
	PrefixPath  *string  `json:"prefix_path,omitempty"     gorm:"column:prefix_path;type:varchar(30)"`

	Table
}
//partner 产品渠道
func (this *PdPartner) TableName() string {
	return "pd_partner"
}

func (this * PdPartner) ParseAdd(p *api.BkCardClassBigTpAdd) *PdPartner {
	cc := &PdPartner{
		Isp: p.ISP,
		Name: p.Name,
		Detail: p.Detail,
		Code: p.Code,
		GsdProvince: p.GsdProvince,
		GsdCity: p.GsdCity,
		GsdProvinceCode : p.GsdProvinceCode,
		GsdCityCode: p.GsdCityCode,
		MadeIn: p.MadeIn,
		NoExpAddr: p.NoExpAddr,
		MinAge : p.MinAge,
		MaxAge: p.MaxAge,
		LimitCardCount: p.LimitCardCount,
		LimitCardPeriod : p.LimitCardPeriod,
		IdcardFiveFlag: p.IdcardFiveFlag,
		IdcardFivePeriod: p.IdcardFivePeriod,
		RepeatExpAddrCount: p.RepeatExpAddrCount,
		RepeatExpAddrPeriod: p.RepeatExpAddrPeriod,
		RepeatPhoneCount: p.RepeatPhoneCount,
		RepeatPhonePeriod : p.RepeatPhonePeriod,
		PrefixPath: p.PrefixPath,
	}
	cc.Valid = new(int)
	*cc.Valid = 1
	return  cc
}

func (this * PdPartner) Parse(p *api.BkCardClassBigTp) *PdPartner {
	d :=  &PdPartner{
		Id: p.Id,

		Detail: p.Detail,

		Isp: p.ISP,
		Name: p.Name,
		Code: p.Code,
		GsdProvince: p.GsdProvince,
		GsdCity: p.GsdCity,
		GsdProvinceCode : p.GsdProvinceCode,
		GsdCityCode: p.GsdCityCode,
		MadeIn: p.MadeIn,
		NoExpAddr: p.NoExpAddr,
		MinAge : p.MinAge,
		MaxAge: p.MaxAge,
		LimitCardCount: p.LimitCardCount,
		LimitCardPeriod : p.LimitCardPeriod,
		IdcardFiveFlag: p.IdcardFiveFlag,
		IdcardFivePeriod: p.IdcardFivePeriod,
		RepeatExpAddrCount: p.RepeatExpAddrCount,
		RepeatExpAddrPeriod: p.RepeatExpAddrPeriod,
		RepeatPhoneCount: p.RepeatPhoneCount,
		RepeatPhonePeriod : p.RepeatPhonePeriod,
		PrefixPath: p.PrefixPath,
	}
	d.Valid = p.Valid
	return d
}

func (this * PdPartner) ParseList(p *api.BkCardClassBigTpList) *PdPartner {
	d:= &PdPartner{
		Isp: p.ISP,
		Name: p.Name,
	}

	d.Valid = p.Valid
	return d
}

func (this *PdPartner) Get() (*PdPartner, error) {
	err := db.GDbMgr.Get().Model(this).Where("id = ? ", this.Id).Last(this).Error
	if err == gorm.ErrRecordNotFound {
		return nil,nil
	}
	if err != nil {
		return nil,err
	}
	return this,nil
}

func (this *PdPartner) GetById(id int64) (*PdPartner, error) {
	err := db.GDbMgr.Get().Model(this).Where("id = ? ", id).Last(this).Error
	if err == gorm.ErrRecordNotFound {
		return nil,nil
	}
	if err != nil {
		return nil,err
	}
	return this,nil
}


func (this *PdPartner) Gets() ([]*PdPartner, error) {
	var arr []*PdPartner
	err := db.GDbMgr.Get().Model(this).Find(&arr).Error
	if err == gorm.ErrRecordNotFound {
		return nil,nil
	}
	if err != nil {
		return nil,err
	}
	return arr,nil
}

func (this *PdPartner) Unique() (bool, error) {
	var count int
	err := db.GDbMgr.Get().Model(this).Where("name = ? ", this.Name).Count(&count).Error
	if err != nil {
		return false,err
	}
	return count <= 0,nil
}


//批量导入
func (this *PdPartner) Add() (*PdPartner, error) {
	err := db.GDbMgr.Get().Create(this).Error
	if err != nil {
		return nil, err
	}
	err = db.GDbMgr.Get().Model(this).Where("id = ? ", this.Id).Last(this).Error
	if err != nil {
		return nil,err
	}
	return this,nil
}

func (this *PdPartner) Update() (*PdPartner, error) {
	err := db.GDbMgr.Get().Model(this).Where("id = ? ", this.Id).Updates(this).Error
	if err != nil {
		return nil,err
	}
	err = db.GDbMgr.Get().Model(this).Where("id = ? ", this.Id).Last(this).Error
	if err != nil {
		return nil,err
	}
	return this,nil
}

func (this *PdPartner) ListWithConds(page, size int64, needFields []string, condPair []*SqlPairCondition) (*common.Result, error) {
	var list []*PdPartner
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

	return new(common.Result).PageQuery(query, &PdPartner{}, &list, page, size, nil, "")
}