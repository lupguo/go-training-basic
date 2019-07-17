package url

import (
	"fmt"
	"net/url"
)

func main()  {
	u := &url.URL{
		Scheme:   "https",
		User:     url.UserPassword("me", "pass"),
		Host:     "example.com",
		Path:     "foo/bar",
		RawQuery: "x=1&y=2",
		Fragment: "anchor",
	}
	fmt.Println(u.String())
	fmt.Printf("%#v\n", *u)
	fmt.Println(u.Path)
	u.Opaque = "opaque"
	fmt.Println(u.String())
}