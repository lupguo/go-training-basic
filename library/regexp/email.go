package main

import (
	"fmt"
	"go-example/helper"
	"os"
	"regexp"
	"text/tabwriter"
)

func main() {
	re := regexp.MustCompile(`^(?P<user>[0-9a-z.\-_]+)@[a-z0-9]+(?P<suffix>\.[a-z]{2,3})$`)
	fmt.Println(re.LiteralPrefix())

	emails := []string{
		`terry@163.com`,
		`qq15231223@qq.com`,
		`cod@google.com`,
		`hello@sina.cn`,
		`a-b-c@sina.com`,
		`a_b_c@sina.com`,
		`a.b.c@sina.com`,
		`12350@126.cn`,
		`dd@ss.ssd`,
		`dd@ss.ssdd`,
		`1234`,
		`@sina.com`,
	}

	tw := tabwriter.NewWriter(os.Stdout, 4, 4, 2, ' ', 0)
	defer tw.Flush()
	show := helper.TwShow(tw)
	show(
		"Email",
		"Match Result",
		"SubexpNames",
		"Name",
		"reversed",
		"FindString",
		"FindStringSubmatch",
		"NumSubexp",
	)

	reversed := fmt.Sprintf("${%s} ${%s}", re.SubexpNames()[2], re.SubexpNames()[1])
	for _, e := range emails {

		show(
			e,
			re.MatchString(e),
			re.SubexpNames(),
			fmt.Sprintf("${%s},${%s}", re.SubexpNames()[1], re.SubexpNames()[2]),
			re.ReplaceAllString(e, reversed),
			re.FindString(e),
			re.FindStringSubmatch(e),
			re.NumSubexp(),
		)
	}
}
