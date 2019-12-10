package async

import (
	"go-example/csp/servers/service"
	"time"
)

//User 设置缓冲通道，让User服务在处理完后，不被阻塞
func User(user string, elapsed time.Duration) chan string {
	retCh := make(chan string, 1)
	go func() {
		retCh <- service.User(user, elapsed)
	}()
	return retCh
}
