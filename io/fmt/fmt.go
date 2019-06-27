package main

import "fmt"

type Person struct {
	id   int
	name string
	pass string
	f    bool
}

type ClassPerson []Person

func main() {
	p1 := Person{100, "∂terry", "123", false}
	p2 := Person{200, "åclark", "234", true}

	cp := make(ClassPerson, 10)
	cp = append(cp, p1, p2)

	fmt.Printf("%v\n", p1)
	fmt.Printf("%+v\n", p2)

	// slice
	fmt.Printf("%v\n", cp)
	fmt.Printf("%+v\n", cp)
	fmt.Printf("%#v\n", cp)
	fmt.Printf("%T\n", cp)
	fmt.Printf("%%\n")

	// bool
	fmt.Printf("%t", cp[0].f)

	// int
	// %c
	//	%d
	//	%o
	//	%q
	//	%x
	//	%X
	//	%U"
	fmt.Printf("\n"+
		"%%c=%c, "+
		"%%b=%b, "+
		"%%d=%d, "+
		"%%o=%o, "+
		"%%q=%q, "+
		"%%x=%x, "+
		"%%X=%X,"+
		"%%U=%U",
		'å', 97, 'a', 'b', 'c', 'd', 'e', 'f',
	)

	fmt.Printf("\n%g,%E", 1e100, 1e100)

	fmt.Printf("\n%p", cp)
	fmt.Printf("\n%#p", &cp[0])

	fmt.Printf("\n%+f", -10.89)
	fmt.Printf("\n%#x", "abc")
	fmt.Printf("\n%5.2f", 0.18234)
	fmt.Printf("\n%010s", "bac")

	fmt.Printf("\n%#x", "bac")

	fmt.Printf("\n%U", 'x')
	fmt.Printf("\n%#U", 'x')

	fmt.Printf("\n%[3]*.[2]*[1]f", 1.0, 2, 30)

	fmt.Printf("\n%d %d %#[1]v %#v", 16, 17)
}
