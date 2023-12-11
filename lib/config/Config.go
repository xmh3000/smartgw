package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

const (
	configFile = "config/config.yml"
	configType = "yml"
)

type (
	// Config 配置文件
	Config struct {
		Gateway `json:"gateway"`
		Logger  `json:"logger"`
		Serial  `json:"serial"`
		Server  `json:"server"`
	}
	// Gateway 配置物联平台
	Gateway struct {
		Name     string `json:"name"`
		Ip       string `json:"ip"`
		Port     string `json:"port"`
		ClientID string `json:"clientid"`
		Username string `json:"username"`
		Password string `json:"password"`
	}

	// Logger 配置日志文件
	Logger struct {
		Level      string `json:"level"`
		Filename   string `json:"filename"`
		MaxSize    int    `json:"maxsize"`
		MaxBackups int    `json:"maxbackups"`
		MaxAge     int    `json:"maxage"`
	}

	// Serial 配置串口
	Serial struct {
		OpenEveryTime bool `json:"openeverytime"`
	}

	// Server 配置网关程序
	Server struct {
		Mode    string `json:"mode"`
		Address string `json:"address"`
		Port    string `json:"port"`
	}
)

func NewConfig() *Config {
	defaultConfig := &Config{
		Gateway: Gateway{
			Name:     "物联平台",
			Ip:       "thingsboard.io",
			Port:     "1883",
			ClientID: "SM001",
			Username: "admin",
			Password: "admin",
		},
		Logger: Logger{
			Level:      "info",
			Filename:   "smart.log",
			MaxSize:    3,
			MaxBackups: 3,
			MaxAge:     3,
		},
		Serial: Serial{
			OpenEveryTime: false,
		},
		Server: Server{
			Mode:    "release",
			Address: "127.0.0.1",
			Port:    ":80",
		},
	}
	conf := &Config{}

	viper.SetConfigType(configType)
	viper.SetConfigFile(configFile)
	if err := viper.ReadInConfig(); err != nil {
		return defaultConfig
	}

	viper.OnConfigChange(func(in fsnotify.Event) {
		viper.Unmarshal(conf)
		zap.S().Info("配置信息发生变动", conf)
	})
	viper.WatchConfig()

	if err := viper.Unmarshal(conf); err != nil {
		return defaultConfig
	}

	fmt.Println(conf)
	return conf
}

func SaveConfig(config *Config) {
	//fmt.Println(config)
	viper.SetConfigType(configType)
	viper.SetConfigFile(configFile)
	viper.Set("gateway.ip", config.Gateway.Ip)
	viper.Set("gateway.clientid", config.Gateway.ClientID)
	viper.Set("gateway.port", config.Gateway.Port)

	err := viper.WriteConfig()

	if err != nil {
		fmt.Println(err.Error())
	}
}
