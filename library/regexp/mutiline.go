package main

import (
	"go-example/helper"
	"io/ioutil"
	"log"
	"regexp"
)

func main() {
	url := `<a href="http://tkstorm.com" target="_blank">http://tkstorm.com</a>`

	// open file get content
	data, err := ioutil.ReadFile(helper.TestDataPath + `/html/download.html`)
	if err != nil {
		log.Fatalln(data)
	}
	res := string(data)

	// filter1: content
	fitRe1 := regexp.MustCompile(`(?msU)<div class="content">\s+<div class="col-md-7 middle-col">\s+(?P<title><h1>.*</h1>).*<div class="clearer"></div>\s+<hr/>`)
	helper.RegexDebug("content regex:", fitRe1, res, false)
	res = fitRe1.ReplaceAllStringFunc(res, func(s string) string {
		var result []byte
		template := `<div class="content"><div class="col-md-7 middle-col">${title} <h3>` + url + `</h3><hr/>`
		for _, submatches := range fitRe1.FindAllStringSubmatchIndex(s, -1) {
			result = fitRe1.ExpandString(result, template, s, submatches)
		}
		return string(result)
	})

	// filter2 : filter <script>
	fitRe2 := regexp.MustCompile(`(?msU)<script[^>]*>(<!--)?[^<]*(//-->)?</script>`)
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

	// write to file
	//fmt.Println(res)
	if err := ioutil.WriteFile(helper.TestDataPath+`/html/download_new.html`, []byte(res), 0644); err != nil {
		log.Fatalln(err)
	}
}
