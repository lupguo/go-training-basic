package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// 推荐场景参数
	type SceneParams map[string]interface{}

	// 推荐请求参数
	type Params struct {
		SceneParams
		Region   string `json:"regioncode"`
		SiteCode string `json:"site_code"`
		PipeCode string `json:"pipelinecode"`
		Platform string `json:"platform"`
		Cookie   string `json:"cookie"`
		Lang     string `json:"lang"`
	}

	p := new(Params)
	sp := make(SceneParams)
	sp["name"] = "Username"
	sp["categoryid"] = 37
	sp["storeid"] = "1024890"
	p.SceneParams = sp
	fmt.Printf("%+v\n", p)

	b, _ := json.MarshalIndent(p, "", "\t")
	fmt.Printf("%s\n", b)

}
