package main

import (
	"bytes"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

func main() {
	// default
	viper.SetDefault("ContentDir", "content")
	viper.SetDefault("LayoutDir", "layouts")
	viper.SetDefault("Taxonomies", map[string]string{"tag": "tags", "category": "categories"})

	// read from config file
	viper.SetConfigName("config")             // name of config file (without extension)
	viper.AddConfigPath("/etc/appname/")      // path to look for the config file in
	viper.AddConfigPath("$HOME/.appname")     // call multiple times to add many search paths
	viper.AddConfigPath("./testdata/config/") // call multiple times to add many search paths
	viper.AddConfigPath(".")                  // optionally look for config in the working directory
	err := viper.ReadInConfig()               // Find and read the config file
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	// 写入成功，按配置文件最新寻找到的配置文件，覆盖写
	err = viper.WriteConfig() // writes current config to predefined path set by 'viper.AddConfigPath()' and 'viper.SetConfigName'
	if err != nil {
		fmt.Printf("1. write config file err: %s\n", err)
	}

	// 写入失败，文件存在，不会覆盖已有配置写入
	err = viper.SafeWriteConfig()
	if err != nil {
		fmt.Printf("2. write config file err: %s\n", err)
	}

	// 写入成功，文件不存在，创建文件，写入
	err = viper.WriteConfigAs("./testdata/temp/config.yml")
	if err != nil {
		fmt.Printf("3. write config file err: %s\n", err)
	}

	// 写入失败，文件存在，不会覆盖已有配置写入
	err = viper.SafeWriteConfigAs("./testdata/temp/config.yml") // will error since it has already been written
	if err != nil {
		fmt.Printf("4. write config file err: %s\n", err)
	}

	err = viper.SafeWriteConfigAs("./testdata/temp/other_config.yml")
	if err != nil {
		fmt.Printf("5. write config file err: %s\n", err)
	}

	// 配置监听
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})

	// 基于IO读取配置信息
	viper.SetConfigType("yaml") // or viper.SetConfigType("YAML")
	// any approach to require this configuration into your program.
	var yamlExample = []byte(`
Hacker: true
name: steve
hobbies:
- skateboarding
- snowboarding
- go
clothing:
  jacket: leather
  trousers: denim
age: 35
eyes : brown
beard: true
company:
  inner:
    - alibaba
    - tencent
  outer:
    - google
    - aws
    - netflix
`)
	err = viper.ReadConfig(bytes.NewBuffer(yamlExample))
	if err != nil {
		fmt.Printf("read config err: %s\n", err)
	}
	fmt.Printf("%+v\n", viper.Get("name"))
	fmt.Printf("%+v\n", viper.GetBool("beard"))
	fmt.Printf("%+v\n", viper.GetInt("age"))
	fmt.Printf("%+v\n", viper.GetStringSlice("hobbies"))
	fmt.Printf("%+v\n", viper.GetStringMap("clothing"))
	fmt.Printf("%+v\n", viper.GetStringMapStringSlice("company"))
	fmt.Printf("outer company => %+v\n", viper.GetStringSlice("company.outer"))
	fmt.Printf("clothing.trousers=>%s\n", viper.Get("clothing.trousers"))

	// 值覆盖
	viper.Set("name", "clark")
	viper.Set("age", 31)
	fmt.Printf("name:%+v, age:%d\n", viper.Get("name"), viper.GetInt("age"))

	// 环境变量读取
	os.Setenv("SPF_ID", "13") // typically done outside of the app
	os.Setenv("THIS_YEAR", "2019")
	viper.SetEnvPrefix("spf") // will be uppercased automatically
	viper.BindEnv("id")
	viper.BindEnv("year", "THIS_YEAR")
	fmt.Printf("id=%s, this year is %s\n", viper.Get("id"), viper.Get("year"))

	// Flag变量读取
	pflag.Int("port", 1234, "Port to run Application server on")
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)
	fmt.Printf("Flag port: %d\n", viper.Get("port"))

	// Marshalling to string
	c := viper.AllSettings()
	bs, err := yaml.Marshal(c)
	if err != nil {
		log.Fatalf("unable to marshal config to YAML: %v", err)
	}
	fmt.Printf("----\n%s\n", bs)

	<-make(chan struct{})
}

