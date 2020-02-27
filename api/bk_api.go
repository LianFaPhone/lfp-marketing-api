package api

type (
	//****后端**************************************************************************//

	//创建订单, 查询

	BkCardOrderNewImport struct {
		TrueName     *string `valid:"required" json:"true_name,omitempty"`
		IdCard       *string `valid:"required" json:"idcard,omitempty"`
		Phone        *string `valid:"required" json:"phone,omitempty"`
		Province     *string `valid:"optional" json:"province,omitempty"`      //上海市
		ProvinceCode *string `valid:"optional" json:"province_code,omitempty"` //上海市
		City         *string `valid:"optional" json:"city,omitempty"`          //上海市
		CityCode     *string `valid:"optional" json:"city_code,omitempty"`     //上海市
		Area         *string `valid:"optional" json:"area,omitempty"`          //浦东新区
		AreaCode     *string `valid:"optional" json:"area_code,omitempty"`     //浦东新区
		Town         *string `valid:"optional" json:"town,omitempty"`          //东明路街道, town street
		TownCode     *string `valid:"optional" json:"town_code,omitempty"`     //东明路街道
		Address      *string `valid:"optional" json:"address,omitempty"`       //小区及门牌号
		PartnerId    *int64    `json:"valid:"optional" partner_id,omitempty"  `                 //手机卡套餐类型
		PartnerGoodsCode    *string    `valid:"optional" json:"partner_goods_code,omitempty"   `                 //手机卡套餐类型
		Isp          *int   `json:"valid:"optional" isp,omitempty"  `
	}

	ResBkCardOrderNewImport struct {
		SuccCount int `valid:"optional"    json:"succ_count,omitempty" `
		FailCount int `valid:"optional"    json:"fail_count,omitempty" `
	}

	BkCardOrder struct {
		OrderNo   *string `valid:"required" json:"order_no,omitempty" ` //订单号
		PartnerId    *int64    `json:"valid:"optional" partner_id,omitempty"  `                 //手机卡套餐类型
		PartnerGoodsCode    *string    `valid:"optional" json:"partner_goods_code,omitempty"   `                 //手机卡套餐类型
		ClassISP      *int    `valid:"optional" json:"isp,omitempty"`
		Status        *int    `valid:"optional" json:"status,omitempty"`
		TrueName      *string `valid:"optional" json:"true_name,omitempty"`
		IdCard        *string `valid:"optional" json:"idcard,omitempty"`
		Phone         *string `valid:"optional" json:"phone,omitempty"`
		NewPhone      *string `valid:"optional" json:"new_phone,omitempty"`
		ICCID         *string `valid:"optional" json:"ICCID,omitempty"`
		Province      *string `valid:"optional" json:"province,omitempty"`
		DeliverAt     *int64  `valid:"optional" json:"deliver_at,omitempty" `
		City          *string `valid:"optional" json:"city,omitempty" `
		Area          *string `valid:"optional" json:"area,omitempty"`
		Town          *string `valid:"optional" json:"town,omitempty" `
		Address       *string `valid:"optional" json:"address,omitempty" `
		PhoneOSTp     *int    `valid:"optional" json:"device_os_tp,omitempty" `
		IP            *string `valid:"optional" json:"ip,omitempty" `
		Valid         *int    `valid:"optional"  json:"valid,omitempty"`
		Express       *string `valid:"optional" json:"express,omitempty"  `    //快递名称
		ExpressNo     *string `valid:"optional" json:"express_no,omitempty"  ` //快递单号
		ExpressRemark *string `valid:"optional" json:"express_remark,omitempty"  `     //备注

	}

	BkCardOrderStatusUpdates struct {
		OrderNo []*string `valid:"required" json:"order_no,omitempty"`

		Status *int `valid:"required" json:"status,omitempty"`
	}

	FtCardOrderPhotoUrlUpdates struct {
		ActiveKey *string `valid:"optional" json:"active_key,omitempty" `
		OrderNo   *string `valid:"required" json:"order_no,omitempty" ` //订单号
		Phone     *string `valid:"required" json:"phone,omitempty"`
		Dataurl1  *string `valid:"required" json:"dataurl1,omitempty"` //身份证照片地址
		Dataurl2  *string `valid:"required" json:"dataurl2,omitempty"` //身份证照片地址
		Dataurl3  *string `valid:"required" json:"dataurl3,omitempty"` //免冠照地址
	}

	BkCardOrderList struct {
		Id       *int64  `valid:"optional" json:"id,omitempty"`
		OrderNo   *string `valid:"optional" json:"order_no,omitempty" ` //订单号
		ClassISP *int    `valid:"optional" json:"isp,omitempty"`
		PartnerId    *int64    `json:"valid:"optional" partner_id,omitempty"  `                 //手机卡套餐类型
		PartnerGoodsCode    *string    `valid:"optional" json:"partner_goods_code,omitempty"   `                 //手机卡套餐类型

		Status   *int    `valid:"optional" json:"status,omitempty"`
		TrueName *string `valid:"optional" json:"true_name,omitempty"`
		IdCard   *string `valid:"optional" json:"idcard,omitempty"`
		Phone    *string `valid:"optional" json:"phone,omitempty"`
		Province *string `valid:"optional" json:"province,omitempty"`
		//		DeliverAt 		*int64        `valid:"optional" json:"deliver_at,omitempty" `
		City       *string `valid:"optional" json:"city,omitempty" `
		Area       *string `valid:"optional" json:"area,omitempty"`
		Town       *string `valid:"optional" json:"town,omitempty" `
		Address    *string `valid:"optional" json:"address,omitempty" `
		PhoneOSTp  *int    `valid:"optional" json:"device_os_tp,omitempty" `
		IP         *string `valid:"optional" json:"ip,omitempty" `
		MinIps     *int    `valid:"optional" json:"min_ips,omitempty" `
		UploadFlag *int    `valid:"optional" json:"upload_flag,omitempty" `

		StartCreatedAt *int64  `valid:"optional" json:"start_created_at,omitempty"`
		EndCreatedAt   *int64  `valid:"optional" json:"end_created_at,omitempty"`
		StartDeliverAt *int64  `valid:"optional" json:"start_deliver_at,omitempty"`
		EndDeliverAt   *int64  `valid:"optional" json:"end_deliver_at,omitempty"`
		LikeStr        *string `valid:"optional" json:"like_str,omitempty"`
		IdCardPicFlag   *int    `valid:"optional" json:"idcard_pic_flag,omitempty"`
		BlackSwitch *int `valid:"optional" json:"black_switch,omitempty"`

		Valid *int  `valid:"optional"  json:"valid,omitempty"`
		Page  int64 `valid:"optional" json:"page,omitempty"`
		Size  int64 `valid:"optional" json:"size,omitempty"`
	}

	BkCardClassSheetList struct {
		Date       *string `valid:"optional"  json:"date,omitempty" ` //订单号
		GroupIspFlag  *int    `valid:"optional"  json:"group_isp_flag,omitempty"`
		GroupPartnerFlag *int    `valid:"optional"  json:"group_partner_flag,omitempty"`
		GroupPartnerGoodsFlag    *int   `valid:"optional"  json:"group_partner_goods_flag,omitempty"`
		DateTp     *int    `valid:"optional"  json:"date_tp,omitempty"`
		StartCreatedAt *int64 `valid:"optional" json:"start_created_at,omitempty"`
		EndCreatedAt   *int64 `valid:"optional" json:"end_created_at,omitempty"`
		Valid          *int   `valid:"optional"  json:"valid,omitempty"`
		Page           int64  `valid:"optional" json:"page,omitempty"`
		Size           int64  `valid:"optional" json:"size,omitempty"`
	}

	BkCardAreaSheetList struct {
		GroupProvinceFlag *int `valid:"optional" json:"group_province_flag,omitempty" `
		GroupCityFlag *int `valid:"optional" json:"group_city_flag,omitempty" `

		DateTp     *int    `valid:"optional"  json:"date_tp,omitempty"`
		StartCreatedAt  *int64    `valid:"optional" json:"start_created_at,omitempty"`
		EndCreatedAt    *int64    `valid:"optional" json:"start_created_at,omitempty"`
		LikeStr *string `valid:"optional" json:"like_str,omitempty"`
		Valid   *int    `valid:"optional"  json:"valid,omitempty"`
		Page    int64   `valid:"optional" json:"page,omitempty"`
		Size    int64   `valid:"optional" json:"size,omitempty"`
	}

	BkCardDateSheetList struct {
		PartnerId  *int64    `valid:"optional" json:"partner_id,omitempty" `
		PartnerGoodsCode  *string    `valid:"optional" json:"partner_goods_code,omitempty" `
		ClassISP *int    `valid:"optional" json:"isp,omitempty"`
		Province *string `valid:"optional" json:"province,omitempty" `
		City *string `valid:"optional" json:"city,omitempty" `
		GroupPartnerFlag  *int    `valid:"optional" json:"group_partner_flag,omitempty" `
		GroupPartnerGoodsFlag  *int    `valid:"optional" json:"group_partner_goods_flag,omitempty" `
		GroupClassISPFlag *int    `valid:"optional" json:"group_isp_flag,omitempty"`
		GroupProvinceFlag *int `valid:"optional" json:"group_province_flag,omitempty" `
		GroupCityFlag *int `valid:"optional" json:"group_city_flag,omitempty" `
		StartCreatedAt  *int64    `valid:"optional" json:"start_created_at,omitempty"`
		EndCreatedAt    *int64    `valid:"optional" json:"start_created_at,omitempty"`
		LikeStr *string `valid:"optional" json:"like_str,omitempty"`
		Valid   *int    `valid:"optional"  json:"valid,omitempty"`
		Page    int64   `valid:"optional" json:"page,omitempty"`
		Size    int64   `valid:"optional" json:"size,omitempty"`
	}

	BkPhoneNumberLock struct{
		Id  *int64   `valid:"required" json:"id,omitempty"`
		//User *string `valid:"required" json:"user,omitempty"`
	}

	BkPhoneNumberUnLock struct{
		Id  *int64   `valid:"required" json:"id,omitempty"`
		//User *string `valid:"required" json:"user,omitempty"`
	}

	BkPhoneNumberUse struct{
		Id  *int64   `valid:"required" json:"id,omitempty"`
		OrderNo  *string   `valid:"required" json:"order_no,omitempty"`
		BuyerName *string  `valid:"required" json:"buyer_name,omitempty"`
		BuyerPhone *string `valid:"required" json:"buyer_phone,omitempty"`
	}

	BkPhoneNumberUnUse struct{
		Id  *int64   `valid:"required" json:"id,omitempty"`
	}


	BkPhoneNumberList struct {
		StartCreatedAt *int64  `valid:"optional" json:"start_created_at,omitempty"`
		EndCreatedAt   *int64  `valid:"optional" json:"end_created_at,omitempty"`
		Number         *string `valid:"optional"  json:"number,omitempty"`
		UseFlag        *int    `valid:"optional"  json:"use_flag,omitempty"`
		LockExpireAt *int64 `valid:"optional" json:"lock_expire_at,omitempty"`
		LockUser  *string  `valid:"optional" json:"lock_user,omitempty"`
		Level   *int    `valid:"optional" json:"level,omitempty"`
		OrderNo  *string  `valid:"optional" json:"order_no,omitempty"`
		BuyerName *string `valid:"optional" json:"buyer_name,omitempty" `
		BuyerPhone *string `valid:"optional" json:"buyer_phone,omitempty"`
		Valid          *int    `valid:"optional"  json:"valid,omitempty"`
		Page           int64   `valid:"optional" json:"page,omitempty"`
		Size           int64   `valid:"optional" json:"size,omitempty"`
	}

	BkPhoneNumber struct {
		Id      *int64  `valid:"required" json:"id,omitempty"`
		Number  *string `valid:"optional"  json:"number,omitempty"`
		UseFlag *int    `valid:"optional" json:"use_flag,omitempty"`
		Level   *int    `valid:"optional" json:"level,omitempty"`
		LockExpireAt *int64 `valid:"optional" json:"lock_expire_at,omitempty"`
		LockUser  *string  `valid:"optional" json:"lock_user,omitempty"`
		OrderNo  *string  `valid:"optional" json:"order_no,omitempty"`
		BuyerName *string `valid:"optional" json:"buyer_name,omitempty" `
		BuyerPhone *string `valid:"optional" json:"buyer_phone,omitempty"`
		Valid   *int    `valid:"optional"  json:"valid,omitempty"`
	}

	BkPhoneNumberGet struct {
		Id     *int64  `valid:"optional" json:"id,omitempty"`
		Number *string `valid:"optional"  json:"number,omitempty"`
	}

	BkPhoneNumberSave struct {
		Number  *string `valid:"required"  json:"number,omitempty"`
		UseFlag int     `valid:"-" json:"-"`
		Level   *int    `valid:"optional" json:"level,omitempty"`
		Valid   *int    `valid:"-"  json:"-"`
	}

	BkResPhoneNumberSave struct {
		SuccCount  int       `json:"succ_count,omitempty"`
		FailCount  int       `json:"fail_count,omitempty"`
		SuccNumber []*string `json:"succ_number,omitempty"`
		FailNumber []*string `json:"fail_number,omitempty"`
	}

	BkPhoneNumberLevelAdd struct {
		//		Id              *int64       `valid:"optional" json:"id,omitempty"`
		Desc  *string `valid:"optional" json:"desc,omitempty" `
		Level *int    `valid:"required" json:"level,omitempty"`
		Valid *int    `valid:"optional"  json:"valid,omitempty"`
	}

	BkPhoneNumberLevelUpdate struct {
		Id    *int64  `valid:"required" json:"id,omitempty"`
		Desc  *string `valid:"optional"  json:"desc,omitempty" `
		Level *int    `valid:"optional" json:"level,omitempty"`
		Valid *int    `valid:"optional"  json:"valid,omitempty"`
	}
	BkPhoneNumberLevelList struct {
		Id    *int64  `valid:"optional" json:"id,omitempty"`
		Desc  *string `valid:"optional"  json:"desc,omitempty"`
		Level *int    `valid:"optional" json:"level,omitempty"`
		Valid *int    `valid:"optional"  json:"valid,omitempty"`
		Page  int64   `valid:"required" json:"page,omitempty"`
		Size  int64   `valid:"optional" json:"size,omitempty"`
	}
	BkPhoneNumberLevelGets struct {
		Desc  *string `valid:"optional"  json:"desc,omitempty"`
		Level *int    `valid:"optional" json:"level,omitempty"`
		Valid *int    `valid:"optional"  json:"valid,omitempty"`
	}

	BkBlacklistPhoneAdd struct {
		Phone *string `valid:"required" json:"phone,omitempty"`
	}

	BkBlacklistPhone struct {
		Id    *int64  `valid:"required" json:"id,omitempty"`
		Phone *string `valid:"optional" json:"phone,omitempty"`
		Valid *int    `valid:"optional"  json:"valid,omitempty"`
	}

	BkBlacklistPhoneList struct {
		Id    *int64  `valid:"optional" json:"id,omitempty"`
		Phone *string `valid:"optional" json:"phone,omitempty"`
		Valid *int    `valid:"optional"  json:"valid,omitempty"`
		Page  int64   `valid:"required" json:"page,omitempty"`
		Size  int64   `valid:"optional" json:"size,omitempty"`
	}

	BkBlacklistIdcardAdd struct {
		IdCard *string `valid:"required" json:"idcard,omitempty"`
	}

	BkBlacklistIdcard struct {
		Id     *int64  `valid:"required" json:"id,omitempty"`
		IdCard *string `valid:"optional" json:"idcard,omitempty"`
		Valid  *int    `valid:"optional"  json:"valid,omitempty"`
	}

	BkBlacklistIdcardList struct {
		Id     *int64  `valid:"optional" json:"id,omitempty"`
		IdCard *string `valid:"optional" json:"idcard,omitempty"`
		Valid  *int    `valid:"optional"  json:"valid,omitempty"`
		Page   int64   `valid:"required" json:"page,omitempty"`
		Size   int64   `valid:"optional" json:"size,omitempty"`
	}

	BkBlacklistAreaList struct {
		Id           *int64  `valid:"optional" json:"id,omitempty"`
		Province     *string `valid:"optional" json:"province,omitempty" `     //省份
		ProvinceCode *string `valid:"optional" json:"province_code,omitempty"` //省份
		City         *string `valid:"optional" json:"city,omitempty"`          //城市
		CityCode     *string `valid:"optional" json:"city_code,omitempty"`     //城市
		Area         *string `valid:"optional" json:"area,omitempty"`          //区
		AreaCode     *string `valid:"optional" json:"area_code,omitempty" `    //区
		Valid        *int    `valid:"optional"  json:"valid,omitempty"`
		Page         int64   `valid:"required" json:"page,omitempty"`
		Size         int64   `valid:"optional" json:"size,omitempty"`
	}

	BkBlacklistAreaAdd struct {

		Province     *string `valid:"optional" json:"province,omitempty" `     //省份
		ProvinceCode *string `valid:"optional" json:"province_code,omitempty"` //省份
		City         *string `valid:"optional" json:"city,omitempty"`          //城市
		CityCode     *string `valid:"optional" json:"city_code,omitempty"`     //城市
		Area         *string `valid:"optional" json:"area,omitempty"`          //区
		AreaCode     *string `valid:"optional" json:"area_code,omitempty" `    //区

	}

	BkBlacklistArea struct {
		Id           *int64  `valid:"optional" json:"id,omitempty"`

		Province     *string `valid:"optional" json:"province,omitempty" `     //省份
		ProvinceCode *string `valid:"optional" json:"province_code,omitempty"` //省份
		City         *string `valid:"optional" json:"city,omitempty"`          //城市
		CityCode     *string `valid:"optional" json:"city_code,omitempty"`     //城市
		Area         *string `valid:"optional" json:"area,omitempty"`          //区

		AreaCode     *string `valid:"optional" json:"area_code,omitempty"`     //区
		Valid        *int    `valid:"optional"  json:"valid,omitempty"`
	}

	BkCardOrderExtraImprot struct {
		OrderNo   *string `valid:"required" json:"order_no,omitempty" `    //订单号
		Express   *string `valid:"optional" json:"express,omitempty"  `    //快递名称
		ExpressNo *string `valid:"optional" json:"express_no,omitempty"  ` //快递单号
		ICCID     *string `valid:"optional" json:"ICCID,omitempty"`        //手机唯一识别码
		NewPhone  *string `valid:"optional" json:"new_phone,omitempty"`
		Guishudi  *string `valid:"optional"    json:"guishudi,omitempty" `
		DeliverAt  *int64 `valid:"-"    json:"-" `
	}

	BkCardOrderActiveImprot struct {
		NewPhone  *string `valid:"optional" json:"new_phone,omitempty"`
	}

	BkResCardOrderExtraImprot struct {
		SuccCount int `valid:"optional"    json:"succ_count,omitempty" `
		FailCount int `valid:"optional"    json:"fail_count,omitempty" `
	}

	BkCardOrderIdCardCheck struct {
		OrderNo       []string   `valid:"required" json:"order_no,omitempty" ` //订单号
		TrueName []*string `valid:"optional" json:"true_name,omitempty"`
		IdCard   []*string `valid:"optional" json:"idcard,omitempty"`
	}

	BkPartnerGoodsAdd struct{
		PartnerId    *int64    `valid:"required" json:"partner_id,omitempty"       gorm:"column:partner_id;type:int(11)"` //加上type:int(11)后AUTO_INCREMENT无效
		Code     *string  `valid:"optional" json:"code,omitempty"     gorm:"column:code;type:varchar(10)" `
		JdCode  *string  `valid:"optional" json:"jd_code,omitempty"     gorm:"column:jd_code;type:varchar(15)" `
		Name  *string `valid:"optional" json:"name,omitempty"     gorm:"column:name;type:varchar(50)" `
		UrlParam *string  `valid:"optional" json:"url_param,omitempty"     gorm:"column:url_param;type:varchar(200)" `
		Detail *string `valid:"optional" json:"detail,omitempty"     gorm:"column:detail;type:varchar(50)"`
		ImgUrl *string `valid:"optional" json:"img_url,omitempty"     gorm:"column:img_url;type:varchar(250)"`
	//	FileUrl *string `valid:"optional" json:"file_url,omitempty"     gorm:"column:file_url;type:varchar(250)"`
		ShortChain *string `valid:"optional" json:"short_chain,omitempty"     gorm:"column:short_chain;type:varchar(50)"`
		LongChain *string `valid:"optional" json:"long_chain,omitempty"     gorm:"column:long_chain;type:varchar(250)"`
		ThirdLongChain *string `valid:"optional" json:"third_long_chain,omitempty"     gorm:"column:third_long_chain;type:varchar(250)"`
//		MaxLimit    *int    `valid:"optional" json:"max_limit,omitempty"       gorm:"column:max_limit;type:int(11)"` //加上type:int(11)后AUTO_INCREMENT无效
	}

	BkPartnerGoodsStatusUpdate struct {
		Id     *int64  `valid:"required"  json:"id,omitempty"   "` //加上type:int(11)后AUTO_INCREMENT无效
		Valid    *int    `valid:"optional" json:"valid,omitempty"`
	}

	BkPartnerGoods struct{
		Id    *int64   `valid:"required"  json:"id,omitempty"   "` //加上type:int(11)后AUTO_INCREMENT无效
		JdCode  *string  `valid:"optional" json:"jd_code,omitempty"     gorm:"column:jd_code;type:varchar(15)" `
		PartnerId    *int64    `valid:"optional" json:"partner_id,omitempty"       gorm:"column:partner_id;type:int(11)"` //加上type:int(11)后AUTO_INCREMENT无效

		Name  *string `valid:"optional" json:"name,omitempty"     gorm:"column:name;type:varchar(50)" `
		Detail *string `valid:"optional" json:"detail,omitempty"     gorm:"column:detail;type:varchar(50)"`
		ImgUrl *string `valid:"optional" json:"img_url,omitempty"     gorm:"column:img_url;type:varchar(250)"`
	//	FileUrl *string `valid:"optional" json:"file_url,omitempty"     gorm:"column:file_url;type:varchar(250)"`
	//	SmsFlag    *int    `valid:"optional" json:"sms_flag,omitempty"       gorm:"column:sms_flag;type:int(11)"` //加上type:int(11)后AUTO_INCREMENT无效
	//	IdcardDispplay *int    `valid:"optional" json:"idcard_display,omitempty"      gorm:"column:idcard_display;type:tinyint(3)"`

		ShortChain *string `valid:"optional" json:"short_chain,omitempty"     gorm:"column:short_chain;type:varchar(50)"`
		LongChain *string `valid:"optional" json:"long_chain,omitempty"     gorm:"column:long_chain;type:varchar(250)"`
		ThirdLongChain *string `valid:"optional" json:"third_long_chain,omitempty"     gorm:"column:third_long_chain;type:varchar(250)"`
	//	MaxLimit    *int    `valid:"optional" json:"max_limit,omitempty"       gorm:"column:max_limit;type:int(11)"` //加上type:int(11)后AUTO_INCREMENT无效

	}

	BkPartnerGoodsList struct{
		ISP   *int    `valid:"optional" json:"isp,omitempty"      gorm:"column:isp;type:int(11)"`
		PartnerId    *int64    `valid:"optional" json:"partner_id,omitempty"       gorm:"column:partner_id;type:int(11)"` //加上type:int(11)后AUTO_INCREMENT无效
		Name  *string `valid:"optional" json:"name,omitempty"     gorm:"column:name;type:varchar(50)" `
		Detail *string `valid:"optional" json:"detail,omitempty"     gorm:"column:detail;type:varchar(50)"`
		ImgUrl *string `valid:"optional" json:"img_url,omitempty"     gorm:"column:img_url;type:varchar(250)"`
	//	FileUrl *string `valid:"optional" json:"file_url,omitempty"     gorm:"column:file_url;type:varchar(250)"`
	//	SmsFlag    *int    `valid:"optional" json:"sms_flag,omitempty"       gorm:"column:sms_flag;type:int(11)"` //加上type:int(11)后AUTO_INCREMENT无效
	//	IdcardDispplay *int    `valid:"optional" json:"idcard_display,omitempty"      gorm:"column:idcard_display;type:tinyint(3)"`
		Valid    *int    `valid:"optional" json:"valid,omitempty"      `

		Page         int64   `valid:"required" json:"page,omitempty"`
		Size         int64   `valid:"optional" json:"size,omitempty"`
	}
	//

	BkPartnerAdd struct{
		ISP   *int    `valid:"optional" json:"isp,omitempty"   `
		Detail *string `valid:"optional" json:"detail,omitempty"  `
		Name *string `valid:"optional" json:"name,omitempty"   `
		Code *string `valid:"optional" json:"code,omitempty"  `
		GsdProvince *string `valid:"optional" json:"gsd_province,omitempty" `
		GsdCity *string     `valid:"optional" json:"gsd_city,omitempty"  `
		GsdProvinceCode *string   `valid:"optional" json:"gsd_province_code,omitempty" `
		GsdCityCode *string       `valid:"optional" json:"gsd_city_code,omitempty"  `
		MadeIn      *string       `valid:"optional" json:"made_in,omitempty" `
		NoExpAddr   *string      `valid:"optional" json:"no_exp_addr,omitempty"  `
		MinAge    *int    `valid:"optional" json:"min_age,omitempty"     `
		MaxAge    *int    `valid:"optional" json:"max_age,omitempty" `
		LimitCardCount  *int         `valid:"optional" json:"limit_card_count,omitempty" `
		LimitCardPeriod *int64     `valid:"optional" json:"limit_card_period,omitempty" `
		IdcardFiveFlag  *int      `valid:"optional" json:"idcard_five_flag,omitempty"   `
		IdcardFivePeriod  *int64  `valid:"optional" json:"idcard_five_period,omitempty" `
		RepeatExpAddrCount *int   `valid:"optional" json:"repeat_exp_addr_count,omitempty" `
		RepeatExpAddrPeriod *int  `valid:"optional" json:"repeat_exp_addr_period,omitempty"  `
		RepeatPhoneCount  *int   `valid:"optional" json:"repeat_phone_count,omitempty"  `
		RepeatPhonePeriod  *int  `valid:"optional" json:"repeat_phone_period,omitempty" `
		PrefixPath  *string  `valid:"optional" json:"prefix_path,omitempty"  `
		IdcardDispplay *int    `valid:"optional" json:"idcard_display,omitempty"      gorm:"column:idcard_display;type:tinyint(3)"`
		SmsFlag    *int    `valid:"optional" json:"sms_flag,omitempty"       gorm:"column:sms_flag;type:int(11)"` //加上type:int(11)后AUTO_INCREMENT无效
		Stock *int    `valid:"optional" json:"stock,omitempty"      gorm:"column:stock;type:int(11);default 0"`
		ProductionNotes *string `valid:"optional" json:"production_notes,omitempty"      gorm:"column:production_notes;type:varchar(50);"`

	}

	BkPartner struct{
		Id    *int64   `valid:"required"  json:"id,omitempty"   "`
		ISP   *int    `valid:"optional" json:"isp,omitempty"    `
		Name *string `valid:"optional" json:"name,omitempty"  `
		Detail *string `valid:"optional" json:"detail,omitempty"  `
		Code *string `valid:"optional" json:"code,omitempty"  `
		GsdProvince *string `valid:"optional" json:"gsd_province,omitempty" `
		GsdCity *string     `valid:"optional" json:"gsd_city,omitempty"  `
		GsdProvinceCode *string   `valid:"optional" json:"gsd_province_code,omitempty" `
		GsdCityCode *string       `valid:"optional" json:"gsd_city_code,omitempty"  `
		MadeIn      *string       `valid:"optional" json:"made_in,omitempty" `
		NoExpAddr   *string      `valid:"optional" json:"no_exp_addr,omitempty"  `
		MinAge    *int    `valid:"optional" json:"min_age,omitempty"     `
		MaxAge    *int    `valid:"optional" json:"max_age,omitempty" `
		LimitCardCount  *int         `valid:"optional" json:"limit_card_count,omitempty" `
		LimitCardPeriod *int64     `valid:"optional" json:"limit_card_period,omitempty" `
		IdcardFiveFlag  *int      `valid:"optional" json:"idcard_five_flag,omitempty"   `
		IdcardFivePeriod  *int64  `valid:"optional" json:"idcard_five_period,omitempty" `
		RepeatExpAddrCount *int   `valid:"optional" json:"repeat_exp_addr_count,omitempty" `
		RepeatExpAddrPeriod *int  `valid:"optional" json:"repeat_exp_addr_period,omitempty"  `
		RepeatPhoneCount  *int   `valid:"optional" json:"repeat_phone_count,omitempty"  `
		RepeatPhonePeriod  *int  `valid:"optional" json:"repeat_phone_period,omitempty" `
		PrefixPath  *string  `valid:"optional" json:"prefix_path,omitempty"  `
		IdcardDispplay *int    `valid:"optional" json:"idcard_display,omitempty"      gorm:"column:idcard_display;type:tinyint(3)"`
		SmsFlag    *int    `valid:"optional" json:"sms_flag,omitempty"       gorm:"column:sms_flag;type:int(11)"` //加上type:int(11)后AUTO_INCREMENT无效
		Stock *int    `valid:"optional" json:"stock,omitempty"      gorm:"column:stock;type:int(11);default 0"`
		ProductionNotes *string `valid:"optional" json:"production_notes,omitempty"      gorm:"column:production_notes;type:varchar(50);"`

		Valid    *int    `valid:"optional" json:"valid,omitempty"      `
	}

	BkPartnerList struct{
		ISP   *int    `valid:"optional" json:"isp,omitempty"      gorm:"column:isp;type:int(11)"`

		Name *string `valid:"optional" json:"name,omitempty"     gorm:"column:detail;type:varchar(50)"`
		Valid    *int    `valid:"optional" json:"valid,omitempty"       gorm:"column:valid;type:int(11)"` //加上type:int(11)后AUTO_INCREMENT无效
		Page         int64   `valid:"required" json:"page,omitempty"`
		Size         int64   `valid:"optional" json:"size,omitempty"`
	}

	BkPdPartnerGoodsStatusUpdate struct{
		Id    *int64   `valid:"required"  json:"id,omitempty"   "`
		Valid    *int    `valid:"optional" json:"valid,omitempty"       gorm:"column:valid;type:int(11)"` //加上type:int(11)后AUTO_INCREMENT无效
	}

	BkPhotoDownload struct{
		Url  *string  `valid:"required" form:"url" json:"url,omitempty"`
	}
	//BkResPhotoDownload struct{
	//	Content  []byte  `json:"content,omitempty"`
	//}

	BkQaAdd struct{
		Title *string   `valid:"required" json:"title,omitempty"`
		Content *string `valid:"optional" json:"content,omitempty"`
	}

	BkQa struct{
		Id *int64   `valid:"required" json:"id,omitempty"`
		Title *string   `valid:"optional" json:"title,omitempty"`
		Content *string `valid:"optional" json:"content,omitempty"`
	}

	BkQaList struct{
		LikeStr      *string  `valid:"required" json:"like_str,omitempty"`
		Page         int64   `valid:"required" json:"page,omitempty"`
		Size         int64   `valid:"optional" json:"size,omitempty"`
	}

	BkCardOrderLogList struct{
		OrderId    *int64   `valid:"optional" json:"order_id,omitempty"  gorm:"column:order_id;type:bigint(20);"` //订单号
		OrderNo    *string  `valid:"optional" json:"order_no,omitempty"   gorm:"column:order_no;type:varchar(30);index"` //订单号
		Page         int64   `valid:"required" json:"page,omitempty"`
		Size         int64   `valid:"optional" json:"size,omitempty"`
	}
	BkCardOrderIdcardPicGet struct{
		OrderNo    *string  `valid:"required" json:"order_no,omitempty"   gorm:"column:order_no;type:varchar(30);index"` //订单号
	}
)
