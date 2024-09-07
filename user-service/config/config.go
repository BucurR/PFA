package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type DbConfig struct {
	Server   string
	Port     int
	Database string
	Passwrod string
	User     string
}

func NewConfig() *DbConfig {
	return &DbConfig{}
}
func (c *DbConfig) InitConfig() {
	godotenv.Load("user-service/.env")
	c.Server = readEnv("AZURE_DB_SERVER", "localhost")
	port, err := strconv.Atoi(readEnv("AZURE_DB_PORT", "8080"))
	if err != nil {
		port = 1433
	}
	c.Port = port
	c.Database = readEnv("AZURE_DB_NAME", "my_db")
	c.User = readEnv("AZURE_DB_USERNAME", "")
	c.Passwrod = readEnv("AZURE_DB_PASSWORD", "")

}
func readEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
