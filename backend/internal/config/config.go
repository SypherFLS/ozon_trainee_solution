package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Env     string     `yaml:"env"`
	HTTP    HTTPConfig `yaml:"http"`
	Storage string     `yaml:"storage"`
	DB      DBConfig   `yaml:"db"`
}

type HTTPConfig struct {
	Host        string        `yaml:"host"`
	Port        int           `yaml:"port"`
	Timeout     time.Duration `yaml:"timeout"`
	IdleTimeout time.Duration `yaml:"idle_timeout"`
}

type DBConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
	SSLMode  string `yaml:"sslmode"`
}

func resolveConfigPath() (string, error) {
	if configPath := os.Getenv("CONFIG_PATH"); configPath != "" {
		if _, err := os.Stat(configPath); err != nil {
			return "", fmt.Errorf("config file does not exist: %s", configPath)
		}
		return configPath, nil
	}

	candidates := []string{
		"config.yaml",
		"config.yml",
		filepath.Join("config", "config.yaml"),
		filepath.Join("config", "config.yml"),
		filepath.Join("deploy", "config", "config.yml"),
	}

	for _, candidate := range candidates {
		if _, err := os.Stat(candidate); err == nil {
			if absCandidate, absErr := filepath.Abs(candidate); absErr == nil {
				return absCandidate, nil
			}
			return candidate, nil
		}
	}

	return "", fmt.Errorf("config file not found; set CONFIG_PATH explicitly")
}

func MustLoad() *Config {
	configPath, err := resolveConfigPath()
	if err != nil {
		log.Fatal(err)
	}

	viper.SetConfigFile(configPath)
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("error reading config file: %s", err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("error unmarshaling config: %s", err)
	}

	return &cfg
}

