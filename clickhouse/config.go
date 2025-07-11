package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Database DatabaseConfig `mapstructure:"database"`
	Server   ServerConfig   `mapstructure:"server"`
}

type DatabaseConfig struct {
	ClickHouse ClickHouseConfig `mapstructure:"clickhouse"`
}

type ClickHouseConfig struct {
	Host             string `mapstructure:"host"`
	Port             int    `mapstructure:"port"`
	Database         string `mapstructure:"database"`
	User             string `mapstructure:"user"`
	Password         string `mapstructure:"password"`
	DialTimeout      string `mapstructure:"dial_timeout"`
	MaxExecutionTime string `mapstructure:"max_execution_time"`
}

type ServerConfig struct {
	Port string `mapstructure:"port"`
}

func LoadConfig(configPath string) (*Config, error) {
	viper.AddConfigPath(configPath)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return &config, nil
}

func (c *ClickHouseConfig) GetDSN() string {
	return fmt.Sprintf("clickhouse://%s:%s@%s:%d/%s?dial_timeout=%s&max_execution_time=%s",
		c.User, c.Password, c.Host, c.Port, c.Database, c.DialTimeout, c.MaxExecutionTime)
}