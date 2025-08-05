// pkg/config/config.go
package config

import "github.com/spf13/viper"

type ServerConfig struct {
	Port string `mapstructure:"port"` // 服务器端口
}

type DatabaseConfig struct {
	DSN string `mapstructure:"dsn"` // 数据库连接字符串
}

type Config struct {
	Server   ServerConfig   `mapstructure:"server"`
	Database DatabaseConfig `mapstructure:"database"`
}

// LoadConfig 从 configs/config.yaml 加载配置
func LoadConfig() (config Config, err error) {
	// 设置配置文件的路径和名称
	viper.AddConfigPath("./api/configs") // 路径相对于项目根目录
	viper.SetConfigName("config")        // 文件名是 config.yaml
	viper.SetConfigType("yaml")

	viper.AutomaticEnv() // 允许从环境变量读取

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
