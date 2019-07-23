package body

import (
	"bytes"
	"fmt"
	"go-example/helper"
	"golang.org/x/net/html"
	"log"
	"net/http"
	"regexp"
	"time"
)

var client *helper.HttpClient

func FindContent(url string, re *regexp.Regexp) (content []byte, err error) {
	if client == nil {
		client = helper.NewHttpClient()
	}

	fmt.Printf("%p\n", client)

	// open url, get its content
	resp, err := client.Request("GET", url, 1*time.Second)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}

	// parse match find
	contentNode := new(html.Node)
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "div" {
			for _, a := range n.Attr {
				if a.Key == "class" && a.Val == "content" {
					contentNode = n
					return
				}
				break
			}
		}
		return
	}

	// doc
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}
	forEachNode(doc, visitNode, nil)

	// save to buffer and return
	buf := bytes.NewBuffer(make([]byte, 4096))
	if err := html.Render(buf, contentNode); err != nil {
		fmt.Println(err)
	}

	log.Println("len size: ", buf.Len())
	return buf.Bytes(), err
}
