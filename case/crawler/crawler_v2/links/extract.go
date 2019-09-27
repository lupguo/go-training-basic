package links

import (
	"crypto/tls"
	"golang.org/x/net/html"
	"golang.org/x/net/publicsuffix"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"sort"
	"strings"
	"time"
)

type linksMap map[string]bool

//FilterAttr using function filter the href which the function test for true
var FilterAttr func(attr html.Attribute) bool

//Extract link page document to links
func Extract(link string) (links []string) {
	resp, err := getResponse(link)
	if err != nil {
		log.Println("request link error:", err)
		return
	}
	defer resp.Body.Close()
	// parse
	root, err := html.Parse(resp.Body)
	if err != nil {
		log.Println("parse error:", err)
	}
	// link map
	mapLink := make(linksMap)
	forEachNode(root, mapLink)
	return perfectLinks(link, mapLink)
}

// get request response by user defined http client
func getResponse(link string) (resp *http.Response, err error) {
	urlLink, err := url.Parse(link)
	if err != nil {
		return
	}
	jar, _ := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	jar.SetCookies(urlLink, []*http.Cookie{
		{
			Name:  "Hey",
			Value: "Man",
			Path:  "/",
		},
	})
	httpClient := &http.Client{
		Transport: &http.Transport{
			Proxy:       nil,
			DialContext: nil,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
		CheckRedirect: nil,
		Jar:           nil,
		Timeout:       10 * time.Second,
	}

	req, _ := http.NewRequest("GET", link, nil)
	return httpClient.Do(req)
}

// perfect links: /pkg => https://google.org/pkg
func perfectLinks(link string, mapLink linksMap) (links []string) {
	if len(mapLink) == 0 {
		return
	}
	// origin url
	originUrl, err := url.ParseRequestURI(link)
	if err != nil {
		log.Fatal("parse link error:", originUrl)
	}
	// link slice
	for l, _ := range mapLink {
		newUrl, _ := originUrl.Parse(l)
		links = append(links, newUrl.String())
	}
	sort.Strings(links)
	return
}

// visitNode visit <a href=""/> html node, and return match link
func visitNode(n *html.Node) (link string) {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, attr := range n.Attr {
			if attr.Key != "href" ||
				strings.HasPrefix(attr.Val, "mailto:") ||
				strings.HasPrefix(attr.Val, "javascript:") {
				continue
			}
			// user defined func filter rule
			if FilterAttr != nil && FilterAttr(attr) {
				continue
			}
			return attr.Val
		}
	}
	return
}

// visit html node by bfs
func forEachNode(n *html.Node, ml linksMap) {
	// pre visit node
	if link := visitNode(n); link != "" {
		ml[link] = true
	}
	// visit child node (bfs)
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, ml)
	}
}
