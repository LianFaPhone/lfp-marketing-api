package tasker

import "time"

func GenDay2(tt int64) int64 {
	ut := time.Unix(tt, 0)
	t1 := ut.Year()  //年
	t2 := ut.Month() //月
	t3 := ut.Day()   //日
	loc := time.FixedZone("UTC", 8*3600)
	currentTimeData := time.Date(t1, t2, t3, 0, 0, 0, 0, loc)
	return currentTimeData.Unix()
}

func GenDay(ut time.Time) int64 {
	t1 := ut.Year()  //年
	t2 := ut.Month() //月
	t3 := ut.Day()   //日
	loc := time.FixedZone("UTC", 8*3600)
	currentTimeData := time.Date(t1, t2, t3, 0, 0, 0, 0, loc)
	return currentTimeData.Unix()
}


