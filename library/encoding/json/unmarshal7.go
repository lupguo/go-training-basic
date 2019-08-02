package main

import (
	"encoding/json"
	"log"
)

func main() {
	// Ai响应的商品基本信息
	type Goods struct {
		GoodsTitle      string `json:"goodstitle"`
		UrlTitle        string `json:"url_title"`
		WareCode        string `json:"warecode"`
		ShopPrice       string `json:"shopprice"`
		DisplayPrice    string `json:"displayprice"`
		Discount        string `json:"discount"`
		ImgUrl          string `json:"imgurl"`
		GridUrl         string `json:"gridurl"`
		GoodsNum        string `json:"goodsnum"`
		AppDisplayPrice string `json:"appdisplayprice"`
	}
	type GoodsList []Goods

	// BackJson
	type BackJson struct {
		Result   GoodsList         `json:"result"`
		Msg      string            `json:"msg"`
		Status   string            `json:"status"`
		AggsData map[string]string `json:"aggsData"`
	}

	bj := new(BackJson)
	d := []byte(`{"msg":"params is not a regular or valid json","result":[],"status":"0"}`)
	if err := json.Unmarshal(d, bj); err != nil {
		log.Fatal(err)
	}
	log.Printf("%#v", bj)
}
