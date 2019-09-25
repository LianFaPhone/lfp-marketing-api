package models

type ClassTp struct{
	ISP                  int      `json:"isp,omitempty"`
	Tp         		     int      `json:"tp,omitempty"` //加上type:int(11)后AUTO_INCREMENT无效
	Name                 string   `json:"name,omitempty"`
	Alias                string   `json:"alias,omitempty"`  //拼音首字母缩写
}

var ClassTpArr  []*ClassTp
var ClassTpMap  map[int] *ClassTp

const(
	CONST_ISP_UnKnown = 0
	CONST_ISP_Dianxin = 1
	CONST_ISP_YiDong = 2
	CONST_ISP_LiTong = 3
	CONST_ISP_Ali = 4
	CONST_ISP_JD = 5

)

//以后得搞成配置文件才行
func init(){
	ClassTpArr = make([]*ClassTp, 0)
	ClassTpMap = make(map[int] *ClassTp)

	temp := &ClassTp{CONST_ISP_UnKnown,1, "王卡", "wk"}
	ClassTpArr = append(ClassTpArr, temp)
	ClassTpMap[temp.Tp] = temp

	temp = &ClassTp{CONST_ISP_UnKnown,2, "小鱼卡", "xyk"}
	ClassTpArr = append(ClassTpArr, temp)
	ClassTpMap[temp.Tp] = temp

	temp = &ClassTp{CONST_ISP_LiTong,3, "联通超王卡", "ltcwk"}
	ClassTpArr = append(ClassTpArr, temp)
	ClassTpMap[temp.Tp] = temp

	temp = &ClassTp{CONST_ISP_UnKnown,4, "鱼卡-抖音", "ykdy"}
	ClassTpArr = append(ClassTpArr, temp)
	ClassTpMap[temp.Tp] = temp

	temp =	&ClassTp{CONST_ISP_LiTong,5, "联通大王卡", "ltdwk"}
	ClassTpArr = append(ClassTpArr, temp)
	ClassTpMap[temp.Tp] = temp

	temp =	&ClassTp{CONST_ISP_LiTong,6, "联通大王卡-抖音", "ltdwkdy"}
	ClassTpArr = append(ClassTpArr, temp)
	ClassTpMap[temp.Tp] = temp

	temp =	&ClassTp{CONST_ISP_LiTong,7, "联通大王卡1", "ltdwk1"}
	ClassTpArr = append(ClassTpArr, temp)
	ClassTpMap[temp.Tp] = temp

	temp =	&ClassTp{CONST_ISP_LiTong,8, "联通大王卡2", "ltdwk2"}
	ClassTpArr = append(ClassTpArr, temp)
	ClassTpMap[temp.Tp] = temp

	temp =	&ClassTp{CONST_ISP_LiTong,9, "联通导学号", "ltdxh"}
	ClassTpArr = append(ClassTpArr, temp)
	ClassTpMap[temp.Tp] = temp

	temp =	&ClassTp{CONST_ISP_LiTong,10, "电信High卡", "dxhighk"}
	ClassTpArr = append(ClassTpArr, temp)
	ClassTpMap[temp.Tp] = temp

	temp =	&ClassTp{CONST_ISP_LiTong,11, "联通导学号2", "ltdxh2"}
	ClassTpArr = append(ClassTpArr, temp)
	ClassTpMap[temp.Tp] = temp

	temp =	&ClassTp{CONST_ISP_LiTong,12, "联通大王卡-抖音2", "ltdwkdy2"}
	ClassTpArr = append(ClassTpArr, temp)
	ClassTpMap[temp.Tp] = temp

	temp =	&ClassTp{CONST_ISP_LiTong,13, "联通大王卡-抖音3", "ltdwkdy3"}
	ClassTpArr = append(ClassTpArr, temp)
	ClassTpMap[temp.Tp] = temp


	temp =	&ClassTp{CONST_ISP_LiTong,14, "联通大王卡-抖音4", "ltdwkdy4"}
	ClassTpArr = append(ClassTpArr, temp)
	ClassTpMap[temp.Tp] = temp

	temp =	&ClassTp{CONST_ISP_LiTong,15, "联通大王卡-抖音5", "ltdwkdy5"}
	ClassTpArr = append(ClassTpArr, temp)
	ClassTpMap[temp.Tp] = temp

	temp =	&ClassTp{CONST_ISP_LiTong,16, "联通大王卡-抖音6", "ltdwkdy6"}
	ClassTpArr = append(ClassTpArr, temp)
	ClassTpMap[temp.Tp] = temp


	temp =	&ClassTp{CONST_ISP_LiTong,17, "联通大王卡-抖音55", "ltdwkdy55"}
	ClassTpArr = append(ClassTpArr, temp)
	ClassTpMap[temp.Tp] = temp


	temp =	&ClassTp{CONST_ISP_LiTong,18, "联通大王卡-抖音66", "ltdwkdy66"}
	ClassTpArr = append(ClassTpArr, temp)
	ClassTpMap[temp.Tp] = temp

	temp =	&ClassTp{CONST_ISP_YiDong,19, "移动花卡-快手", "ydhkks"}
	ClassTpArr = append(ClassTpArr, temp)
	ClassTpMap[temp.Tp] = temp

	temp =	&ClassTp{CONST_ISP_YiDong,20, "移动花卡-抖音", "ydhkdy"}
	ClassTpArr = append(ClassTpArr, temp)
	ClassTpMap[temp.Tp] = temp

	temp =	&ClassTp{CONST_ISP_YiDong,21, "移动花卡-抖音1", "ydhkdy1"}
	ClassTpArr = append(ClassTpArr, temp)
	ClassTpMap[temp.Tp] = temp

	temp =	&ClassTp{CONST_ISP_YiDong,22, "移动花卡-抖音2", "ydhkdy2"}
	ClassTpArr = append(ClassTpArr, temp)
	ClassTpMap[temp.Tp] = temp

	temp =	&ClassTp{CONST_ISP_YiDong,23, "移动花卡-4", "ydhk4"}
	ClassTpArr = append(ClassTpArr, temp)
	ClassTpMap[temp.Tp] = temp


	temp =	&ClassTp{CONST_ISP_YiDong,24, "移动花卡-5", "ydhk5"}
	ClassTpArr = append(ClassTpArr, temp)
	ClassTpMap[temp.Tp] = temp


	temp =	&ClassTp{CONST_ISP_YiDong,25, "移动花卡-6", "ydhk6"}
	ClassTpArr = append(ClassTpArr, temp)
	ClassTpMap[temp.Tp] = temp

	temp =	&ClassTp{CONST_ISP_UnKnown,26, "大圣卡BPS", "dskbps"}
	ClassTpArr = append(ClassTpArr, temp)
	ClassTpMap[temp.Tp] = temp


	temp =	&ClassTp{CONST_ISP_UnKnown,27, "大圣卡BPS-快手", "dskbpsks"}
	ClassTpArr = append(ClassTpArr, temp)
	ClassTpMap[temp.Tp] = temp


	temp =	&ClassTp{CONST_ISP_UnKnown,28, "大圣卡BPS-快手1", "dskbpsks1"}
	ClassTpArr = append(ClassTpArr, temp)
	ClassTpMap[temp.Tp] = temp


	temp =	&ClassTp{CONST_ISP_UnKnown,29, "大圣卡BPS-抖音", "dskbpsdy"}
	ClassTpArr = append(ClassTpArr, temp)
	ClassTpMap[temp.Tp] = temp

	temp =	&ClassTp{CONST_ISP_UnKnown,30, "大圣卡BPS-抖音1", "dskbpsdy1"}
	ClassTpArr = append(ClassTpArr, temp)
	ClassTpMap[temp.Tp] = temp

	temp =	&ClassTp{CONST_ISP_UnKnown,31, "大圣卡BPS-抖音2", "dskbpsdy2"}
	ClassTpArr = append(ClassTpArr, temp)
	ClassTpMap[temp.Tp] = temp

}

