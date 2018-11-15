package config

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Init() {
	if gin.Mode() == gin.TestMode {
		viper.Set("mongodb.host", "127.0.0.1")
		viper.Set("mongodb.port", "27017")
		viper.Set("mongodb.name", "schoolnetwork")
		viper.Set("mongodb.username", "")
		viper.Set("mongodb.password", "")
		viper.Set("name", "schoolnetwork")
		viper.Set("port", "8080")
		viper.Set("jwt.secret", "bda25a151f65a52z1a551f5az5e52s56156re5eas65ac45uy63z3d5g4y1f5d5t5u4iu4sd5s5sz11")
		viper.Set("secret", "a15f6a165rg415r161d564649a61b13g1aw53ed3a5135r153s13a51z123d1f56ae4fraw65")
	} else {
		viper.SetConfigName("config")
		viper.SetConfigType("json")
		viper.AddConfigPath("./")
		err := viper.ReadInConfig() // 搜索路径，并读取配置数据
		if err != nil {
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}
	}
}

func Get(key string) interface{} {
	return viper.Get(key)
}

func Set(key string, value interface{}) {
	viper.Set(key, value)
}
