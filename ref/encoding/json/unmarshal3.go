package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

func main() {
	// open file
	type Goods struct {
		DisplayPrice string `json:"displayprice"`
		Gridurl      string `json:"gridurl"`
		Imgurl       string `json:"imgurl"`
		ShopPrice    string `json:"shopprice"`
	}
	type AiRs struct {
		Result []*Goods `json:"result"`
	}
	aiRs := new(AiRs)
	f, err := ioutil.ReadFile("testdata/json/resp.json")
	if err != nil {
		log.Fatal(err)
	}
	// print data
	//log.Printf("%s\n", f)

	if err := json.Unmarshal(f, &aiRs); err != nil {
		log.Fatalln(err)
	}

	// write
	aiRs.Result = aiRs.Result[:2]
	if b, err := json.MarshalIndent( aiRs,"", "\t"); err != nil {
		log.Fatal(err)
	} else {
		log.Printf("\n%s\n", b)
	}
}
