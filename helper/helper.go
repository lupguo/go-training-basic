package helper

import (
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"regexp"
	"runtime"
	"text/tabwriter"
	"time"
)

const (
	ToFile      = 0
	ToStdout    = 1
	LogFilename = "./storage/running/running.log"
)

type kv map[string]interface{}

// MonitorKey setting num goroutine
func monitorKey() kv {
	return kv{
		"NumGoroutine": runtime.NumGoroutine(),
	}
}

// GoStatus get running status background, and store into ./storage/running
func StatusLog(toDst int, interval time.Duration) {
	go statusTo(toDst, interval)
}

func statusTo(toDst int, interval time.Duration) {
	var dst io.Writer
	if toDst == ToFile {
		// running log
		dst = MustOpenFile(LogFilename)
	} else {
		dst = os.Stdout
	}
	// header write
	tw := tabwriter.NewWriter(dst, 0, 0, 2, ' ', 0)
	_, _ = fmt.Fprintf(tw, "Time\tKey\tVal\n---\t---\t---\n")
	// status write map
	for {
		for k, v := range monitorKey() {
			_, _ = fmt.Fprintf(tw, "%s\t%s\t%v\n", time.Now().Format(time.RFC3339), k, v)
		}
		//_, _ = fmt.Fprintf(tw, "---\t---\t---\n")
		_ = tw.Flush()
		time.Sleep(interval)
	}
}

func MustOpenFile(filename string) *os.File {
	dir := path.Dir(filename)
	if err := os.MkdirAll(dir, 0755); err != nil {
		log.Fatalln(err)
	}
	logFile, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalln(err)
	}
	return logFile
}

// TwShow help print format
func TwShow(tw *tabwriter.Writer) func(args ...interface{}) {
	lp := false
	show := func(args ...interface{}) {
		for _, v := range args {
			_, _ = fmt.Fprintf(tw, "%#v\t", v)
		}
		_, _ = tw.Write([]byte("\n"))

		// line print
		if lp == false {
			lp = true
			for range args {
				_, _ = tw.Write([]byte("----\t"))
			}
			_, _ = tw.Write([]byte("\n"))
		}
	}
	return show
}

// TwStdoutShow write to format
type ShowFun func(args ...interface{})

func TwStdoutShow() (tw *tabwriter.Writer, fn ShowFun) {
	tw = tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	return tw, TwShow(tw)
}

// RegexDebug show regex debug detail
func RegexDebug(mark string, re *regexp.Regexp, res string, showDetail bool) {
	log.Println(mark,re.MatchString(res), len(re.FindAllString(res, -1)))
	if showDetail {
		log.Println(re.FindAllString(res, -1))
	}
}
