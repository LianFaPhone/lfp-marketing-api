package models

import "LianFaPhone/lfp-marketing-api/api"
import (
	"LianFaPhone/lfp-marketing-api/common"
	"LianFaPhone/lfp-marketing-api/db"
	"github.com/jinzhu/gorm"
)

type (
	PhoneNumberPool struct {
		Id      *int64  `json:"id,omitempty"        gorm:"column:id;primary_key;AUTO_INCREMENT:1;not null"` //加上type:int(11)后AUTO_INCREMENT无效
		Number  *string `json:"number,omitempty"     gorm:"column:number;type:varchar(20);unique;index"`    //订单号
		UseFlag *int    `json:"use_flag,omitempty"     gorm:"column:use_flag;type:tinyint(2);index"`
		LockExpireAt *int64 `json:"lock_expire_at,omitempty"     gorm:"column:lock_expire_at;type:bigint(20);"`
		LockUser  *string  `json:"lock_user,omitempty"     gorm:"column:lock_user;type:varchar(70);"`
		UseUser  *string  `json:"use_user,omitempty"     gorm:"column:use_user;type:varchar(70);"`
		Level   *int    `json:"level,omitempty"     gorm:"column:level;type:int(10)"`
		OrderNo  *string  `json:"order_no,omitempty"     gorm:"column:order_no;type:varchar(70);"`
		BuyerName *string `json:"buyer_name,omitempty"     gorm:"column:buyer_name;type:varchar(20);"`
		BuyerPhone *string `json:"buyer_phone,omitempty"     gorm:"column:buyer_phone;type:varchar(20);"`
		Table
	}
)

func (this *PhoneNumberPool) TableName() string {
	return "phone_number_pool"
}

func (this *PhoneNumberPool) Parse(p *api.BkPhoneNumber) *PhoneNumberPool {
	this.Id = p.Id
	this.Number = p.Number
	this.UseFlag = p.UseFlag
	this.Level = p.Level
	this.Valid = p.Valid
	this.LockExpireAt = p.LockExpireAt
	this.LockUser = p.LockUser
	this.Level = p.Level
	this.OrderNo = p.OrderNo
	this.BuyerName = p.BuyerName
	this.BuyerPhone = p.BuyerPhone
	return this
}

func (this *PhoneNumberPool) ParseAdd(number string, level int) *PhoneNumberPool {
	p := &PhoneNumberPool{
		Number:  &number,
		Level:   &level,
		UseFlag: new(int),
		LockExpireAt: new(int64),
		LockUser: new(string),
		UseUser: new(string),
		OrderNo: new(string),
	}
	p.Valid = new(int)
	*p.Valid = 1
	*p.UseFlag = 0
	return p
}

func (this *PhoneNumberPool) ParseGet(p *api.BkPhoneNumberGet) *PhoneNumberPool {
	this.Id = p.Id
	this.Number = p.Number
	return this
}

func (this *PhoneNumberPool) FtParseList(p *api.FtPhoneNumberList) *PhoneNumberPool {
	this.Valid = &p.Valid
	this.UseFlag = &p.UseFlag
	return this
}

func (this *PhoneNumberPool) BkParseList(p *api.BkPhoneNumberList) *PhoneNumberPool {
	this.Valid = p.Valid
	this.UseFlag = p.UseFlag
	this.Number = p.Number
	this.LockExpireAt = p.LockExpireAt
	this.LockUser = p.LockUser
	this.Level = p.Level
	this.OrderNo = p.OrderNo
	this.BuyerName = p.BuyerName
	this.BuyerPhone = p.BuyerPhone
	return this
}

//批量导入
func (this *PhoneNumberPool) Add() error {
	err := db.GDbMgr.Get().Create(this).Error
	if err != nil {
		return err
	}
	return nil
}

func (this *PhoneNumberPool) Update() error {
	err := db.GDbMgr.Get().Model(this).Where("id = ? ", this.Id).Updates(this).Error
	if err != nil {
		return err
	}
	return nil
}

func (this *PhoneNumberPool) UniqueByNumber(number string, condTime int64) (bool, error) {
	count := 0
	err := db.GDbMgr.Get().Model(this).Where("number = ? and created_at > ?", number, condTime).Count(&count).Error
	if err == gorm.ErrRecordNotFound {
		return true, nil
	}
	if err != nil {
		return false, err
	}
	return count == 0, nil
}

//func (this *PhoneNumberPool) UniqueByNumberValid(number string) (bool, error) {
//	count := 0
//	err := db.GDbMgr.Get().Model(this).Where("number = ?", number).Count(&count).Error
//	if err == gorm.ErrRecordNotFound {
//		return true,nil
//	}
//	if err != nil {
//		return  false,err
//	}
//	return count == 0, nil
//}

func (this *PhoneNumberPool) UseNumber(number string) (bool, error) {
	query := db.GDbMgr.Get().Model(this).Where("number = ? ", number).Update("use_flag", 1)
	if query.Error != nil {
		return false, query.Error
	}
	return query.RowsAffected > 0, nil
}

func (this *PhoneNumberPool) LockNumberById( id, nowTime, expireTime int64, user string) (bool, error) {
	query := db.GDbMgr.Get().Model(this).Where("id = ? and valid = 1 and lock_expire_at <= ? and use_flag = 0", id, nowTime).Updates(map[string]interface{}{"lock_expire_at":nowTime+expireTime, "lock_user":user})
	if query.Error != nil {
		return false, query.Error
	}
	return query.RowsAffected > 0, nil
}

func (this *PhoneNumberPool) LockNumberByNumber( number string, nowTime, expireTime int64, user string) (bool, error) {
	query := db.GDbMgr.Get().Model(this).Where("number = ? and valid = 1 and lock_expire_at <= ? and use_flag = 0", number, nowTime).Updates(map[string]interface{}{"lock_expire_at":nowTime+expireTime, "lock_user":user})
	if query.Error != nil {
		return false, query.Error
	}
	return query.RowsAffected > 0, nil
}

func (this *PhoneNumberPool) UnLockNumberById( id, nowTime int64, user string) (bool, error) {
	query := db.GDbMgr.Get().Model(this).Where("id = ? and valid = 1 and lock_expire_at <= ? and use_flag = 0 and lock_user = ?", id, nowTime, user).Updates(map[string]interface{}{"lock_expire_at":0})
	if query.Error != nil {
		return false, query.Error
	}
	return query.RowsAffected > 0, nil
}

func (this *PhoneNumberPool) UnLockNumberByNumber( number string, nowTime int64, user string) (bool, error) {
	query := db.GDbMgr.Get().Model(this).Where("number = ? and valid = 1 and lock_expire_at <= ? and use_flag = 0 and lock_user = ?", number, nowTime, user).Updates(map[string]interface{}{"lock_expire_at":0})
	if query.Error != nil {
		return false, query.Error
	}
	return query.RowsAffected > 0, nil
}

func (this *PhoneNumberPool) UseNumberById(id, nowTime int64, user, orderNo, buyerName, buyerPhone string) (bool, error) {
	query := db.GDbMgr.Get().Model(this).Where("id = ? and valid = 1 and lock_expire_at >= ? and use_flag = 0 and lock_user = ?", id, nowTime, user).Updates(map[string]interface{}{"order_no":orderNo, "use_flag": 1, "buyer_name":buyerName, "buyer_phone": buyerPhone, "use_user":user})
	if query.Error != nil {
		return false, query.Error
	}
	return query.RowsAffected > 0, nil
}

func (this *PhoneNumberPool) UseNumberByNumber(number string, nowTime int64, user, orderNo, buyerName, buyerPhone string) (bool, error) {
	query := db.GDbMgr.Get().Model(this).Where("number = ? and valid = 1 and lock_expire_at >= ? and use_flag = 0 and lock_user = ?", number, nowTime, user).Updates(map[string]interface{}{"order_no":orderNo, "use_flag": 1, "buyer_name":buyerName, "buyer_phone": buyerPhone, "use_user":user})
	if query.Error != nil {
		return false, query.Error
	}
	return query.RowsAffected > 0, nil
}

func (this *PhoneNumberPool) UnUseNumberById(id int64, user string) (bool, error) {
	query := db.GDbMgr.Get().Model(this).Where("id = ? and valid = 1 and use_flag = 1 and use_user = ?", id,  user).Updates(map[string]interface{}{"order_no":"", "use_flag": 0, "buyer_name":"", "buyer_phone": "", "lock_expire_at":0})
	if query.Error != nil {
		return false, query.Error
	}
	return query.RowsAffected > 0, nil
}

func (this *PhoneNumberPool) Get() (*PhoneNumberPool, error) {
	p := new(PhoneNumberPool)
	err := db.GDbMgr.Get().Model(this).Where(this).Last(p).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (this *PhoneNumberPool) GetByNumber(number string) (*PhoneNumberPool, error) {
	p := new(PhoneNumberPool)
	err := db.GDbMgr.Get().Model(this).Where("number = ? and valid =1", number).Last(p).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (this *PhoneNumberPool) ListWithConds(page, size int64, needFields []string, condPair []*SqlPairCondition) (*common.Result, error) {
	var list []*PhoneNumberPool
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

	return new(common.Result).PageQuery(query, &PhoneNumberPool{}, &list, page, size, nil, "")
}

func (this *PhoneNumberPool) CountWithConds(condPair []*SqlPairCondition) (int64, error) {
	var count int64
	query := db.GDbMgr.Get().Where(this)

	for i := 0; i < len(condPair); i++ {
		if condPair[i] == nil {
			continue
		}
		query = query.Where(condPair[i].Key, condPair[i].Value)
	}

	//query = query.Order("valid desc").Order("id desc")
	err := query.Count(&count).Error
	return count ,err
}
