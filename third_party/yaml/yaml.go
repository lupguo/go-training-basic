package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type Config struct {
	Http struct {
		EsPrice     string `yaml:"es_price"`
		AiRecommend string `yaml:"ai_recommend"`
	}
}

var cfg *Config

// 获取配置
func GetConfig() *Config {
	return cfg
}

// 解析Config配置
func main() {
	if cfg == nil {
		cfg = new(Config)
	}
	// open yaml config
	f, err := os.Open("./testdata/config/app.yaml")
	if err != nil {
		panic(err)
	}
	// read all
	d, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	if err := yaml.Unmarshal(d, cfg); err != nil {
		panic(err)
	}

	fmt.Printf("%+v", cfg)
}