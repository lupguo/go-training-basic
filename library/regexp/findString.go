// http(s)://www.google.com/ab/cd.html
package main

import (
	"go-example/helper"
	"os"
	"regexp"
	"text/tabwriter"
)

func main() {
	urls := []string{
		"http://google.com",
		"https://www.google.com",
		"https://tkstorm.com/dir1/goods",
		"https://www.qq.com/dir1/user",
		"ab.com",
		"abcd",
		"aa.bb.cc.dd",
	}

	tw := tabwriter.NewWriter(os.Stdout, 8, 0, 4, ' ', 0)
	defer tw.Flush()
	show := helper.TwShow(tw)
	show(
		"URL",
		"Match",
		"FindString",
		"FindAllString",
		"FindStringSubmatch",
		"FindAllStringSubmatch",
		"FindStringIndex",
		"FindStringSubmatchIndex",
	)

	re := regexp.MustCompile(`^(http(?:s)?://)?([^/]*\.)+([^/]*).*$`)
	for _, url := range urls {
		show(
			url,
			re.MatchString(url),
			re.FindString(url),
			re.FindAllString(url, -1),
			re.FindStringSubmatch(url),
			re.FindAllStringSubmatch(url,-1),
			re.FindStringIndex(url),
			re.FindStringSubmatchIndex(url),
		)
	}

}
