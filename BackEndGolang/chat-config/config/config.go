package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v2"
)

type DatabaseConfig struct {
	Name     string `yaml:"name"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type ServerConfig struct {
	Port         int    `yaml:"port"`
	Host         string `yaml:"host"`
	IPAddress    string `yaml:"ip_address"`
	SecretKey    string `yaml:"secret_key"`
	ReadTimeout  int    `yaml:"read_timeout"`
	WriteTimeout int    `yaml:"write_timeout"`
	IdleTimeout  int    `yaml:"idle_timeout"`
}

type CacheConfig struct {
	Addr string `yaml:"addr"`
}

type ApiGatewayConfig struct {
	Addr string `yaml:"addr"`
}

type AppConfig struct {
	AppName    string           `yaml:"app_name"`
	Version    string           `yaml:"version"`
	Database   DatabaseConfig   `yaml:"database"`
	Server     ServerConfig     `yaml:"server"`
	Cache      CacheConfig      `yaml:"cache"`
	ApiGateway ApiGatewayConfig `yaml:"chat-backend"`
}

func LoadConfig(filePath string) (AppConfig, error) {
	// Load environment variables from .env file if present
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: No .env file found")
	}

	fmt.Printf("Open file %s...\n", filePath)

	// Open the YAML file
	yamlFile, err := os.ReadFile(filePath)

	fmt.Printf("Config %s...\n", yamlFile)

	if err != nil {
		return AppConfig{}, fmt.Errorf("error reading YAML file: %v", err)
	}

	// Replace environment variables in YAML content
	yamlContent := os.ExpandEnv(string(yamlFile))

	fmt.Printf("Content %s...\n", yamlContent)

	// Create a variable to hold the configuration
	var config AppConfig
	if err := yaml.Unmarshal([]byte(yamlContent), &config); err != nil {
		return AppConfig{}, fmt.Errorf("error parsing YAML file: %v", err)
	}

	// Return the parsed configuration
	return config, nil
}
