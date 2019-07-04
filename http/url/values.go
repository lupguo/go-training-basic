package main

import (
	"fmt"
	"log"
	"net/url"
)

func main() {
	v := url.Values{}
	v.Set("name", "Ava")
	v.Add("friend", "Jess")
	v.Add("friend", "Sarah")
	v.Add("friend", "Zoe")
	// v.Encode() == "name=Ava&friend=Jess&friend=Sarah&friend=Zoe"
	fmt.Println(v.Get("name"))
	fmt.Println(v.Get("friend"))
	fmt.Println(v["friend"])
	fmt.Println(v.Encode())

	// create new url
	u, err := url.Parse("https://google.com/search?apikey=key")
	if err != nil {
		log.Fatalln(err)
	}

	// concat query
	if us := u.RawQuery; us != "" {
		u.RawQuery = us + "&" + v.Encode()
	}else {
		u.RawQuery = v.Encode()
	}

	fmt.Println(u.String())
	fmt.Printf("%+v", u)
}