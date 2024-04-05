package config

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

var CONF *Config = nil

// Config 配置
type Config struct {
	Proxy  Proxy  `toml:"proxy"`
	Logger Logger `toml:"logger"`
}

// Proxy 代理
type Proxy struct {
	ProxyPort        int32  `toml:"proxy_port"`
	RedirectSdk      string `toml:"redirect_sdk"`
	RedirectDispatch string `toml:"redirect_dispatch"`
}

// Logger 日志
type Logger struct {
	Level        string `toml:"level"`
	TrackLine    bool   `toml:"track_line"`
	TrackThread  bool   `toml:"track_thread"`
	EnableFile   bool   `toml:"enable_file"`
	FileMaxSize  int32  `toml:"file_max_size"`
	DisableColor bool   `toml:"disable_color"`
	EnableJson   bool   `toml:"enable_json"`
}

func InitConfig(filePath string) {
	CONF = new(Config)
	CONF.loadConfigFile(filePath)
}

func GetConfig() *Config {
	return CONF
}

// 加载配置文件
func (c *Config) loadConfigFile(filePath string) {
	_, err := toml.DecodeFile(filePath, &c)
	if err != nil {
		info := fmt.Sprintf("config file load error: %v\n", err)
		panic(info)
	}
}
