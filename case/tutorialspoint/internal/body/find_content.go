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
	buf := bytes.NewBuffer([]byte{})
	if err := html.Render(buf, contentNode); err != nil {
		fmt.Println(err)
	}

	// filter
	//content = buf.Bytes()
	content = []byte(filter(url, buf.String()))

	log.Println("len size: ", buf.Len())
	return
}

// filter is html content page filter, return page data
func filter(url, res string) (left string) {
	// filter1: content
	fitRe1 := regexp.MustCompile(`(?msU)<div class="content">\s+<div class="col-md-7 middle-col">\s+(?P<title><h1>.*</h1>).*<div class="clearer"></div>\s+<hr[^>]*>`)
	helper.RegexDebug("content regex:", fitRe1, res, false)
	res = fitRe1.ReplaceAllStringFunc(res, func(s string) string {
		var result []byte
		template := `
<div class="content">
<div class="col-md-7 middle-col">
${title}
<h3> <a href="` + url + `" target="_blank">` + url + `</a></h3>
<hr/>
`
		for _, submatches := range fitRe1.FindAllStringSubmatchIndex(s, -1) {
			result = fitRe1.ExpandString(result, template, s, submatches)
		}
		return string(result)
	})

	// filter2 : filter <script>
	fitRe2 := regexp.MustCompile(`(?msU)<script[^>]*>.*</script>`)
	helper.RegexDebug("script regex:", fitRe2, res, false)
	res = fitRe2.ReplaceAllString(res, ``)

	// filter3 <hr/> pre-btn &  nxt-btn <hr/>
	fitRe3 := regexp.MustCompile(`(?msU)<hr[^>]*>\s+<div class="pre-btn">.*<div class="nxt-btn">.*<hr[^>]*>`)
	helper.RegexDebug("pre-btn regex:", fitRe3, res, false)
	res = fitRe3.ReplaceAllString(res, ``)

	// filter4 <img>
	fitRe4 := regexp.MustCompile(`<img[^>]*>`)
	helper.RegexDebug("image regex:", fitRe4, res, false)
	res = fitRe4.ReplaceAllString(res, ``)

	// filter5 Advertisements
	fitRe5 := regexp.MustCompile(`<div[^>]*>Advertisements</div>`)
	helper.RegexDebug("advertisement regex:", fitRe5, res, false)
	res = fitRe5.ReplaceAllString(res, ``)

	return res
}
