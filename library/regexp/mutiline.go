package main

import (
	"go-example/helper"
	"io/ioutil"
	"log"
	"regexp"
)

func main() {
	// open file get content
	data, err := ioutil.ReadFile(helper.TestDataPath + `/html/download.html`)
	if err != nil {
		log.Fatalln(data)
	}

	var res string
	// regex replace data
	fitRe1 := regexp.MustCompile(`(?ms)<div class="content">\s+<div class="col-md-7 middle-col">.*<div class="clearer"></div>\s+<hr/>\s+<h1>`)
	log.Println(fitRe1.MatchString(res), len(fitRe1.FindAllString(res, -1)))

	// replace
	res = fitRe1.ReplaceAllString(string(data), `
<div class="content">
<div class="col-md-7 middle-col">
<hr/>
<h1>`)

	// filter2 : filter <script>
	fitRe2 := regexp.MustCompile(`(?ms)<script>(<!--)?.*(//-->)?</script>`)
	//log.Println(fitRe2.MatchString(res))
	res = fitRe2.ReplaceAllString(res, `@@@@@`)

	// filter3 <hr/> pre-btn &  nxt-btn <hr/>
	fitRe3 := regexp.MustCompile(`(?ms)(<hr/>)?.<div class="pre-btn">.*print-btn center.*nxt-btn.*<hr/>`)
	//log.Println(fitRe3.MatchString(res), len(fitRe3.FindAllString(res, -1)))
	res = fitRe3.ReplaceAllString(res, `=====`)

	// filter4 <img>
	fitRe4 := regexp.MustCompile(`<img.*>`)
	log.Println(fitRe4.MatchString(res), len(fitRe4.FindAllString(res, -1)))
	res = fitRe4.ReplaceAllString(res, `++++`)

	// write to file
	//fmt.Println(res)
	if err := ioutil.WriteFile(helper.AppStorage+`/download_new.html`, []byte(res), 0644); err != nil {
		log.Fatalln(err)
	}
}
