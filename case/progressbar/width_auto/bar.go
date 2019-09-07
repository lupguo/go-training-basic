package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"golang.org/x/sys/unix"
)

var (
	total = 50
	count = 0
	wscol = 20
)

func init() {
	err := updateWSCol()
	if err != nil {
		panic(err)
	}
}

func updateWSCol() error {
	ws, err := unix.IoctlGetWinsize(syscall.Stdout, unix.TIOCGWINSZ)
	if err != nil {
		return err
	}
	wscol = int(ws.Col)
	return nil
}

func renderbar() {
	fmt.Print("\x1b7")       // 保存光标位置
	fmt.Print("\x1b[2k")     // 清除当前行内容
	defer fmt.Print("\x1b8") // 恢复光标位置

	barwidth := wscol - len("Progress: 100% []")
	done := int(float64(barwidth) * float64(count) / float64(total))

	fmt.Printf("Progress: \x1b[33m%3d%%\x1b[0m ", count*100/total)
	fmt.Printf("[%s%s]",
		strings.Repeat("=", done),
		strings.Repeat("-", barwidth-done))
}

func main() {
	// set signal handler
	sigwinch := make(chan os.Signal, 1)
	defer close(sigwinch)
	signal.Notify(sigwinch, syscall.SIGWINCH)
	go func() {
		for {
			if _, ok := <-sigwinch; !ok {
				return
			}
			_ = updateWSCol()
			renderbar()
		}
	}()

	for count = 1; count <= 50; count++ {
		renderbar()
		time.Sleep(time.Second)
	}
	fmt.Println()
}