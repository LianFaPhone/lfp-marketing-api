package order

import (
	"LianFaPhone/lfp-marketing-api/models"
)

//func GetSelectFields(input string) []string{
//	if len(input) <= 0 || input == " "{
//		return nil
//	}
//	res := make([]string, 0, len(input))
//	inArr := strings.Split(input, ",")
//	for i:=0; i < len(inArr);i++{
//		v, ok := models.CardOrderFieldMap[inArr[i]]
//		if !ok {
//			continue
//		}
//		res = append(res, v)
//	}
//	return res
//}

// userid--->GoodsId---->orders
func GetsGoodsByUserId(input string) ([]*models.PdPartnerGoods,error) {
	if len(input) <= 0 || input == " "{
		return nil,nil
	}

	return new(models.PdPartnerGoods).GetsByUserId(input)
}
