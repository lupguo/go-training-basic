package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"time"
)

func main()  {
	// alternatively, you can create a new viper instance.
	var runtime_viper = viper.New()
	_ = runtime_viper.AddRemoteProvider("etcd", "http://127.0.0.1:4001", "/config/hugo.yml")
	runtime_viper.SetConfigType("yaml")
	if err := runtime_viper.ReadRemoteConfig(); err != nil {
		panic(err)
	}

	// unmarshal config
	runtime_conf := ""
	_ = runtime_viper.Unmarshal(&runtime_conf)

	// open a goroutine to watch remote changes forever
	go func(){
		for {
			time.Sleep(time.Second * 5) // delay after each request

			// currently, only tested with etcd support
			err := runtime_viper.WatchRemoteConfig()
			if err != nil {
				logrus.Errorf("unable to read remote config: %v", err)
				continue
			}

			// unmarshal new config into our runtime config struct. you can also use channel
			// to implement a signal to notify the system of the changes
			runtime_viper.Unmarshal(&runtime_conf)
		}
	}()
}
