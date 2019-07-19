// https://golang.org/pkg/regexp/
// Find(All)?(String)?(Submatch)?(Index)?

package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile(`foo.?`)

	// []byte
	fmt.Println(`=> Find Byte<=`)
	fmt.Printf("%q\n", re.Find([]byte(`seafood fool`)))
	fmt.Printf("%q\n", re.FindAll([]byte(`seafood fool`),-1))
	fmt.Printf("%q\n", re.FindSubmatch([]byte(`seafood fool`)))
	fmt.Printf("%q\n", re.FindAllSubmatch([]byte(`seafood fool`),-1))

	// Index
	fmt.Println(`=> Find Index <=`)
	fmt.Printf("%#v\n", re.FindIndex([]byte(`seafood fool`)))
	fmt.Printf("%#v\n", re.FindAllIndex([]byte(`seafood fool`),-1))
	fmt.Printf("%#v\n", re.FindSubmatchIndex([]byte(`seafood fool`)))
	fmt.Printf("%#v\n", re.FindAllSubmatchIndex([]byte(`seafood fool`),-1))
	fmt.Printf("%#v\n", re.FindStringIndex(`seafood fool`))
	fmt.Printf("%#v\n", re.FindAllStringIndex(`seafood fool`,-1))

	// string
	fmt.Println(`=> Find String<=`)
	fmt.Printf("%q\n", re.FindString(`seafood fool`))
	fmt.Printf("%q\n", re.FindAllString(`seafood fool`,-1))
	fmt.Printf("%q\n", re.FindAllString(`seafood fool`,1))
	fmt.Printf("%q\n", re.FindStringSubmatch(`seafood fool`))
}
