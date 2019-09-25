package models


type PhoneOs struct{
	Tp    int   `json:"tp,omitempty"`
	Name  string   `json:"name,omitempty"`
}

const(
	CONST_PHONEOS_Other   = 4
	CONST_PHONEOS_Other_Name   = "Other"
	CONST_PHONEOS_Android = 1
	CONST_PHONEOS_Android_Name = "Android"
	CONST_PHONEOS_Iphone = 2
	CONST_PHONEOS_Iphone_Name = "Iphone"
	CONST_PHONEOS_Ipad   = 3
	CONST_PHONEOS_Ipad_Name   = "Ipad"
)

var PhoneOsArr  []*PhoneOs
var PhoneOsMap  map[int] *PhoneOs

func init() {
	PhoneOsArr = make([]*PhoneOs, 0)
	PhoneOsMap = make(map[int] *PhoneOs)

	tmp := &PhoneOs{CONST_PHONEOS_Android, CONST_PHONEOS_Android_Name}
	PhoneOsArr = append(PhoneOsArr, tmp)
	PhoneOsMap[CONST_PHONEOS_Android] = tmp

	tmp = &PhoneOs{CONST_PHONEOS_Iphone, CONST_PHONEOS_Iphone_Name}
	PhoneOsArr = append(PhoneOsArr, tmp)
	PhoneOsMap[CONST_PHONEOS_Iphone] = tmp

	tmp = &PhoneOs{CONST_PHONEOS_Ipad, CONST_PHONEOS_Ipad_Name}
	PhoneOsArr = append(PhoneOsArr, tmp)
	PhoneOsMap[CONST_PHONEOS_Ipad] = tmp

	tmp = &PhoneOs{CONST_PHONEOS_Other, CONST_PHONEOS_Other_Name}
	PhoneOsArr = append(PhoneOsArr, tmp)
	PhoneOsMap[CONST_PHONEOS_Other] = tmp
}

