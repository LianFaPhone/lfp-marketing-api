package api

type(
	//****后端**************************************************************************//

	//创建订单, 查询

	BkCardOrder struct{
		Id         		*int64      `valid:"required" json:"id,omitempty"`
		ClassId    		*int     `valid:"optional" json:"class_tp,omitempty" `
		ClassISP    *int      `valid:"optional" json:"class_isp,omitempty"`
		Status  	*int    `valid:"optional" json:"status,omitempty"`
		TrueName  		*string    `valid:"optional" json:"true_name,omitempty"`
		IdCard  		*string      `valid:"optional" json:"idcard,omitempty"`
		Phone  			*string      `valid:"optional" json:"phone,omitempty"`
		NewPhone  		*string  `valid:"optional" json:"new_phone,omitempty"`
		ICCID          	*string    `valid:"optional" json:"ICCID,omitempty"`
		Province 		*string   `valid:"optional" json:"province,omitempty"`
		DeliverAt 		*int64   `valid:"optional" json:"deliver_at,omitempty" `
		City    		*string    `valid:"optional" json:"city,omitempty" `
		Area    		*string    `valid:"optional" json:"area,omitempty"`
		Town    		*string    `valid:"optional" json:"town,omitempty" `
		Address    		*string    `valid:"optional" json:"address,omitempty" `
		PhoneOSTp    		*int   `valid:"optional" json:"phone_os_tp,omitempty" `
		IP          	*string    `valid:"optional" json:"ip,omitempty" `
		Valid     		*int       `valid:"optional"  json:"valid,omitempty"`
		Express    		*string    `valid:"optional" json:"express,omitempty"  ` //快递名称
		ExpressNo    	*string     `valid:"optional" json:"express_no,omitempty"  ` //快递单号
		ExpressRemark   *string    `valid:"optional" json:"remark,omitempty"  `  //备注

	}

	BkCardOrderStatusUpdates struct{
		Id         []*int64      `valid:"required" json:"id,omitempty"`

		Status  	*int    `valid:"required" json:"status,omitempty"`
	}

	FtCardOrderPhotoUrlUpdates struct{
		ActiveKey       *string      `valid:"required" json:"active_key,omitempty" `
		OrderNo    	    *string     `valid:"required" json:"order_no,omitempty" ` //订单号
		Phone  			*string    `valid:"required" json:"phone,omitempty"`
		Dataurl1       	*string    `valid:"required" json:"dataurl1,omitempty"`   //身份证照片地址
		Dataurl2       	*string    `valid:"required" json:"dataurl2,omitempty"`   //身份证照片地址
		Dataurl3       	*string    `valid:"required" json:"dataurl3,omitempty"`   //免冠照地址
	}

	BkCardOrderList struct{
		Id              *int64        `valid:"optional" json:"id,omitempty"`
		ClassISP    	*int      `valid:"optional" json:"class_isp,omitempty"`
		ClassTp    		*int          `valid:"optional" json:"class_tp,omitempty" `
		Status  	    *int          `valid:"optional" json:"status,omitempty"`
		TrueName  		*string       `valid:"optional" json:"true_name,omitempty"`
		IdCard  		*string       `valid:"optional" json:"idcard,omitempty"`
		Phone  			*string       `valid:"optional" json:"phone,omitempty"`
		Province 		*string       `valid:"optional" json:"province,omitempty"`
//		DeliverAt 		*int64        `valid:"optional" json:"deliver_at,omitempty" `
		City    		*string       `valid:"optional" json:"city,omitempty" `
		Area    		*string       `valid:"optional" json:"area,omitempty"`
		Town    		*string       `valid:"optional" json:"town,omitempty" `
		Address    		*string       `valid:"optional" json:"address,omitempty" `
		PhoneOSTp    		*int       `valid:"optional" json:"device_os_tp,omitempty" `
		IP          	*string       `valid:"optional" json:"ip,omitempty" `
		MinIps          *int         `valid:"optional" json:"min_ips,omitempty" `

		StartCreatedAt  *int64    `valid:"optional" json:"start_created_at,omitempty"`
		EndCreatedAt    *int64    `valid:"optional" json:"end_created_at,omitempty"`
		StartDeliverAt  *int64    `valid:"optional" json:"start_deliver_at,omitempty"`
		EndDeliverAt    *int64    `valid:"optional" json:"end_deliver_at,omitempty"`
		LikeStr         *string   `valid:"optional" json:"like_str,omitempty"`

		Valid      *int       `valid:"optional"  json:"valid,omitempty"`
		Page       int64       `valid:"optional" json:"page,omitempty"`
		Size       int64       `valid:"optional" json:"size,omitempty"`
	}

	BkCardDateSheetList struct{
		Date              string     `valid:"optional"  json:"date,omitempty" ` //订单号
		OrderCount        int64      `valid:"optional"  json:"order_count,omitempty" `
		Tp                int        `valid:"optional"  json:"tp,omitempty"`
		ClassTp           *int        `valid:"optional"  json:"class_tp,omitempty"`


		StartCreatedAt  *int64    `valid:"optional" json:"start_created_at,omitempty"`
		EndCreatedAt    *int64    `valid:"optional" json:"end_created_at,omitempty"`
		Valid     *int       `valid:"optional"  json:"valid,omitempty"`
		Page       int64       `valid:"optional" json:"page,omitempty"`
		Size       int64       `valid:"optional" json:"size,omitempty"`
	}

	BkCardAreaSheetList struct{
		ClassTp    		*int     `valid:"optional" json:"class_tp,omitempty" `
		ClassISP        *int      `valid:"optional" json:"class_isp,omitempty"`
		CityCode          *string    `valid:"optional" json:"city_code,omitempty" `
		//StartCreatedAt  *int64    `valid:"optional" json:"start_created_at,omitempty"`
		//EndCreatedAt    *int64    `valid:"optional" json:"start_created_at,omitempty"`
		LikeStr         *string   `valid:"optional" json:"like_str,omitempty"`
		Valid            *int       `valid:"optional"  json:"valid,omitempty"`
		Page             int64       `valid:"optional" json:"page,omitempty"`
		Size             int64       `valid:"optional" json:"size,omitempty"`
	}

	BkPhoneNumberList struct{
		StartCreatedAt  *int64    `valid:"optional" json:"start_created_at,omitempty"`
		EndCreatedAt    *int64    `valid:"optional" json:"start_created_at,omitempty"`
		Number      *string    `valid:"optional"  json:"number,omitempty"`
		UseFlag    *int        `valid:"optional"  json:"use_flag,omitempty"`
		Valid      *int       `valid:"optional"  json:"valid,omitempty"`
		Page       int64       `valid:"optional" json:"page,omitempty"`
		Size       int64       `valid:"optional" json:"size,omitempty"`
	}

	BkPhoneNumber struct{
		Id              *int64       `valid:"required" json:"id,omitempty"`
		Number          *string    		`valid:"optional"  json:"number,omitempty"`
		UseFlag         *int        `valid:"optional" json:"use_flag,omitempty"`
		Level           *int        `valid:"optional" json:"level,omitempty"`
		Valid           *int       	`valid:"optional"  json:"valid,omitempty"`
	}

	BkPhoneNumberGet struct{
		Id              *int64       `valid:"optional" json:"id,omitempty"`
		Number      *string    		`valid:"optional"  json:"number,omitempty"`
	}

	BkPhoneNumberSave struct{
		Number          *string    		`valid:"required"  json:"number,omitempty"`
		UseFlag         int         `valid:"-" json:"-"`
		Level           *int        `valid:"required" json:"level,omitempty"`
		Valid           *int       	`valid:"-"  json:"-"`
	}

	BkResPhoneNumberSave struct{
		SuccCount           int   		`json:"succ_count,omitempty"`
		FailCount         int  		`json:"fail_count,omitempty"`
		SuccNumber          []*string    		`json:"succ_number,omitempty"`
		FailNumber          []*string    		`json:"fail_number,omitempty"`
	}

	BkPhoneNumberLevelAdd struct{
//		Id              *int64       `valid:"optional" json:"id,omitempty"`
		Desc            *string      `valid:"optional" json:"desc,omitempty" `
		Level           *int        `valid:"required" json:"level,omitempty"`
		Valid           *int       	`valid:"optional"  json:"valid,omitempty"`
	}

	BkPhoneNumberLevelUpdate struct{
		Id              *int64       `valid:"required" json:"id,omitempty"`
		Desc            *string      `valid:"optional"  json:"desc,omitempty" `
		Level           *int        `valid:"optional" json:"level,omitempty"`
		Valid           *int       	`valid:"optional"  json:"valid,omitempty"`
	}
	BkPhoneNumberLevelList struct{
		Id              *int64       `valid:"optional" json:"id,omitempty"`
		Desc            *string      `valid:"optional"  json:"desc,omitempty"`
		Level           *int        `valid:"optional" json:"level,omitempty"`
		Valid           *int       	`valid:"optional"  json:"valid,omitempty"`
		Page       int64       `valid:"required" json:"page,omitempty"`
		Size       int64       `valid:"optional" json:"size,omitempty"`
	}
	BkPhoneNumberLevelGets struct{
		Desc            *string      `valid:"optional"  json:"desc,omitempty"`
		Level           *int        `valid:"optional" json:"level,omitempty"`
		Valid           *int       	`valid:"optional"  json:"valid,omitempty"`
	}

	BkBlacklistPhoneAdd struct{
		Phone    	    *string     `valid:"required" json:"phone,omitempty"`
	}

	BkBlacklistPhone struct{
		Id              *int64       `valid:"required" json:"id,omitempty"`
		Phone    	    *string       `valid:"optional" json:"phone,omitempty"`
		Valid           *int       	`valid:"optional"  json:"valid,omitempty"`
	}

	BkBlacklistPhoneList struct{
		Id              *int64       `valid:"optional" json:"id,omitempty"`
		Phone    	    *string       `valid:"optional" json:"phone,omitempty"`
		Valid           *int       	`valid:"optional"  json:"valid,omitempty"`
		Page       int64       `valid:"required" json:"page,omitempty"`
		Size       int64       `valid:"optional" json:"size,omitempty"`
	}

	BkBlacklistIdcardAdd struct{
		IdCard    	    *string     `valid:"required" json:"idcard,omitempty"`
	}

	BkBlacklistIdcard struct{
		Id              *int64       `valid:"required" json:"id,omitempty"`
		IdCard    	    *string       `valid:"optional" json:"idcard,omitempty"`
		Valid           *int       	`valid:"optional"  json:"valid,omitempty"`
	}

	BkBlacklistIdcardList struct{
		Id              *int64       `valid:"optional" json:"id,omitempty"`
		IdCard    	    *string       `valid:"optional" json:"idcard,omitempty"`
		Valid           *int       	`valid:"optional"  json:"valid,omitempty"`
		Page       int64       `valid:"required" json:"page,omitempty"`
		Size       int64       `valid:"optional" json:"size,omitempty"`
	}

	BkCardOrderExtraImprot struct{
		OrderNo    	    *string     `valid:"required" json:"order_no,omitempty" ` //订单号
		Express    		*string    `valid:"optional" json:"express,omitempty"  ` //快递名称
		ExpressNo    	*string     `valid:"optional" json:"express_no,omitempty"  ` //快递单号
		ICCID          	*string    `valid:"optional" json:"ICCID,omitempty"`    //手机唯一识别码
		NewPhone  		*string  `valid:"optional" json:"new_phone,omitempty"`
		Guishudi       	*string    `valid:"optional"    json:"guishudi,omitempty" `

	}

	BkResCardOrderExtraImprot struct{
		SuccCount       	int    `valid:"optional"    json:"succ_count,omitempty" `
		FailCount       	int    `valid:"optional"    json:"fail_count,omitempty" `
	}
)