package models


import (
	"LianFaPhone/lfp-marketing-api/api"
	"LianFaPhone/lfp-marketing-api/config"
	"github.com/jinzhu/gorm"
	"github.com/juju/errors"
	"math/rand"
	"time"
)

import (
	"LianFaPhone/lfp-marketing-api/common"
	"LianFaPhone/lfp-marketing-api/db"

)




const (
	CONST_ISP_UnKnown      = 0
	CONST_ISP_Dianxin      = 1
	CONST_ISP_Dianxin_Name = "电信"
	CONST_ISP_YiDong       = 2
	CONST_ISP_YiDong_Name  = "移动"
	CONST_ISP_LiTong       = 3
	CONST_ISP_LiTong_Name  = "联通"
	CONST_ISP_Ali          = 4
	CONST_ISP_JD           = 5
)

type PdPartnerGoods struct{
	Id    *int64   `json:"id,omitempty"        gorm:"column:id;primary_key;AUTO_INCREMENT:1;not null"` //加上type:int(11)后AUTO_INCREMENT无效
	ISP   *int    `json:"isp,omitempty"      gorm:"column:isp;type:int(11)"`

	BigTp    *int64    `json:"big_tp_id,omitempty"       gorm:"column:big_tp_id;type:bigint(20)"` //
	Code     *string  `json:"code,omitempty"     gorm:"column:code;type:varchar(10)" `
	JdCode  *string  `json:"jd_code,omitempty"     gorm:"column:jd_code;type:varchar(15)" `
	MaxLimit    *int    `json:"max_limit,omitempty"       gorm:"column:max_limit;type:int(11)"` //加上type:int(11)后AUTO_INCREMENT无效
	Name  *string `json:"name,omitempty"     gorm:"column:name;type:varchar(15)" `
	UrlParam *string  `json:"url_param,omitempty"     gorm:"column:url_param;type:varchar(200)" `
	Detail *string `json:"detail,omitempty"     gorm:"column:detail;type:varchar(20)"` //拼音首字母缩写
	SmsFlag *int    `json:"sms_flag,omitempty"      gorm:"column:sms_flag;type:tinyint(4)"`

	IdcardDispplay *int    `json:"idcard_display,omitempty"      gorm:"column:idcard_display;type:tinyint(3);default 0"`
	ShortChain *string `json:"short_chain,omitempty"     gorm:"column:short_chain;type:varchar(50)"`
	ImgUrl *string `json:"img_url,omitempty"     gorm:"column:img_url;type:varchar(250)"`
//	FileUrl *string `json:"file_url,omitempty"     gorm:"column:file_url;type:varchar(250)"`
	LongChain *string `json:"long_chain,omitempty"     gorm:"column:long_chain;type:varchar(250)"`
	ThirdLongChain *string `json:"third_long_chain,omitempty"     gorm:"column:third_long_chain;type:varchar(250)"`
	///////////////////////////

	//////////////////////////
	Table
}

func (this *PdPartnerGoods) TableName() string {
	return "pd_partner_goods"
}

func (this * PdPartnerGoods) ParseAdd(p *api.BkCardClassAdd) *PdPartnerGoods {
	cc := &PdPartnerGoods{
		ISP: p.ISP,
		//Tp:  p.Tp,
		Name: p.Name,
		MaxLimit: p.MaxLimit,
		Detail: p.Detail,
		ImgUrl: p.ImgUrl,
//		FileUrl: p.FileUrl,
		SmsFlag: p.SmsFlag,
		BigTp: p.BigTp,
		ShortChain: p.ShortChain,
		LongChain : p.LongChain,
		ThirdLongChain :p.ThirdLongChain,
		IdcardDispplay: p.IdcardDispplay,
		Code: p.Code,
		UrlParam: p.UrlParam,
		JdCode: p.JdCode,
	}
	cc.Valid = new(int)
	*cc.Valid = 1
	return  cc
}

func (this * PdPartnerGoods) Parse(p *api.BkCardClass) *PdPartnerGoods {
	return &PdPartnerGoods{
		Id: p.Id,
		ISP: p.ISP,
		//Tp:  p.Tp,
		Name: p.Name,
		Detail: p.Detail,
		ImgUrl: p.ImgUrl,
//		FileUrl: p.FileUrl,
		SmsFlag: p.SmsFlag,
		BigTp: p.BigTp,
		ShortChain: p.ShortChain,
		LongChain : p.LongChain,
		ThirdLongChain :p.ThirdLongChain,
		MaxLimit: p.MaxLimit,
		IdcardDispplay: p.IdcardDispplay,

	}
}

func (this * PdPartnerGoods) ParseList(p *api.BkCardClassList) *PdPartnerGoods {
	return &PdPartnerGoods{
		ISP: p.ISP,
		//Tp:  p.Tp,
		Name: p.Name,

		Detail: p.Detail,
		ImgUrl: p.ImgUrl,
//		FileUrl: p.FileUrl,
		SmsFlag: p.SmsFlag,
		BigTp: p.BigTp,
		IdcardDispplay: p.IdcardDispplay,
	}
}

func (this *PdPartnerGoods) Get() (*PdPartnerGoods, error) {
	err := db.GDbMgr.Get().Model(this).Where("id = ? ", this.Id).Last(this).Error
	if err != nil {
		return nil,err
	}
	return this,nil
}


func (this *PdPartnerGoods) Gets() ([]*PdPartnerGoods, error) {
	var arr []*PdPartnerGoods
	err := db.GDbMgr.Get().Model(this).Find(&arr).Error
	if err == gorm.ErrRecordNotFound {
		return nil,nil
	}
	if err != nil {
		return nil,err
	}
	return arr,nil
}

func (this *PdPartnerGoods) GetByCode(name string) (*PdPartnerGoods, error) {
	err := db.GDbMgr.Get().Model(this).Where("code = ? ", name).Last(this).Error
	if err == gorm.ErrRecordNotFound {
		return nil,nil
	}
	if err != nil {
		return nil,err
	}
	return this,nil
}

func (this *PdPartnerGoods) GetById(tp int64) (*PdPartnerGoods, error) {
	err := db.GDbMgr.Get().Model(this).Where("id = ? ", tp).Last(this).Error
	if err == gorm.ErrRecordNotFound {
		return nil,nil
	}
	if err != nil {
		return nil,err
	}
	return this,nil
}

func (this *PdPartnerGoods) Unique() (bool, error) {
	var count int
	err := db.GDbMgr.Get().Model(this).Where("name = ? ", this.Name).Count(&count).Error
	if err != nil {
		return false,err
	}
	return count <= 0,nil
}

//批量导入
func (this *PdPartnerGoods) Add() (*PdPartnerGoods, error) {
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

func (this *PdPartnerGoods) Update() (*PdPartnerGoods, error) {
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

func (this *PdPartnerGoods) ListWithConds(page, size int64, needFields []string, condPair []*SqlPairCondition) (*common.Result, error) {
	var list []*PdPartnerGoods
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

	return new(common.Result).PageQuery(query, &PdPartnerGoods{}, &list, page, size, nil, "")
}

func (this *PdPartnerGoods) GetByCodeFromCache(name string) (*PdPartnerGoods, error) {
	data,err := db.GCache.GetPdPartnerGoodsByCode(name)
	if err == gorm.ErrRecordNotFound {
		return nil,nil
	}
	if err != nil {
		return nil,err
	}
	if data == nil {
		return nil, nil
	}

	acty, ok := data.(*PdPartnerGoods)
	if !ok {
		return nil, errors.Annotate(err, "type err")
	}
	return acty,nil
}

func (this *PdPartnerGoods) InnerGetByCode(input interface{}) (interface{}, *time.Duration, error) {
	expire := time.Second * time.Duration(config.GConfig.Cache.CardClassByNameTimeout+rand.Intn(600))
	userKey,ok := input.(string)
	if !ok {

		return nil,nil, errors.New("type err")
	}
	acty,err := new(PdPartnerGoods).GetByCode(userKey)
	if err != nil {
		return nil, nil, errors.Annotate(err, "Activity GetByIdAndFields")
	}
	if acty == nil {
		return nil, nil, gorm.ErrRecordNotFound  // nil,nil,nil可能将是永远不超时
	}
	return  acty, &expire, nil
}

//
func (this *PdPartnerGoods) GetByIdFromCache(id int64) (*PdPartnerGoods, error) {
	data,err := db.GCache.GetPdPartnerGoodsById(id)
	if err == gorm.ErrRecordNotFound {
		return nil,nil
	}
	if err != nil {
		return nil,err
	}
	if data == nil {
		return nil, nil
	}

	acty, ok := data.(*PdPartnerGoods)
	if !ok {
		return nil, errors.Annotate(err, "type err")
	}
	return acty,nil
}

func (this *PdPartnerGoods) InnerGetById(input interface{}) (interface{}, *time.Duration, error) {
	expire := time.Second * time.Duration(config.GConfig.Cache.CardClassByNameTimeout+rand.Intn(600))
	id,ok := input.(int64)
	if !ok {

		return nil,nil, errors.New("type err")
	}
	acty,err := new(PdPartnerGoods).GetById(id)
	if err != nil {
		return nil, nil, errors.Annotate(err, "Activity GetByIdAndFields")
	}
	if acty == nil {
		return nil, nil, gorm.ErrRecordNotFound  // nil,nil,nil可能将是永远不超时
	}
	return  acty, &expire, nil
}