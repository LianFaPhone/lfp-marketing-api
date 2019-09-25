package controllers

import (
	"github.com/shopspring/decimal"
	"strings"
	"math"
	"sync"
)

const(
	CONST_MIN_RED_DIV_BASE = 10
)

func GetFileExt(name string ) string {
	idx := strings.Index(name, ".")
	if idx < 0 {
		return ""
	}
	return name[idx:]
}
//
//func GetClaimsFromCtx(ctx iris.Context) *ApiKeyClaims {
//	cl := ctx.Values().Get("apikey_claims")
//	if cl == nil {
//		return nil
//	}
//	claims,_ := cl.(*ApiKeyClaims)
//	if claims == nil {
//		return nil
//	}
//	return claims
//}

//func GenApiKeyUuidSqlCondFromCtx(ctx iris.Context) *models.SqlPairCondition {
//	claims := GetClaimsFromCtx(ctx)
//	if claims == nil {
//		return nil
//	}
//	return &models.SqlPairCondition{"uuid in (?)", claims}
//}
//
//func GenApiKeyActyUuidSqlCondFromCtx(ctx iris.Context) *models.SqlPairCondition {
//	claims := GetClaimsFromCtx(ctx)
//	if claims == nil {
//		return nil
//	}
//	return &models.SqlPairCondition{"activity_uuid in (?)", claims}
//}



func GenMinPrecision(totalCoin *decimal.Decimal, totalRed, totalRob *int64){
	testPrecision := totalCoin.Div(decimal.New((*totalRed)*(*totalRob)*CONST_MIN_RED_DIV_BASE, 0))
	n := decimal.New(1, 0).Div(testPrecision)
	f,_ := n.Float64()

	math.Ceil(math.Log10(f))

}

func Substr(str string, start int, end int) string {
	rs := []rune(str)
	length := len(rs)

	if start < 0 || start > length {
		return ""
	}

	if end < 0 || end > length {
		return ""
	}

	return string(rs[start:end])
}


func SecretPhone(phone string) string {
	if phone == "" {
		return ""
	}

	return Substr(phone, 0, 3) + "***" + Substr(phone, len(phone)-4, len(phone))
}

type IdGener struct{
	id  int
	sync.Mutex
}

func (this * IdGener) Gen() int {
	data := 0
	this.Lock()
	this.id++
	data = (this.id)%1000
	this.Unlock()
	return data
}