package main

import (
	"fmt"
	"go-example/helper"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func init() {
	fmt.Println("init")
	runtime.GOMAXPROCS(1)
	rand.Seed(time.Now().UnixNano())
	helper.StatusLog(helper.ToFile, 200*time.Millisecond)
}

func main() {
	fmt.Println("main")
	wg := sync.WaitGroup{}
	fin := make(chan struct{})

	// wrong wg.wait here! wg.wait may run over! (when this goroutine run before wg.Add())
	//go func() {
	//	fmt.Println("run.wg.Wait 1, wait for work fin!")
	//	wg.Wait()
	//	fmt.Println("run.wg.Wait 2, work has finished!!")
	//	// fin <- struct{}{}
	//	close(fin)
	//	fmt.Println("run.wg.Wait 3!")
	//}()

	for i := 0; i < 5; i++ {
		wg.Add(1)
		//time.Sleep(500*time.Millisecond)
		go func(i int) {
			//time.Sleep(time.Duration(rand.Intn(10)*100) * time.Millisecond)
			fmt.Printf("do work %d\n", i)
			wg.Done()
		}(i)
	}

	// right wg.wait here!
	go func() {
		fmt.Println("run.wg.Wait 1, wait for work fin!")
		wg.Wait()
		fmt.Println("run.wg.Wait 2, work has finished!!")
		// using close(chan) broadcast chan is closed
		// fin <- struct{}{}
		close(fin)
		fmt.Println("run.wg.Wait 3!")
	}()

	// receive fin
	fmt.Println("receive fin chan... chan block wait for event!")
	if _, ok := <-fin; ok {
		fmt.Println("fin receive ok !")
	}else{
		fmt.Println("fin is close")
	}

	// receive again (hang stop because fin is block, if using `fin <- struct{}{}`)
	if _, ok := <-fin; ok {
		fmt.Println("fin receive ok !")
	}else{
		fmt.Println("fin is close")
	}

	//time.Sleep(time.Second)
	fmt.Println("finish main!")
}
