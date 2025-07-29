package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// IConfig interface định nghĩa các phương thức cần thiết cho config
type IConfig interface {
	GetPort() int
	GetDBHost() string
	GetDBPort() int
	GetDBUser() string
	GetDBPassword() string
	GetDBName() string
	GetEnv() string
	GetLogLevel() string
	GetDB() *gorm.DB
}

// Config struct chứa tất cả cấu hình
type Config struct {
	Port     int    `json:"port"`
	DBHost   string `json:"db_host"`
	DBPort   int    `json:"db_port"`
	DBUser   string `json:"db_user"`
	DBPass   string `json:"db_pass"`
	DBName   string `json:"db_name"`
	Env      string `json:"env"`
	LogLevel string `json:"log_level"`
}

var instance *Config

// GetInstance trả về singleton instance của config
func GetInstance() IConfig {
	if instance == nil {
		instance = loadConfig()
	}
	return instance
}

// loadConfig load config từ environment variables
func loadConfig() *Config {
	// Load .env file nếu có
	godotenv.Load()

	port, _ := strconv.Atoi(getEnv("PORT", "8080"))
	dbPort, _ := strconv.Atoi(getEnv("DB_PORT", "3306"))

	return &Config{
		Port:     port,
		DBHost:   getEnv("DB_HOST", "localhost"),
		DBPort:   dbPort,
		DBUser:   getEnv("DB_USER", "root"),
		DBPass:   getEnv("DB_PASS", ""),
		DBName:   getEnv("DB_NAME", "notion"),
		Env:      getEnv("ENV", "development"),
		LogLevel: getEnv("LOG_LEVEL", "info"),
	}
}

// getEnv helper function để lấy environment variable với default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// Implement các phương thức của interface IConfig
func (c *Config) GetPort() int {
	return c.Port
}

func (c *Config) GetDBHost() string {
	return c.DBHost
}

func (c *Config) GetDBPort() int {
	return c.DBPort
}

func (c *Config) GetDBUser() string {
	return c.DBUser
}

func (c *Config) GetDBPassword() string {
	return c.DBPass
}

func (c *Config) GetDBName() string {
	return c.DBName
}

func (c *Config) GetEnv() string {
	return c.Env
}

func (c *Config) GetLogLevel() string {
	return c.LogLevel
}

func (c *Config) GetDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open(c.GetDBHost()), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
