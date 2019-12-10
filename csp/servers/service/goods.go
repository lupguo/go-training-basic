package service

import (
	"time"
)

//User Service Goods
func Goods(str string, dur time.Duration) string {
	time.Sleep(dur)
	return "Goods Info:" + str
}

