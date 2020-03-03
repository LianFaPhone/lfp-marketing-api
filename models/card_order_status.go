package models

type OrderStatus struct {
	Status int    `json:"status"`
	Name   string `json:"name"`
}

const (
	CONST_OrderStatus_New  = 1     //新订单
	CONST_OrderStatus_Name = "新订单已完成" //


	CONST_OrderStatus_New_UnFinish  = 10     //新订单
	CONST_OrderStatus_New_UnFinish_Name = "新订单未完成" //

	CONST_OrderStatus_Already_Export      = 2     //已导出
	CONST_OrderStatus_Already_Export_Name = "已导出" //

	CONST_OrderStatus_Already_Delivered      = 3     //已发货
	CONST_OrderStatus_Already_Delivered_Name = "已发货" //

	CONST_OrderStatus_Wait_Done      = 5     //待处理
	CONST_OrderStatus_Wait_Done_Name = "待处理" //

	CONST_OrderStatus_Already_Done      = 6     //已处理
	CONST_OrderStatus_Already_Done_Name = "已处理" //

	CONST_OrderStatus_Recyclebin      = 4     //垃圾箱
	CONST_OrderStatus_Recyclebin_Name = "回收站" //

	CONST_OrderStatus_UnMatch      = 7     //未匹配
	CONST_OrderStatus_UnMatch_Name = "未匹配" //

	CONST_OrderStatus_Already_Activated      = 8     //已激活
	CONST_OrderStatus_Already_Activated_Name = "已激活" //已激活

	CONST_OrderStatus_Fail  = 11     //新订单
	CONST_OrderStatus_Fail_Name = "失败" //
)

var PathToOrderStatus = map[string] int {
	"all": -1,
	"new": CONST_OrderStatus_New ,
	"new-unfinish":CONST_OrderStatus_New_UnFinish,
	"export": CONST_OrderStatus_Already_Export ,
	"deliver": CONST_OrderStatus_Already_Delivered ,
	"waitdone": CONST_OrderStatus_Wait_Done ,
	"alreadydone": CONST_OrderStatus_Already_Done ,
	"recyclebin": CONST_OrderStatus_Recyclebin ,
	"unmatch": CONST_OrderStatus_UnMatch ,
	"activated": CONST_OrderStatus_Already_Activated ,
	"fail":CONST_OrderStatus_Fail,
}

var OrderStatusMap = map[int]string{
	CONST_OrderStatus_New_UnFinish:      CONST_OrderStatus_New_UnFinish_Name,
	CONST_OrderStatus_New:               CONST_OrderStatus_Name,
	CONST_OrderStatus_Already_Export:    CONST_OrderStatus_Already_Export_Name,
	CONST_OrderStatus_Already_Delivered: CONST_OrderStatus_Already_Delivered_Name,
	CONST_OrderStatus_Wait_Done:         CONST_OrderStatus_Wait_Done_Name,
	CONST_OrderStatus_Already_Done:      CONST_OrderStatus_Already_Done_Name,
	CONST_OrderStatus_Recyclebin:        CONST_OrderStatus_Recyclebin_Name,
	CONST_OrderStatus_UnMatch:           CONST_OrderStatus_UnMatch_Name,
	CONST_OrderStatus_Already_Activated: CONST_OrderStatus_Already_Activated_Name,
	CONST_OrderStatus_Fail:              CONST_OrderStatus_Fail_Name,
}

var OrderStatusArr = []*OrderStatus{
	&OrderStatus{CONST_OrderStatus_New_UnFinish, CONST_OrderStatus_New_UnFinish_Name},
	&OrderStatus{CONST_OrderStatus_New, CONST_OrderStatus_Name},
	&OrderStatus{CONST_OrderStatus_Already_Export, CONST_OrderStatus_Already_Export_Name},
	&OrderStatus{CONST_OrderStatus_Already_Delivered, CONST_OrderStatus_Already_Delivered_Name},
	&OrderStatus{CONST_OrderStatus_Wait_Done, CONST_OrderStatus_Wait_Done_Name},
	&OrderStatus{CONST_OrderStatus_Already_Done, CONST_OrderStatus_Already_Done_Name},
	&OrderStatus{CONST_OrderStatus_Recyclebin, CONST_OrderStatus_Recyclebin_Name},
	&OrderStatus{CONST_OrderStatus_UnMatch, CONST_OrderStatus_UnMatch_Name},
	&OrderStatus{CONST_OrderStatus_Already_Activated, CONST_OrderStatus_Already_Activated_Name},
	&OrderStatus{CONST_OrderStatus_Fail, CONST_OrderStatus_Fail_Name},
}
