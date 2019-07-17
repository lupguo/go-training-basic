package main

import (
	"io/ioutil"
	"os"
	"text/template"
)

func main() {
	// new template
	tmpl := template.New("test")
	tplBytes, err := ioutil.ReadFile("testdata/tpl/exp.tpl")
	if err != nil {
		panic(err)
	}
	// template parse
	if _, err = tmpl.Parse(string(tplBytes)); err != nil {
		panic(err)
	}
	// data apply
	if err := tmpl.Execute(os.Stdout, "Hey, man!"); err != nil {
		panic(err)
	}

	// execute template
	if err := tmpl.ExecuteTemplate(os.Stdout, "T2", "no data needed"); err != nil {
		panic(err)
	}
}
