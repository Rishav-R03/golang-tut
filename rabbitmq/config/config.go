package config

import "github.com/spf13/viper"

type Config struct {
	Server   ServerConfig   `mapstructure:"server"`
	RabbitMQ RabbitConfig   `mapstructure:"rabbitmq"`
	Database DatabaseConfig `mapstructure:"database"`
}

type ServerConfig struct {
	PORT int `mapstructure:"port"`
}

type RabbitConfig struct {
	URL       string `mapstructure:"url"`
	QueueName string `mapstructure:"queue"`
}

type DatabaseConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("config") // name of configg file without extension
	viper.SetConfigType("yml")
	viper.AddConfigPath(".") // look for config in the working directory

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config

	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
