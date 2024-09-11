package utils

import "time"

var NowTime func() time.Time 

func init() {
	jst, err := time.LoadLocation("Asia/Tokyo")
    if err != nil {
        panic(err)
    }

	NowTime = func() time.Time {
		return time.Now().In(jst)
	}
}
