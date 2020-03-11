package models

const(
	CONST_ADTRACK_Tp_KuaiShou = 1
	CONST_ADTRACK_Tp_DouYin = 2
)

type AdTp struct {
	Tp int    `json:"tp"`
	Name   string `json:"name"`
}

var AdTpArr [] AdTp = []AdTp{
	AdTp{CONST_ADTRACK_Tp_KuaiShou, "快手"},
	AdTp{CONST_ADTRACK_Tp_DouYin, "抖音"},
}

var AdTpMap map[int] string = map[int] string{
	CONST_ADTRACK_Tp_KuaiShou:"快手",
	CONST_ADTRACK_Tp_DouYin: "抖音",
}