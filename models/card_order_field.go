package models


type CardOrderField struct{
	Id     int    `json:"id"`
	Field   string `json:"field"`
	Name   string `json:"name"`
}

var CardOrderFieldMap = map[string] string {
	"1": "order_no",
	"2": "isp",
	"3": "partner_id",
	"4": "partner_goods_code",
	"5": "true_name",
	"6": "status",
	"7": "idcard",
	"8": "phone",
	"9": "express",
	"10": "express_no",
	"11": "express_remark",
	"12": "deliver_at",
	"13": "ICCID",
	"14": "new_phone",
	"15": "guishudi",
	"16": "ad_tp",
	"17": "ip",
	"18": "third_order_no",
	"19": "third_order_at",
	"20": "active_at",
	"21": "valid",
	"22": "created_at",
	"23": "updated_at",
	"24": "province",
	"25": "province_code",
	"26": "city",
	"27": "city_code",
	"28": "area",
	"29": "area_code",
	"30": "town",
	"31": "address",
}

var CardOrderFieldArr = []*CardOrderField{
	&CardOrderField{1, "order_no", "订单号"},
	&CardOrderField{2, "isp", "运营商"},
	&CardOrderField{3, "partner_id", "渠道商ID"},
	&CardOrderField{4, "partner_goods_code", "渠道商品码"},
	&CardOrderField{5, "true_name", "姓名"},
	&CardOrderField{6, "status", "订单状态"},
	&CardOrderField{7, "idcard", "身份证"},
	&CardOrderField{8, "phone",   "手机号"},
	&CardOrderField{9, "express",   "快递公司"},
	&CardOrderField{10, "express_no",   "快递单号"},
	&CardOrderField{11, "express_remark",   "备注"},
	&CardOrderField{12, "deliver_at",    "发货时间"},
	&CardOrderField{13, "ICCID",    "ICCID"},
	&CardOrderField{14, "new_phone",     "新手机号"},
	&CardOrderField{15, "guishudi", "归属地"},
	&CardOrderField{16, "ad_tp",     "广告类型"},
	&CardOrderField{17, "ip",   "ip"},
	&CardOrderField{18, "third_order_no", "第三方下单号"},
	&CardOrderField{19, "third_order_at",    "第三方下单时间"},
	&CardOrderField{20, "active_at",     "激活时间"},
	&CardOrderField{21, "valid",    "有效性"},
	&CardOrderField{22, "created_at", "订单创建时间"},
	&CardOrderField{23, "updated_at",   "订单更新时间"},
	&CardOrderField{24, "province",   "省"},
	&CardOrderField{25, "province_code",  "省码"},
	&CardOrderField{26, "city",    "城市"},
	&CardOrderField{27, "city_code",    "城市码"},
	&CardOrderField{28, "area", "区"},
	&CardOrderField{29, "area_code", "区码"},
	&CardOrderField{30, "town", "镇"},
	&CardOrderField{31, "address", "详细地址"},

}
