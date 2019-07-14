package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	type KV map[string]interface{}

	// make map
	ume := KV{
		"uname":  "user01",
		"age":    37,
		"height": 182.20,
		"likes": []string{
			"cat",
			"dog",
			"pig",
		},
		"schools": KV{
			"primary":    "HanShan School",
			"middle":     "TangTian School",
			"high":       "JiuZhong",
			"university": "JiangXiShiFan",
		},
	}
	
	// json encode map
	b, _ := json.MarshalIndent(ume, "", "\t")
	fmt.Printf("%s\n", b)
}
