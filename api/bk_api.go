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
		ClassTp      *int    `valid:"optional" json:"class_tp,omitempty"`
		ClassAlias   *string `valid:"optional" json:"class_alias,omitempty"`
	}

	ResBkCardOrderNewImport struct {
		SuccCount int `valid:"optional"    json:"succ_count,omitempty" `
		FailCount int `valid:"optional"    json:"fail_count,omitempty" `
	}

	BkCardOrder struct {
		OrderNo   *string `valid:"required" json:"order_no,omitempty" ` //订单号
		ClassId       *int    `valid:"optional" json:"class_tp,omitempty" `
		ClassISP      *int    `valid:"optional" json:"class_isp,omitempty"`
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
		PhoneOSTp     *int    `valid:"optional" json:"phone_os_tp,omitempty" `
		IP            *string `valid:"optional" json:"ip,omitempty" `
		Valid         *int    `valid:"optional"  json:"valid,omitempty"`
		Express       *string `valid:"optional" json:"express,omitempty"  `    //快递名称
		ExpressNo     *string `valid:"optional" json:"express_no,omitempty"  ` //快递单号
		ExpressRemark *string `valid:"optional" json:"remark,omitempty"  `     //备注

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
		ClassISP *int    `valid:"optional" json:"class_isp,omitempty"`
		ClassTp  *int    `valid:"optional" json:"class_tp,omitempty" `
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

		BlackSwitch *int `valid:"optional" json:"black_switch,omitempty"`

		Valid *int  `valid:"optional"  json:"valid,omitempty"`
		Page  int64 `valid:"optional" json:"page,omitempty"`
		Size  int64 `valid:"optional" json:"size,omitempty"`
	}

	BkCardDateSheetList struct {
		Date       string `valid:"optional"  json:"date,omitempty" ` //订单号
		OrderCount int64  `valid:"optional"  json:"order_count,omitempty" `
		Tp         int    `valid:"optional"  json:"tp,omitempty"`
		ClassTp    *int   `valid:"optional"  json:"class_tp,omitempty"`

		StartCreatedAt *int64 `valid:"optional" json:"start_created_at,omitempty"`
		EndCreatedAt   *int64 `valid:"optional" json:"end_created_at,omitempty"`
		Valid          *int   `valid:"optional"  json:"valid,omitempty"`
		Page           int64  `valid:"optional" json:"page,omitempty"`
		Size           int64  `valid:"optional" json:"size,omitempty"`
	}

	BkCardAreaSheetList struct {
		ClassTp  *int    `valid:"optional" json:"class_tp,omitempty" `
		ClassISP *int    `valid:"optional" json:"class_isp,omitempty"`
		CityCode *string `valid:"optional" json:"city_code,omitempty" `
		//StartCreatedAt  *int64    `valid:"optional" json:"start_created_at,omitempty"`
		//EndCreatedAt    *int64    `valid:"optional" json:"start_created_at,omitempty"`
		LikeStr *string `valid:"optional" json:"like_str,omitempty"`
		Valid   *int    `valid:"optional"  json:"valid,omitempty"`
		Page    int64   `valid:"optional" json:"page,omitempty"`
		Size    int64   `valid:"optional" json:"size,omitempty"`
	}

	BkPhoneNumberList struct {
		StartCreatedAt *int64  `valid:"optional" json:"start_created_at,omitempty"`
		EndCreatedAt   *int64  `valid:"optional" json:"start_created_at,omitempty"`
		Number         *string `valid:"optional"  json:"number,omitempty"`
		UseFlag        *int    `valid:"optional"  json:"use_flag,omitempty"`
		Valid          *int    `valid:"optional"  json:"valid,omitempty"`
		Page           int64   `valid:"optional" json:"page,omitempty"`
		Size           int64   `valid:"optional" json:"size,omitempty"`
	}

	BkPhoneNumber struct {
		Id      *int64  `valid:"required" json:"id,omitempty"`
		Number  *string `valid:"optional"  json:"number,omitempty"`
		UseFlag *int    `valid:"optional" json:"use_flag,omitempty"`
		Level   *int    `valid:"optional" json:"level,omitempty"`
		Valid   *int    `valid:"optional"  json:"valid,omitempty"`
	}

	BkPhoneNumberGet struct {
		Id     *int64  `valid:"optional" json:"id,omitempty"`
		Number *string `valid:"optional"  json:"number,omitempty"`
	}

	BkPhoneNumberSave struct {
		Number  *string `valid:"required"  json:"number,omitempty"`
		UseFlag int     `valid:"-" json:"-"`
		Level   *int    `valid:"required" json:"level,omitempty"`
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

	BkCardClassAdd struct{
		ISP   *int    `valid:"optional" json:"isp,omitempty"      gorm:"column:isp;type:int(11)"`
		Tp    *int    `valid:"optional" json:"tp,omitempty"       gorm:"column:tp;type:int(11)"` //加上type:int(11)后AUTO_INCREMENT无效
		Name  *string `valid:"optional" json:"name,omitempty"     gorm:"column:name;type:varchar(50)" `
		Alias *string `valid:"optional" json:"alias,omitempty"     gorm:"column:alias;type:varchar(50)"`
		ImgUrl *string `valid:"optional" json:"img_url,omitempty"     gorm:"column:img_url;type:varchar(250)"`
		FileUrl *string `valid:"optional" json:"file_url,omitempty"     gorm:"column:file_url;type:varchar(250)"`
	}

	BkCardClass struct{
		Id    *int   `valid:"required"  json:"id,omitempty"   "` //加上type:int(11)后AUTO_INCREMENT无效
		ISP   *int    `valid:"optional" json:"isp,omitempty"      gorm:"column:isp;type:int(11)"`
		Tp    *int    `valid:"optional" json:"tp,omitempty"       gorm:"column:tp;type:int(11)"` //加上type:int(11)后AUTO_INCREMENT无效
		Name  *string `valid:"optional" json:"name,omitempty"     gorm:"column:name;type:varchar(50)" `
		Alias *string `valid:"optional" json:"alias,omitempty"     gorm:"column:alias;type:varchar(50)"`
		ImgUrl *string `valid:"optional" json:"img_url,omitempty"     gorm:"column:img_url;type:varchar(250)"`
		FileUrl *string `valid:"optional" json:"file_url,omitempty"     gorm:"column:file_url;type:varchar(250)"`
	}

	BkCardClassList struct{
		ISP   *int    `valid:"optional" json:"isp,omitempty"      gorm:"column:isp;type:int(11)"`
		Tp    *int    `valid:"optional" json:"tp,omitempty"       gorm:"column:tp;type:int(11)"` //加上type:int(11)后AUTO_INCREMENT无效
		Name  *string `valid:"optional" json:"name,omitempty"     gorm:"column:name;type:varchar(50)" `
		Alias *string `valid:"optional" json:"alias,omitempty"     gorm:"column:alias;type:varchar(50)"`
		ImgUrl *string `valid:"optional" json:"img_url,omitempty"     gorm:"column:img_url;type:varchar(250)"`
		FileUrl *string `valid:"optional" json:"file_url,omitempty"     gorm:"column:file_url;type:varchar(250)"`
		Page         int64   `valid:"required" json:"page,omitempty"`
		Size         int64   `valid:"optional" json:"size,omitempty"`
	}

	BkPhotoDownload struct{
		Url  *string  `valid:"required" form:"url" json:"url,omitempty"`
	}
	//BkResPhotoDownload struct{
	//	Content  []byte  `json:"content,omitempty"`
	//}
)
