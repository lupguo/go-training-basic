package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

type Data []byte

// Page Resource
type PageA interface {
	Analyze() []interface{}
}

// FirstPage(left, regex)
type Page struct {
	url  string
	regx *regexp.Regexp
}

// Analyze IndexPage using regex got match content, write to chan
func (p *Page) Analyze() chan interface{} {
	contCh := make(chan interface{})

	// crawler get page
	ct, err := crawlContent(p.url)
	if err != nil {
		log.Println(err)
		return nil
	}

	// regex match
	

	// write to chan

	return contCh
}

func matchData()  {
	
}

// crawlContent crawl url page and get html body content
func crawlContent(url string) (cont []byte, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	// get body content
	ct, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return ct, nil
}

// Content is html content data
type Content struct {
	chapter int
	data    Data
}

// ContentList is list of content
type ContentList []*Content

// Sort ContentList
func (cs *ContentList) Sort() {

}

// Write Content List Data to file
func (cs *ContentList) WriteToFile(filename string, content []byte) {

}

func main() {

}
