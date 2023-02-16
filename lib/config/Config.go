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
	// Logger 配置日志文件
	Logger struct {
		Level      string `json:"level"`
		Filename   string `json:"filename"`
		MaxSize    int    `json:"maxsize"`
		MaxBackups int    `json:"maxbackups"`
		MaxAge     int    `json:"maxage"`
	}

	// Server 配置网关程序
	Server struct {
		Mode    string `json:"mode"`
		Address string `json:"address"`
		Port    string `json:"port"`
	}

	// Config 配置文件
	Config struct {
		Logger `json:"logger"`
		Server `json:"server"`
	}
)

func NewConfig() *Config {
	defaultConfig := &Config{
		Logger: Logger{
			Level:      "info",
			Filename:   "smart.log",
			MaxSize:    3,
			MaxBackups: 3,
			MaxAge:     3,
		},
		Server: Server{
			Mode:    "release",
			Address: "127.0.0.1",
			Port:    ":8080",
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
