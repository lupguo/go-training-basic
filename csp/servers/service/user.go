package service

import (
	"time"
)

//User Service User
func User(str string, dur time.Duration) string {
	time.Sleep(dur)
	return "User Name:" + str
}

