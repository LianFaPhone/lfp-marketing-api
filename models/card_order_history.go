package models

import (
	"LianFaPhone/lfp-marketing-api/api"
	"LianFaPhone/lfp-marketing-api/common"
	"LianFaPhone/lfp-marketing-api/db"
	"github.com/jinzhu/gorm"
)

type CardOrderHistory struct {
	Id         *int64  `json:"id,omitempty"        gorm:"column:id;primary_key;not null"`                   //加上type:int(11)后AUTO_INCREMENT无效
	OrderNo    *string `json:"order_no,omitempty"     gorm:"column:order_no;type:varchar(30);unique;index"` //订单号
	ClassIsp   *int    `json:"class_isp,omitempty"     gorm:"column:class_isp;type:int(11);"`               //运营商
	ClassTp    *int    `json:"class_tp,omitempty"     gorm:"column:class_tp;type:int(11);"`                 //手机卡套餐类型
	ClassName  *string `json:"class_name,omitempty"     gorm:"-"`
	Status     *int    `json:"status,omitempty"     gorm:"column:status;type:int(11);"` //订单状态
	StatusName *string `json:"status_name,omitempty"     gorm:"-"`
	TrueName   *string `json:"true_name,omitempty"     gorm:"column:true_name;type:varchar(10)"` //姓名
	IdCard     *string `json:"idcard,omitempty"     gorm:"column:idcard;type:varchar(25)"`       //身份证
	//	CountryCode     *string    `json:"country_code,omitempty"     gorm:"column:country_code;type:varchar(4)"` //国家码
	Phone         *string `json:"phone,omitempty"     gorm:"column:phone;type:varchar(20)"`                 //手机号
	Province      *string `json:"province,omitempty"     gorm:"column:province;type:varchar(15)"`           //省份
	ProvinceCode  *string `json:"province_code,omitempty"     gorm:"column:province_code;type:varchar(15)"` //省份
	City          *string `json:"city,omitempty"     gorm:"column:city;type:varchar(15)"`                   //城市
	CityCode      *string `json:"city_code,omitempty"     gorm:"column:city_code;type:varchar(15)"`         //城市
	Area          *string `json:"area,omitempty"     gorm:"column:area;type:varchar(20)"`                   //区
	AreaCode      *string `json:"area_code,omitempty"     gorm:"column:area_code;type:varchar(15)"`         //区
	Town          *string `json:"town,omitempty"     gorm:"column:town;type:varchar(20)"`                   //镇街道
	Address       *string `json:"address,omitempty"     gorm:"column:address;type:varchar(50)"`             //剩余地址
	Express       *string `json:"express,omitempty"     gorm:"column:express;type:varchar(20)"`             //快递名称
	ExpressNo     *string `json:"express_no,omitempty"     gorm:"column:express_no;type:varchar(30)"`       //快递单号
	ExpressRemark *string `json:"remark,omitempty"     gorm:"column:remark;type:varchar(50)"`               //备注
	DeliverAt     *int64  `json:"deliver_at,omitempty"     gorm:"column:deliver_at;type:bigint(20)"`        //发货时间
	ICCID         *string `json:"ICCID,omitempty"     gorm:"column:ICCID;type:varchar(20)"`                 //手机卡唯一识别码
	NewPhone      *string `json:"new_phone,omitempty"     gorm:"column:new_phone;type:varchar(20)"`         //新手机号
	Guishudi      *string `json:"guishudi,omitempty"     gorm:"column:guishudi;type:varchar(30)"`           //归属于哪个门店
	Dataurl1      *string `json:"dataurl1,omitempty"     gorm:"column:dataurl1;type:varchar(50)"`           //身份证照片地址
	Dataurl2      *string `json:"dataurl2,omitempty"     gorm:"column:dataurl2;type:varchar(50)"`           //身份证照片地址
	Dataurl3      *string `json:"dataurl3,omitempty"     gorm:"column:dataurl3;type:varchar(50)"`           //免冠照地址
	IP            *string `json:"ip,omitempty"     gorm:"column:ip;type:varchar(40)"`                       //ip地址
	Ips           *int    `json:"ips,omitempty"     gorm:"column:ips;type:int(10)"`                         //同一个ip下单次数
	PhoneOSTp     *int    `json:"device_os_tp,omitempty"     gorm:"column:device_os_tp;type:int(11)"`       //设备类型
	PhoneOSName   *string `json:"device_os_name,omitempty"     gorm:"-"`
	IdCardAudit   *int    `json:"idcard_audit,omitempty"     gorm:"column:idcard_audit;type:tinyint(2)"` //身份证是否审核通过

	BspExpress   *string `json:"bsp_express,omitempty"     gorm:"column:bsp_express;type:varchar(20)"`       //快递名称
	BspExpressNo *string `json:"bsp_express_no,omitempty"     gorm:"column:bsp_express_no;type:varchar(30)"` //快递单号
	//	BspExpressRemark   *string    `json:"bsp_express_remark,omitempty"     gorm:"column:bsp_express_remark;type:varchar(50)"`
	BspStatus     *int    `json:"bsp_status,omitempty"     gorm:"column:bsp_status;type:int(11);"` //订单状态
	BspStatusName *string `json:"bsp_status_name,omitempty"     gorm:"-"`
	Message       *string `json:"message,omitempty"     gorm:"column:message;type:varchar(30)"`  //湖南反馈信息
	BspRsp        *string `json:"bsp_rsp,omitempty"     gorm:"column:bsp_rsp;type:varchar(200)"` //湖南反馈信息
	NbMsg         *string `json:"nb_msg,omitempty"     gorm:"column:nb_msg;type:varchar(20)"`    //宁波反馈信息

	IsBacklist *int `json:"is_blacklist,omitempty"     gorm:"-"`
	Table
}

func (this *CardOrderHistory) TableName() string {
	return "card_order_history"
}

func (this *CardOrderHistory) BkParse(p *api.BkCardOrder) *CardOrder {
	acty := &CardOrder{
//		Id:            p.Id,
		ClassTp:       p.ClassId,
		Status:        p.Status,
		TrueName:      p.TrueName,
		IdCard:        p.IdCard,
		Phone:         p.Phone,
		NewPhone:      p.NewPhone,
		Province:      p.Province,
		DeliverAt:     p.DeliverAt,
		City:          p.City,
		Area:          p.Area,
		Town:          p.Town,
		Address:       p.Address,
		PhoneOSTp:     p.PhoneOSTp,
		IP:            p.IP,
		Express:       p.Express,
		ExpressNo:     p.ExpressNo,
		ExpressRemark: p.ExpressRemark,
		ICCID:         p.ICCID,
	}

	acty.Valid = p.Valid

	return acty
}

//func (this *CardOrderHistory) BkParseExtraImport(p *api.BkCardOrderExtraImprot) *CardOrder {
//	acty := &CardOrder{
//		OrderNo: p.OrderNo,
//		NewPhone: p.NewPhone,
//
//		Express: p.Express,
//		ExpressNo: p.ExpressNo,
//		ICCID: p.ICCID,
//		Guishudi: p.Guishudi,
//		Status: new(int),
//		DeliverAt: new(int64),
//	}
//	*acty.DeliverAt = time.Now().Unix()
//	*acty.Status = CONST_OrderStatus_Already_Delivered
//	return acty
//}

func (this *CardOrderHistory) BkParseList(p *api.BkCardOrderList) *CardOrder {
	acty := &CardOrder{
		Id:        p.Id,
		ClassTp:   p.ClassTp,
		Status:    p.Status,
		TrueName:  p.TrueName,
		IdCard:    p.IdCard,
		Phone:     p.Phone,
		Province:  p.Province,
		City:      p.City,
		Area:      p.Area,
		Town:      p.Town,
		Address:   p.Address,
		PhoneOSTp: p.PhoneOSTp,
		IP:        p.IP,
	}
	acty.Valid = p.Valid
	return acty
}

func (this *CardOrderHistory) BkParseAdd(p *CardOrder) *CardOrderHistory {
	acty := &CardOrderHistory{
		Id:       p.Id,
		OrderNo:  p.OrderNo,
		ClassIsp: p.ClassIsp,
		ClassTp:  p.ClassTp,
		Status:   p.Status,
		TrueName: p.TrueName,
		IdCard:   p.IdCard,
		//		CountryCode  : p.CountryCode,
		Phone:         p.Phone,
		Province:      p.Province,
		ProvinceCode:  p.ProvinceCode,
		City:          p.City,
		CityCode:      p.CityCode,
		Area:          p.Area,
		AreaCode:      p.AreaCode,
		Town:          p.Town,
		Address:       p.Address,
		Express:       p.Express,
		ExpressNo:     p.ExpressNo,
		ExpressRemark: p.ExpressRemark,
		DeliverAt:     p.DeliverAt,
		ICCID:         p.ICCID,
		NewPhone:      p.NewPhone,
		Guishudi:      p.Guishudi,
		Dataurl1:      p.Dataurl1,
		Dataurl2:      p.Dataurl2,
		Dataurl3:      p.Dataurl3,
		IP:            p.IP,
		Ips:           p.Ips,
		PhoneOSTp:     p.PhoneOSTp,
		IdCardAudit:   p.IdCardAudit,
		BspExpress:    p.BspExpress,
		BspExpressNo:  p.BspExpressNo,
		//	BspExpressRemark  : p.BspExpressRemark,
		BspStatus: p.BspStatus,
		BspRsp:    p.BspRsp,
		NbMsg:     p.NbMsg,
		Message:   p.Message,
	}
	acty.Valid = p.Valid
	acty.CreatedAt = p.CreatedAt
	acty.UpdatedAt = p.UpdatedAt
	//acty.DeletedAt = p.DeletedAt
	return acty
}

func (this *CardOrderHistory) TxLockUniqueByIdCardAndTime(tx *gorm.DB, idcard string, createdAt int64, classTp int) (bool, error) {
	count := 0
	err := tx.Model(this).Set("gorm:query_option", "FOR UPDATE").Where("idcard = ? and class_tp = ? and created_at >= ? and valid = 1", idcard, classTp, createdAt).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count == 0, nil
}

func (this *CardOrderHistory) Add() error {
	err := db.GDbMgr.Get().Create(this).Error
	if err != nil {
		return err
	}
	return nil
}

func (this *CardOrderHistory) Unique() (bool, error) {
	count := 0
	err := db.GDbMgr.Get().Where("id = ?", this.Id).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count == 0, nil
}

func (this *CardOrderHistory) ListWithConds(page, size int64, needFields []string, condPair []*SqlPairCondition) (*common.Result, error) {
	var list []*CardOrder
	query := db.GDbMgr.Get().Where(this)

	for i := 0; i < len(condPair); i++ {
		if condPair[i] == nil {
			continue
		}
		query = query.Where(condPair[i].Key, condPair[i].Value)
	}
	if len(needFields) > 0 {
		query = query.Select(needFields)
	}

	query = query.Order("valid desc").Order("created_at desc")

	return new(common.Result).PageQuery(query, &CardOrder{}, &list, page, size, nil, "")
}

func (this *CardOrderHistory) Update() error {
	err := db.GDbMgr.Get().Model(this).Where("id = ? ", this.Id).Updates(this).Error
	if err != nil {
		return err
	}
	return nil
}

func (this *CardOrderHistory) UpdatesStatus(ids []*int64, Status int) error {
	err := db.GDbMgr.Get().Model(this).Where("id IN (?)", ids).Update("status", Status).Error
	if err != nil {
		return err
	}
	return nil
}

func (this *CardOrderHistory) UpdatesPhotosByOrderNo(OrderNo, url1, url2, url3 string) error {
	err := db.GDbMgr.Get().Model(this).Where("order_no = ?", OrderNo).Updates(map[string]string{"dataurl1": url1, "dataurl2": url2, "dataurl3": url3}).Error
	if err != nil {
		return err
	}
	return nil
}

func (this *CardOrderHistory) UpdatesByOrderNo() (int64, error) {
	query := db.GDbMgr.Get().Model(this).Where("order_no = ?", this.OrderNo).Updates(this)
	return query.RowsAffected, query.Error

}

func (this *CardOrderHistory) Get() (*CardOrder, error) {
	acty := new(CardOrder)
	err := db.GDbMgr.Get().Where("id = ?", this.Id).Last(acty).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return acty, err
}

func (this *CardOrderHistory) GetByOrderNo(no string) (*CardOrder, error) {
	acty := new(CardOrder)
	err := db.GDbMgr.Get().Where("order_no = ?", no).Last(acty).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return acty, err
}

func (this *CardOrderHistory) GetById(id int64) (*CardOrder, error) {
	acty := new(CardOrder)
	err := db.GDbMgr.Get().Where("id = ?", id).Last(acty).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return acty, err
}

func (this *CardOrderHistory) GetByIp(ip string, condPair []*SqlPairCondition) (*CardOrder, error) {
	acty := new(CardOrder)
	query := db.GDbMgr.Get().Where("ip = ?", ip)

	for i := 0; i < len(condPair); i++ {
		if condPair[i] == nil {
			continue
		}
		query = query.Where(condPair[i].Key, condPair[i].Value)
	}
	err := query.Last(acty).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return acty, err
}

func (this *CardOrderHistory) GetLimitByCond(limit int, condPair []*SqlPairCondition) ([]*CardOrder, error) {
	var arr []*CardOrder
	query := db.GDbMgr.Get().Where(this)

	for i := 0; i < len(condPair); i++ {
		if condPair[i] == nil {
			continue
		}
		query = query.Where(condPair[i].Key, condPair[i].Value)
	}

	err := query.Order("id DESC").Find(&arr).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return arr, err
}
