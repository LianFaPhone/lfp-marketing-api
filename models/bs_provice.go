package models

import (
	"LianFaPhone/lfp-marketing-api/db"
	"github.com/jinzhu/gorm"
)

type BsProvice struct {
	Id        *int    `json:"id,omitempty"        gorm:"column:id;primary_key;AUTO_INCREMENT:1;not null"`
	Code      *string `json:"code,omitempty"     gorm:"column:code;type:varchar(20)"`
	Name      *string `json:"name,omitempty"     gorm:"column:name;type:varchar(20)"`
	ShortName *string `json:"short_name,omitempty"     gorm:"column:short_name;type:varchar(20)"`
	Lng       *string `json:"lng,omitempty"     gorm:"column:lng;type:varchar(20)"`
	Lat       *string `json:"lat,omitempty"     gorm:"column:lat;type:varchar(20)"`
	Sort      *int    `json:"sort,omitempty"     gorm:"column:sort;type:int(11)"`
	Table
}

func (this *BsProvice) TableName() string {
	return "bs_province"
}

func (this *BsProvice) Gets() ([]*BsProvice, error) {
	var arr []*BsProvice
	err := db.GDbMgr.Get().Model(this).Where("valid = 1").Order("sort").Find(&arr).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return arr, err
}

func (this *BsProvice) Gets2() ([]*BsProvice, error) {
	var arr []*BsProvice
	err := db.GDbMgr.Get().Model(this).Where("valid = 1").Find(&arr).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return arr, err
}

func (this *BsProvice) GetByName(name string) (*BsProvice, error) {
	p := new(BsProvice)
	err := db.GDbMgr.Get().Model(this).Where("name = ? or short_name = ?", name, name).Last(p).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return p, err
}

func (this *BsProvice) GetByCode(code string) (*BsProvice, error) {
	p := new(BsProvice)
	err := db.GDbMgr.Get().Model(this).Where("code = ?", code).Last(p).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return p, err
}

//func (this *BsProvice) GetByNameFromCache(id int) (*BsProvice, error) {
//	data,err := db.GCache.GetProviceByName("0")
//	if err == gorm.ErrRecordNotFound {
//		return nil,nil
//	}
//	if err != nil {
//		return nil,err
//	}
//	if data == nil {
//		return nil, nil
//	}
//
//	acty, ok := data.(*CardClass)
//	if !ok {
//		return nil, errors.Annotate(err, "type err")
//	}
//	return acty,nil
//}
//
//func (this *BsProvice) InnerGets(input interface{}) (interface{}, *time.Duration, error) {
//	expire := time.Second * time.Duration(config.GConfig.Cache.CardClassByNameTimeout+rand.Intn(600))
//	//id,ok := input.(int)
//	//if !ok {
//	//	return nil,nil, errors.New("type err")
//	//}
//	acty,err := new(BsProvice).Gets2()
//	if err != nil {
//		return nil, nil, errors.Annotate(err, "Activity GetByIdAndFields")
//	}
//	if acty == nil {
//		return nil, nil, gorm.ErrRecordNotFound  // nil,nil,nil可能将是永远不超时
//	}
//	m:= make(map[acty])
//	return  acty, &expire, nil
//}
