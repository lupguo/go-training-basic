package async

import (
	"go-example/csp/servers/service"
	"testing"
	"time"
)

func TestServices(t *testing.T) {
	var user string
	userCh := User("Clark", time.Millisecond*500)
	goods := service.Goods("Apple", time.Millisecond*500)
	select {
	case user = <-userCh:
	case <-time.After(600 * time.Millisecond):
		t.Error("time out")
	}
	t.Log(user, goods)
}
