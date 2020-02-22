package models

import (
	"LianFaPhone/lfp-marketing-api/api"
	"LianFaPhone/lfp-marketing-api/common"
	"LianFaPhone/lfp-marketing-api/db"
	"github.com/jinzhu/gorm"
	"time"
)

type CardOrder struct {
	Id         *int64  `json:"id,omitempty"        gorm:"column:id;primary_key;AUTO_INCREMENT:1;not null"`  //加上type:int(11)后AUTO_INCREMENT无效
	OrderNo    *string `json:"order_no,omitempty"     gorm:"column:order_no;type:varchar(30);unique;index"` //订单号
	ClassIsp   *int    `json:"class_isp,omitempty"     gorm:"column:class_isp;type:int(11);"`               //运营商

	ClassBigTp    *int    `json:"class_big_tp,omitempty"     gorm:"column:class_big_tp;type:int(11);"`                 //手机卡套餐类型

	ClassTp    *int    `json:"class_tp,omitempty"     gorm:"column:class_tp;type:int(11);"`                 //手机卡套餐类型
	ClassName  *string `json:"class_name,omitempty"     gorm:"-"`
	ClassDetail  *string `json:"class_detail,omitempty"     gorm:"-"`
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

	//快递信息
	Express       *string `json:"express,omitempty"     gorm:"column:express;type:varchar(20)"`             //快递名称
	ExpressNo     *string `json:"express_no,omitempty"     gorm:"column:express_no;type:varchar(30)"`       //快递单号
	ExpressRemark *string `json:"remark,omitempty"     gorm:"column:express_remark;type:varchar(50)"`               //备注
	DeliverAt     *int64  `json:"deliver_at,omitempty"     gorm:"column:deliver_at;type:bigint(20)"`        //发货时间
	//
	ICCID         *string `json:"ICCID,omitempty"     gorm:"column:ICCID;type:varchar(22)"`                 //手机卡唯一识别码
	NewPhone      *string `json:"new_phone,omitempty"     gorm:"column:new_phone;type:varchar(20)"`         //新手机号
	Guishudi      *string `json:"guishudi,omitempty"     gorm:"column:guishudi;type:varchar(30)"`           //归属于哪个门店
//	Dataurl1      *string `json:"dataurl1,omitempty"     gorm:"column:dataurl1;type:varchar(50)"`           //身份证照片地址
//	Dataurl2      *string `json:"dataurl2,omitempty"     gorm:"column:dataurl2;type:varchar(50)"`           //身份证照片地址
//	Dataurl3      *string `json:"dataurl3,omitempty"     gorm:"column:dataurl3;type:varchar(50)"`           //免冠照地址
	IP            *string `json:"ip,omitempty"     gorm:"column:ip;type:varchar(40)"`                       //ip地址
	Ips           *int    `json:"min_ips,omitempty"     gorm:"column:min_ips;type:int(10)"`                         //同一个ip下单次数
	PhoneOSTp     *int    `json:"device_os_tp,omitempty"     gorm:"column:device_os_tp;type:int(11)"`       //设备类型
	PhoneOSName   *string `json:"device_os_name,omitempty"     gorm:"-"`
	IdCardAudit   *int    `json:"idcard_audit,omitempty"     gorm:"column:idcard_audit;type:tinyint(2)"`      //身份证是否审核通过
	IdCardPicFlag   *int    `json:"idcard_pic_flag,omitempty"     gorm:"column:idcard_pic_flag;type:tinyint(2);default 0;"`      //身份证是否审核通过
	//这个要废弃吧
//	BspExpress    *string `json:"bsp_express,omitempty"     gorm:"column:bsp_express;type:varchar(20)"`       //快递名称
//	BspExpressNo  *string `json:"bsp_express_no,omitempty"     gorm:"column:bsp_express_no;type:varchar(30)"` //快递单号
	//	BspExpressRemark   *string    `json:"bsp_express_remark,omitempty"     gorm:"column:bsp_express_remark;type:varchar(50)"`
//	BspStatus     *int    `json:"bsp_status,omitempty"     gorm:"column:bsp_status;type:int(11);"` //订单状态
//	BspStatusName *string `json:"bsp_status_name,omitempty"     gorm:"-"`
	//都要废弃
//	Message       *string `json:"message,omitempty"     gorm:"column:message;type:varchar(30)"`  //湖南反馈信息
//	BspRsp        *string `json:"bsp_rsp,omitempty"     gorm:"column:bsp_rsp;type:varchar(150)"` //湖南反馈信息
//	NbMsg         *string `json:"nb_msg,omitempty"     gorm:"column:nb_msg;type:varchar(20)"`    //宁波反馈信息
	//
	IsBacklist *int `json:"is_blacklist,omitempty"     gorm:"-"`

	ThirdOrderNo    *string `json:"third_order_no,omitempty"     gorm:"column:third_order_no;type:varchar(30);"` //订单号
	ThirdResp    *string `json:"third_resp,omitempty"     gorm:"column:third_resp;type:varchar(30);"` //订单号

	Table

	//"id": "255418",
	//"orderID": "DD1909021538227895",
	//"ClassID": "5",
	//"ClassName": "联通大王卡",
	//"orderStatus": "已发货",
	//"truename": "王岩平",
	//"idcard": "'152325199807253519",
	//"phone": "17304755629",
	//"province": "内蒙古自治区",
	//"addDate": "2019/9/2 15:38:02",
	//"fahuoDate": "2019/9/2 19:53:24",
	//"city": "通辽市",
	//"area": "库伦旗",
	//"town": "库伦镇",
	//"address": "蒙医医院住院部6楼护士站",
	//"express": "中通快递",
	//"expressID": "73119112290694",
	//"beizhu": "",
	//"ICCID": "",
	//"newPhone": "15606848757",
	//"guishudi": "鄞州鄞中卓途创新电子代理2店",
	//"dataurl1": "",
	//"dataurl2": "",
	//"dataurl3": "",
	//"IP": "39.155.41.102",
	//"IPS": "0",
	//"phoneOS": "Iphone"
}

func (this *CardOrder) TableName() string {
	return "card_order"
}

func (this *CardOrder) FtParseAdd(p *api.CardOrderApply, OrderId string) *CardOrder {
	acty := &CardOrder{
		OrderNo:  &OrderId,
		TrueName: p.TrueName,
		IdCard:   p.IdCard,
		ClassTp:  p.ClassTp,
		//		CountryCode: p.CountryCode,
		Phone:        p.Phone,
		Province:     p.Province,
		ProvinceCode: p.ProvinceCode,
		City:         p.City,
		CityCode:     p.CityCode,
		Area:         p.Area,
		AreaCode:     p.AreaCode,
		Town:         p.Town,
		Address:      p.Address,
		IP:           &p.IP,
		Ips:          p.Ips,
		PhoneOSTp:    p.PhoneOSTp,
		IdCardAudit:  new(int),
		//PushFlag:    new(int),
		//PushTryCount:    new(int),
		//SmsFlag: new(int),
		ClassIsp: p.ClassISP,
		NewPhone: p.Number,
		Status:   p.Status,
		ThirdOrderNo: p.ThirdOrderNo,
	}

	if acty.Valid == nil {
		acty.Valid = new(int)
		*acty.Valid = 1
	}
	*acty.IdCardAudit = CONST_IDCARD_Audit_pending
	//*acty.PushFlag = 0
	//*acty.PushTryCount = 0
	//*acty.SmsFlag = 0

	return acty
}

func (this *CardOrder) FtParseAdd2(p *api.OldCardOrderApply, OrderId string, ClassTp, Status int, Province, City, Area, Town, Address, IP string) *CardOrder {
	acty := &CardOrder{
		OrderNo:  &OrderId,
		TrueName: p.TrueName,
		IdCard:   p.IdCard,
		ClassTp:  &ClassTp,
		Phone:    p.Phone,
		Province: &Province,
		City:     &City,
		Area:     &Area,
		Town:     &Town,
		Address:  &Address,
		IP:       &IP,

		IdCardAudit: new(int),
		//PushFlag:    new(int),
		//PushTryCount:    new(int),
		Status: &Status,
		//ThirdOrderNo: p.ThirdOrderNo,
	}

	if acty.Valid == nil {
		acty.Valid = new(int)
		*acty.Valid = 1
	}
	*acty.IdCardAudit = CONST_IDCARD_Audit_pending
	//*acty.PushFlag = 0
	//*acty.PushTryCount = 0
	//*acty.SmsFlag = 0

	return acty
}

func (this *CardOrder) BkParse(p *api.BkCardOrder) *CardOrder {
	acty := &CardOrder{
		OrderNo:            p.OrderNo,
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
		//ThirdOrderNo: p.ThirdOrderNo,
	}

	acty.Valid = p.Valid

	return acty
}

func (this *CardOrder) BkParseExtraImport(p *api.BkCardOrderExtraImprot) *CardOrder {
	acty := &CardOrder{
		OrderNo:  p.OrderNo,
		NewPhone: p.NewPhone,

		Express:   p.Express,
		ExpressNo: p.ExpressNo,
		ICCID:     p.ICCID,
		Guishudi:  p.Guishudi,
		Status:    new(int),
		DeliverAt: new(int64),
	}
	*acty.DeliverAt = time.Now().Unix()
	*acty.Status = CONST_OrderStatus_Already_Delivered
	return acty
}

func (this *CardOrder) BkParseList(p *api.BkCardOrderList) *CardOrder {
	acty := &CardOrder{
		Id:        p.Id,
		OrderNo:   p.OrderNo,
		ClassIsp: p.ClassISP,
		ClassBigTp: p.ClassBigTp,
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
		IdCardPicFlag: p.IdCardPicFlag,
	}
	acty.Valid = p.Valid
	return acty
}

func (this *CardOrder) FtParseStatus(p *api.FtCardOrderStatus) *CardOrder {
	acty := &CardOrder{
		OrderNo: &p.OrderNo,
		//Status:    p.Status,
	}
	return acty
}

func (this *CardOrder) LimitCheckByIdCardAndTime(idcard string, createdAt int64, classTp, limit int) (bool, error) {
	count := 0
	err := db.GDbMgr.Get().Model(this).Where("idcard = ? and class_tp = ? and created_at >= ? and valid = 1", idcard, classTp, createdAt).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > limit, nil
}

func (this *CardOrder) LockUniqueByIdCardAndTime(idcard string, createdAt int64, classTp int) (bool, error) {
	count := 0
	err := db.GDbMgr.Get().Model(this).Where("idcard = ? and class_tp = ? and created_at >= ? and valid = 1", idcard, classTp, createdAt).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count == 0, nil
}

func (this *CardOrder) Add() error {
	err := db.GDbMgr.Get().Create(this).Error
	if err != nil {
		return err
	}
	return nil
}

func (this *CardOrder) TxAdd(gb *gorm.DB) error {
	err := gb.Create(this).Error
	if err != nil {
		return err
	}
	return nil
}

func (this *CardOrder) CountByConds(condPair []*SqlPairCondition) (int64, error) {
	var count int64
	query := db.GDbMgr.Get().Model(this).Where(this)

	for i := 0; i < len(condPair); i++ {
		if condPair[i] == nil {
			continue
		}
		query = query.Where(condPair[i].Key, condPair[i].Value)
	}

	err := query.Count(&count).Error

	return count, err
}

func (this *CardOrder) ListWithConds(page, size int64, needFields []string, condPair []*SqlPairCondition) (*common.Result, error) {
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

func (this *CardOrder) GetsWithConds(limit int64, needFields []string, condPair []*SqlPairCondition, condStr string) ([]*CardOrder, error) {
	var list []*CardOrder
	query := db.GDbMgr.Get().Where(this)

	for i := 0; i < len(condPair); i++ {
		if condPair[i] == nil {
			continue
		}
		query = query.Where(condPair[i].Key, condPair[i].Value)
	}
	if len(condStr) > 0 {
		query = query.Where(condStr)
	}
	if len(needFields) > 0 {
		query = query.Select(needFields)
	}

	query = query.Order("id desc")

	err := query.Find(&list).Error
	if err == gorm.ErrRecordNotFound {
		return nil,nil
	}
	return list,err
}

func (this *CardOrder) ListAreaCountWithConds(page, size int64, condPair []*SqlPairCondition) (*common.Result, error) {
	var list []*CardAreasheet
	query := db.GDbMgr.Get().Where(this)

	for i := 0; i < len(condPair); i++ {
		if condPair[i] == nil {
			continue
		}
		query = query.Where(condPair[i].Key, condPair[i].Value)
	}
	//, count(id) as order_count
	query = query.Select("province, city, count(*) as order_count").Group("city")

	return new(common.Result).PageQuery2(query, &CardOrder{}, &list, page, size, nil, "")
}

func (this *CardOrder) Update() error {
	err := db.GDbMgr.Get().Model(this).Where("id = ? ", this.Id).Updates(this).Error
	if err != nil {
		return err
	}
	return nil
}


func (this *CardOrder) UpdateByOrderNo() error {
	err := db.GDbMgr.Get().Model(this).Where("order_no = ? ", this.OrderNo).Updates(this).Error
	if err != nil {
		return err

	}
	return nil
}

func (this *CardOrder) UpdateStatusByOrderNo(orderNo string, status int) error {
	err := db.GDbMgr.Get().Model(this).Where("order_no = ? ", orderNo).Update("status", status).Error
	if err != nil {
		return err

	}
	return nil
}

func (this *CardOrder) UpdatesStatus(ids []*int64, Status int) error {
	err := db.GDbMgr.Get().Model(this).Where("id IN (?)", ids).Update("status", Status).Error
	if err != nil {
		return err

	}
	return nil
}

func (this *CardOrder) UpdatesStatusByOrderNo(nos []*string, Status int) error {
	err := db.GDbMgr.Get().Model(this).Where("order_no IN (?)", nos).Update("status", Status).Error
	if err != nil {
		return err

	}
	return nil
}

func (this *CardOrder) UpdatesStatusByNewphone(phones []*string, Status int) (int64,error) {

	query := db.GDbMgr.Get().Model(this).Where("new_phone IN (?)", phones).Update("status", Status)
	return query.RowsAffected, query.Error
}

func (this *CardOrder) UpdatesPhotosByOrderNo(OrderNo, url1, url2, url3 string) error {
	err := db.GDbMgr.Get().Model(this).Where("order_no = ?", OrderNo).Updates(map[string]string{"dataurl1": url1, "dataurl2": url2, "dataurl3": url3}).Error
	if err != nil {
		return err
	}
	return nil
}

func (this *CardOrder) UpdatesByOrderNo() (int64, error) {
	query := db.GDbMgr.Get().Model(this).Where("order_no = ?", this.OrderNo).Updates(this)
	return query.RowsAffected, query.Error

}

func (this *CardOrder) Get() (*CardOrder, error) {
	acty := new(CardOrder)
	err := db.GDbMgr.Get().Where("id = ?", this.Id).Last(acty).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return acty, err
}

func (this *CardOrder) GetByIdcardAndPhone(conds *SqlPairCondition) (*CardOrder, error) {
	acty := new(CardOrder)
	query := db.GDbMgr.Get().Where("idcard = ? and phone = ?", this.IdCard, this.Phone)
	if conds != nil {
		query = query.Where(conds.Key, conds.Value)
	}
	err := query.Last(acty).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return acty, err
}

func (this *CardOrder) GetByOrderNo(no string) (*CardOrder, error) {
	acty := new(CardOrder)
	err := db.GDbMgr.Get().Where("order_no = ?", no).Last(acty).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return acty, err
}

func (this *CardOrder) GetById(id int64) (*CardOrder, error) {
	acty := new(CardOrder)
	err := db.GDbMgr.Get().Where("id = ?", id).Last(acty).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return acty, err
}

func (this *CardOrder) Del() error {
	err := db.GDbMgr.Get().Where("id = ?", this.Id).Delete(&CardOrder{}).Error

	return err
}

func (this *CardOrder) GetByIp(ip string, condPair []*SqlPairCondition) (*CardOrder, error) {
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

func (this *CardOrder) GetLimitByCond(limit int, condPair []*SqlPairCondition, needFields []string) ([]*CardOrder, error) {
	var arr []*CardOrder
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

	err := query.Order("id DESC").Find(&arr).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return arr, err
}

func (this *CardOrder) GetLimitByCond2(limit int, condPair []*SqlPairCondition) ([]*CardOrder, error) {
	var arr []*CardOrder
	query := db.GDbMgr.Get().Where(this)

	for i := 0; i < len(condPair); i++ {
		if condPair[i] == nil {
			continue
		}
		query = query.Where(condPair[i].Key, condPair[i].Value)
	}

	err := query.Order("id").Find(&arr).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return arr, err
}

