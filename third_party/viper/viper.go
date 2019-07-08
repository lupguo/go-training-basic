package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func main()  {
	// default
	viper.SetDefault("ContentDir", "content")
	viper.SetDefault("LayoutDir", "layouts")
	viper.SetDefault("Taxonomies", map[string]string{"tag": "tags", "category": "categories"})

	// read from config file
	viper.SetConfigName("config") // name of config file (without extension)
	viper.AddConfigPath("/etc/appname/")   // path to look for the config file in
	viper.AddConfigPath("$HOME/.appname")  // call multiple times to add many search paths
	viper.AddConfigPath("./testdata/config/")  // call multiple times to add many search paths
	viper.AddConfigPath(".")               // optionally look for config in the working directory
	err := viper.ReadInConfig() // Find and read the config file
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
}
