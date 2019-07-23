package main

import (
	"fmt"
	"go-example/case/tutorialspoint/internal/body"
	"go-example/helper"
	"log"
	"regexp"
	"sync"
)

// Page(url, regex)
type Page struct {
	url string
	re  *regexp.Regexp
}

// Analyze IndexPage using regex got match content, write to chan
func (p *Page) AnalyzeLinks() (urls []string, err error) {
	return body.FindLeft(p.url, p.re)
	//return body.FindLeftMock(p.url, p.re)
}

// Analyze IndexPage using regex got match content, write to chan
func (p *Page) AnalyzeContent() (content []byte, err error) {
	return body.FindContent(p.url, p.re)
}

// Content
func content(url string) (data []byte, err error) {
	contentPage := &Page{
		url,
		regexp.MustCompile(`(?smU)<div class="content">.*</div>`),
	}
	return contentPage.AnalyzeContent()
}

// Launch
func launch(url string) (urls []string, err error) {
	indexPage := &Page{
		url,
		regexp.MustCompile(`(?smU)<ul class="nav nav-list primary left-menu">.*</ul>`),
	}
	return indexPage.AnalyzeLinks()
}

// Content is html content data
type Content struct {
	url  string
	data []byte
}

// concurrency and limit get content
func concurContent(links []string, semaphore int) (contMap map[int]*Content) {
	contMap = make(map[int]*Content)

	// get file content
	wg := sync.WaitGroup{}
	sema := make(chan struct{}, semaphore)
	for i, link := range links {
		wg.Add(1)
		go func(n int, link string) {
			// currency limit
			sema <- struct{}{}
			defer func() {
				wg.Done()
				<-sema
			}()
			// get content
			cont, err := content(link)
			if err != nil {
				log.Println(err)
				return
			}
			// add to content map
			contMap[n] = &Content{
				link,
				cont,
			}
			log.Println("get content", link)
		}(i, link)
	}
	wg.Wait()

	return contMap
}

func main() {
	// index page
	clinks, err := launch(`https://www.tutorialspoint.com/go/index.htm`)
	if err != nil {
		log.Fatalln(err)
	}

	// concurrency get content part
	contMap := concurContent(clinks, 5)

	// write to
	htmlFile := helper.MustOpenFile("./storage/app/download.html")
	defer htmlFile.Close()
	for i := 0; i < len(contMap); i++ {
		//log.Println(i, bytes.SplitN((*contMap[i]).data, []byte(""), 100), "...")
		c := contMap[i]
		htmlFile.WriteString(fmt.Sprintf("<h3>%s</h3>", c.url))
		htmlFile.WriteString(string(c.data))
	}

	fmt.Println("Over!")
}
