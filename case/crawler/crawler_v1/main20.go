package main

import (
	"fmt"
	"go-example/case/crawler/crawler_v1/internal/links"
	"log"
)

func main() {
	worklist := make(chan []string)  // lists of URLs, may have duplicates
	unseenLinks := make(chan string) // de-duplicated URLs

	// Add command-line arguments to worklist.
	go func() {
		//worklist <- os.Args[1:]
		worklist <- []string{
			`https://tkstorm.com`,
		}
	}()

	// Create 20 crawler goroutines to fetch each unseen link.
	for i := 0; i < 20; i++ {
		go func(i int) {
			for link := range unseenLinks {
				//worklist <- crawl(link) // bad!! if you wirte direct, this goroutine will be sync, then dead lock!
				foundLinks := crawl(link)
				go func() {
					worklist <- foundLinks
				}()
			}
			log.Println("close goroutine #", i)
		}(i)
	}

	// The main goroutine de-duplicates worklist items
	// and sends the unseen ones to the crawlers.
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				unseenLinks <- link
			}
		}
	}

}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

// todo
// can not break goroutine
