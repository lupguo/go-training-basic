package main

import (
	"os"
	"text/template"
)

type Inventory struct {
	Material string
	Count    uint
}

func main() {
	// data
	sweaters := Inventory{"wool", 17}

	// new test template
	tpl := template.New("test")
	// template text parse
	_, err := tpl.Parse("{{.Count}} items are made of {{.Material}}\n")
	if err != nil {
		panic(err)
	}
	// template text execute, data apply
	err = tpl.Execute(os.Stdout, sweaters)
	if err != nil {
		panic(err)
	}

	// space
	_, err = tpl.Parse("{{23 -}} < {{- 45}}")
	if err := tpl.Execute(os.Stdout, nil); err != nil {
		panic(err)
	}

}
