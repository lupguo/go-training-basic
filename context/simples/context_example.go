package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 模拟老板安排一个工程，限定3s的时限，否则就取消没有完成的工作
//
// 工头盘点：目前每个工人处理任务时限在100~300ms内，计划20人全部来参与
// 工头尝试：
// 		1. 时间限制：老板限定3s内完成任务，否则直接取消剩余的工作内容
//		2. 串行派遣任务，存在很多工人的超时
//		3. 并行派遣任务，总完成效率高很多
//		4. 并行派遣任务，改善工人工作效率(优化到200s内)，总体达标
//
func main() {
	// 空的上下文
	ctx := context.Background()

	// 1. 这里只是为了更好的体现，cancelFn，超时限制可以直接：ctx, _ = content.WithTimeout(ctx, 3*time.Second)
	// 取消的上下文（包含一个取消函数）
	ctx, cancelFn := context.WithCancel(ctx)

	// 整体程序，限定时间3s钟，主动取消
	go func() {
		time.Sleep(3 * time.Second)
		cancelFn()
	}()

	// 2. 串行地派遣20个工作任务，整体限时3s钟，后续的任务都会超时！
	//for i := 1; i <= 20; i++ {
	//	dispatchWork(ctx, i)
	//}

	// 3. 并发的启动20个goroutine，整体限时3s钟，虽然有的工人工作超时，但整体完成任务较多
	wg := sync.WaitGroup{}
	for i := 1; i <= 20; i++ {
		wg.Add(1)
		go func(i int) {
			dispatchWork(ctx, i)
			wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Println("all over!!")
}

// 工头派遣工作
func dispatchWork(ctx context.Context, workNo int) {

	// 工头中继老板的ctx上下文，在总时间限制前提下，继续把每项工作内容时间限制在300ms指定范围内
	timeoutCtx, _ := context.WithTimeout(ctx, 300*time.Millisecond)

	// 工作完成标识
	finish := make(chan string)

	// 派遣工作给工人
	go hardWork(workNo, finish)

	// 工头等待消息，可能工人完成，也可能单次工作超时了，亦或是整体达到timeout限制时间了
	// tips: 可以调整工人的工作时间看效果
	select {
	case fin := <-finish:
		fmt.Println("headman : ", fin)
	case <-timeoutCtx.Done():
		fmt.Printf("headman : work(%d) too slow, over time limit!!\n", workNo)

		// 关闭finish通道，避免worker的goroutine泄露
		//close(finish)
	}
}

// 工人开始workNo工作，可能很重的工作，如果完成，则将内容通知到finish通道
func hardWork(workNo int, finish chan<- string) {
	rand.Seed(time.Now().UnixNano())

	// 4. 模拟工作，花时100~500ms，当工人工作时间都是200ms以下，整体工作都可以达标
	// tips: 可以调整工人的工作时间看效果
	workExpend := time.Duration(rand.Intn(3)+1) * 100 * time.Millisecond
	time.Sleep(workExpend)

	// 通知派遣工头，完成了任务了
	finish <- fmt.Sprintf("ok, finished no. %d work.", workNo)
}
