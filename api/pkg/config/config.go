// pkg/config/config.go
package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// ServerConfig 结构体保持不变
type ServerConfig struct {
	Port string `mapstructure:"port"` // 服务器端口
}

// DSNConfig 结构体用于存储单个环境的DSN
type DSNConfig struct {
	DSN string `mapstructure:"dsn"` // 数据库连接字符串
}

// DatabaseConfig 结构体被修改以支持多环境
type DatabaseConfig struct {
	Active string    `mapstructure:"active"` // 当前激活的环境
	Local  DSNConfig `mapstructure:"local"`  // 本地环境配置
	Azure  DSNConfig `mapstructure:"azure"`  // Azure环境配置
}

// Config 结构体现在引用了修改后的DatabaseConfig
type Config struct {
	Server   ServerConfig   `mapstructure:"server"`
	Database DatabaseConfig `mapstructure:"database"`
}

// GetDSN 是一个辅助方法，用于根据激活的环境返回正确的DSN字符串
func (c *Config) GetDSN() (string, error) {
	switch c.Database.Active {
	case "local":
		return c.Database.Local.DSN, nil
	case "azure":
		return c.Database.Azure.DSN, nil
	default:
		return "", fmt.Errorf("未知的激活数据库环境: '%s'", c.Database.Active)
	}
}

// LoadConfig 从 configs/config.yaml 加载配置
// 这个函数的逻辑基本不变，因为它只是将文件内容映射到结构体
func LoadConfig() (config Config, err error) {
	// 设置配置文件的路径和名称
	// 注意: 我将路径改回了 "./configs"，请根据您的实际项目结构调整
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
