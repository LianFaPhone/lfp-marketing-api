package models

type OrderStatus struct {
	Status int    `json:"status"`
	Name   string `json:"name"`
}

const (
	CONST_OrderStatus_New  = 1     //新订单
	CONST_OrderStatus_Name = "新订单" //

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
)

var PathToOrderStatus = map[string] int {
	"all": -1,
	"new": 1 ,
	"export": 2 ,
	"deliver": 3 ,
	"waitdone": 5 ,
	"alreadydone": 6 ,
	"recyclebin": 4 ,
	"unmatch": 7 ,
	"activated": 8 ,
}

var OrderStatusMap = map[int]string{
	CONST_OrderStatus_New:               CONST_OrderStatus_Name,
	CONST_OrderStatus_Already_Export:    CONST_OrderStatus_Already_Export_Name,
	CONST_OrderStatus_Already_Delivered: CONST_OrderStatus_Already_Delivered_Name,
	CONST_OrderStatus_Wait_Done:         CONST_OrderStatus_Wait_Done_Name,
	CONST_OrderStatus_Already_Done:      CONST_OrderStatus_Already_Done_Name,
	CONST_OrderStatus_Recyclebin:        CONST_OrderStatus_Recyclebin_Name,
	CONST_OrderStatus_UnMatch:           CONST_OrderStatus_UnMatch_Name,
	CONST_OrderStatus_Already_Activated: CONST_OrderStatus_Already_Activated_Name,
}

var OrderStatusArr = []*OrderStatus{
	&OrderStatus{CONST_OrderStatus_New, CONST_OrderStatus_Name},
	&OrderStatus{CONST_OrderStatus_Already_Export, CONST_OrderStatus_Already_Export_Name},
	&OrderStatus{CONST_OrderStatus_Already_Delivered, CONST_OrderStatus_Already_Delivered_Name},
	&OrderStatus{CONST_OrderStatus_Wait_Done, CONST_OrderStatus_Wait_Done_Name},
	&OrderStatus{CONST_OrderStatus_Already_Done, CONST_OrderStatus_Already_Done_Name},
	&OrderStatus{CONST_OrderStatus_Recyclebin, CONST_OrderStatus_Recyclebin_Name},
	&OrderStatus{CONST_OrderStatus_UnMatch, CONST_OrderStatus_UnMatch_Name},
	&OrderStatus{CONST_OrderStatus_Already_Activated, CONST_OrderStatus_Already_Activated_Name},
}
