package config

import (
    "github.com/spf13/viper"
    "log"
)

type Config struct {
    Port string
}

func LoadConfig() *Config {
    viper.SetDefault("PORT", "8080")
    viper.AutomaticEnv()

    config := &Config{
        Port: viper.GetString("PORT"),
    }

    log.Printf("Starting server on port %s", config.Port)
    return config
}

