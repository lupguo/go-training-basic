package body

import (
	"bytes"
	"fmt"
	"go-example/helper"
	"golang.org/x/net/html"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"
)

func FindLeftMock(url string, re *regexp.Regexp) (links []string, err error) {
	htmlTxt, err := ioutil.ReadFile(`/data/github.com/tkstorm/go-example/testdata/temp/tutorialspoint_go_index.html`)
	if err != nil {
		return nil, err
	}
	re = regexp.MustCompile(`(?smU)<ul class="nav nav-list primary left-menu">.*</ul>`)
	match := re.Find(htmlTxt)

	// parse match find
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				//link, err := resp.Request.URL.Parse(a.Val)
				//if err != nil {
				//	continue // ignore bad URLs
				//}
				links = append(links, a.Val)
			}
		}
	}

	// parse matches content urls
	doc, err := html.Parse(bytes.NewReader(match))
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	forEachNode(doc, visitNode, nil)

	return links, nil
}

// find index page left part
func FindLeft(url string, re *regexp.Regexp) (links []string, err error) {
	if client == nil {
		client = helper.NewHttpClient()
	}

	resp, err := client.Request("GET", url, 1*time.Second)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}

	// regex find
	htmlTxt, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read html body: %s", err)
	}
	matches := re.FindAll(htmlTxt, -1)

	// parse match find
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue // ignore bad URLs
				}
				links = append(links, link.String())
			}
		}
	}

	// parse matches content urls
	for _, b := range matches {
		doc, err := html.Parse(bytes.NewReader(b))
		if err != nil {
			return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
		}
		forEachNode(doc, visitNode, nil)
	}

	return links, nil
}

// Copied from gopl.io/ch5/outline2.
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}
