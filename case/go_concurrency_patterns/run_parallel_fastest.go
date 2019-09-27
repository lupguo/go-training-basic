package main

import (
	"context"
	"fmt"
	"math/rand"
	"runtime"
	"strconv"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	conns := make([]Conn, 5)
	for i := 0; i < len(conns); i++ {
		conns[i] = Conn{i}
	}

	for n := 0; n < 10; n++ {
		rs := Query(conns, "query="+strconv.Itoa(n*n))
		fastConn := rs.ctx.Value("c").(*Conn)
		fmt.Printf("%s, on connId:%d, elapse:%v\n", rs.query, fastConn.connId, rs.elapse)
	}

	time.Sleep(100 * time.Millisecond)
	fmt.Println("GouroutineNums=>", runtime.NumGoroutine())
}

type Conn struct {
	connId int
}

type Result struct {
	query  string
	elapse time.Duration
	ctx    context.Context
}

func (c *Conn) DoQuery(query string) Result {
	ctx := context.Background()
	elapse := time.Duration(rand.Intn(50)) * time.Millisecond
	time.Sleep(elapse)
	return Result{
		query:  query,
		elapse: elapse,
		ctx:    context.WithValue(ctx, "c", c),
	}
}

func Query(conns []Conn, query string) Result {
	ch := make(chan Result)
	for _, conn := range conns {
		// It must be no blocked (default:) or Goroutine will leak
		go func(c Conn) {
			select {
			case ch <- c.DoQuery(query):
			default:
			}
		}(conn)
	}
	return <-ch
}
