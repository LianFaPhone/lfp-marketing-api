package models

type BlacklistArea struct{
	Id         		int64      `json:"id,omitempty"        gorm:"column:id;primary_key;AUTO_INCREMENT:1;not null"` //加上type:int(11)后AUTO_INCREMENT无效
	IdTag    	    int64     `json:"id_tag,omitempty"     gorm:"column:id_tag;type:bigint(20)"` //订单号
	Name    	    string     `json:"name,omitempty"     gorm:"column:name;type:varchar(30);index;unique"` //订单号
	Table
}

func (this * BlacklistArea) TableName() string{
	return "area"
}
