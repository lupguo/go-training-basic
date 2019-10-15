package main

import (
	"flag"
	"fmt"
	newrelic "github.com/newrelic/go-agent"
	"io/ioutil"
	"log"
	"os"
	"path"
	"runtime"
	"runtime/pprof"
	"sync"
	"time"
)

var (
	wg   sync.WaitGroup
	sema = make(chan struct{}, 20)
)

var root = flag.String("root", ".", "path to du")
var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
var memprofile = flag.String("memprofile", "", "write memory profile to `file`")




func main() {
	defer elapsed(time.Now())
	flag.Parse()

	// cpu profile
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		defer f.Close()
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	// new relic
	config := newrelic.NewConfig("YOUR_APP_NAME", "_YOUR_NEW_RELIC_LICENSE_KEY_")
	app, err := newrelic.NewApplication(config)

	filesize := make(chan int64)
	wg.Add(1)
	go func() {
		workDir(*root, filesize)
	}()

	go func() {
		wg.Wait()
		close(filesize)
	}()

	var sumsize int64
	for size := range filesize {
		sumsize += size
	}
	fmt.Printf("total sum size: %0.2f Mb\n", float64(sumsize)/1e6)

	// memory profile
	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		defer f.Close()
		runtime.GC() // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
	}
}

func elapsed(start time.Time) {
	fmt.Printf("program elapsed: %0.2fs\n", time.Since(start).Seconds())
}

func workDir(dir string, filesize chan int64) {
	defer wg.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			wg.Add(1)
			go workDir(path.Join(dir, entry.Name()), filesize)
		} else {
			filesize <- entry.Size()
		}
	}
}

func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}
	defer func() {
		<-sema
	}()
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du read dir error: %v\n", err)
		return nil
	}
	return entries
}
