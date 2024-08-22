package config

import (
	"fmt"
	"github.com/spf13/viper"
)

var Gconfig *viper.Viper

func InitConfig() {
	Gconfig = viper.New()
	Gconfig.SetConfigName("config")   // 配置文件名称(无扩展名)
	Gconfig.SetConfigType("yaml")     // 如果配置文件的名称中没有扩展名，则需要配置此项
	Gconfig.AddConfigPath("./config") // 查找配置文件所在的路径
	Gconfig.AddConfigPath("$HOME/")   // 多次调用以添加多个搜索路径
	Gconfig.AddConfigPath(".")        // 还可以在工作目录中查找配置
	err := Gconfig.ReadInConfig()     // 查找并读取配置文件
	if err != nil {                   // 处理读取配置文件的错误
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}
