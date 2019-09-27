package main

import (
	"fmt"
	"go-example/case/crawler/crawler_v2/links"
)

func main() {
	//link := "https://golang.org"
	//link := "https://tkstorm.com"
	link := "https://www.taobao.com"
	for _, u := range links.Extract(link) {
		fmt.Println(u)
	}
}
