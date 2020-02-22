package models

type CardIdcardPic struct {
	Id       *int64  `json:"id,omitempty"        gorm:"column:id;primary_key;AUTO_INCREMENT:1;not null"`  //加上type:int(11)后AUTO_INCREMENT无效
	OrderNo  *string `json:"order_no,omitempty"     gorm:"column:order_no;type:varchar(30)"` //订单号
	PicUrl1  *string `json:"pic_url1,omitempty"     gorm:"column:pic_url1;type:varchar(300)"` //订单号
	PicUrl2  *string `json:"pic_url2,omitempty"     gorm:"column:pic_url2;type:varchar(300)"` //订单号
	PicUrl3  *string `json:"pic_url3,omitempty"     gorm:"column:pic_url3;type:varchar(300)"` //订单号

	Table
}
