package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath("./")
	err := viper.ReadInConfig() // 搜索路径，并读取配置数据
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

func Get(key string) interface{} {
	return viper.Get(key)
}

func Set(key string, value interface{}) {
	viper.Set(key, value)
}
