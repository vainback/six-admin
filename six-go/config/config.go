package config

import (
	"fmt"
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Notices = make(chan bool, 128)

var _config map[string]any

func init() {
	v := viper.New()
	v.AddConfigPath(".")
	v.SetConfigType("yaml")
	v.SetConfigName("config")
	if err := v.ReadInConfig(); err != nil {
		log.Fatalln(err)
	}
	v.WatchConfig()
	if err := v.Unmarshal(&_config); err != nil {
		log.Fatalln(err)
	}
	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println(in.Name, " 配置文件内容被更改。如有必要，请重启程序！")
		if err := v.Unmarshal(&_config); err != nil {
			log.Println(err)
		}
		Notices <- true
	})
}

func Get(key string) any {
	val, ok := _config[key]
	if !ok {
		return nil
	}
	return val
}
