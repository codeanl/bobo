package utils

import (
	"bobo-server/config"
	"github.com/spf13/viper"
	"log"
)

func InitViper() {
	configPath := "config/config.toml"
	v := viper.New()
	v.SetConfigFile(configPath)
	v.AutomaticEnv() // 允许使用环境变量
	// 读取配置文件
	if err := v.ReadInConfig(); err != nil {
		log.Panic("配置文件读取失败: ", err)
	}
	// 加载配置文件内容到结构体对象
	if err := v.Unmarshal(&config.Cfg); err != nil {
		log.Panic("配置文件内容加载失败: ", err)
	}
	log.Println("配置文件内容加载成功")
}
