package application

import (
	"github.com/joho/godotenv"
	"os"
	"strconv"
	"strings"
)

type Config interface {
	GetString(string) string
	GetBool(string) bool
	GetInt(string) int
	GetFloat(string) float64
	GetStringSlice(string) []string
	GetStringMap(string) map[string]interface{}
}

type DefaultConfig struct{}

func (c *DefaultConfig) Init(app *Application) error {
	filename := ".env"
	if value, exists := os.LookupEnv("APP_ENV"); exists {
		filename = ".env." + value
		if _, err := os.Stat(filename); err == nil {
			if err := godotenv.Load(filename); err != nil {
				app.Logger().Info("No " + filename + " file found")
			}
		}
	}

	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		app.Logger().Info("No .env file found")
	}

	return nil
}

func NewDefaultConfig() Config {
	return &DefaultConfig{}
}

func (c *DefaultConfig) GetString(s string) string {
	if value, exists := os.LookupEnv(s); exists {
		return value
	}

	panic("The key " + s + " is not exists in the .env file")
}

func (c *DefaultConfig) GetInt(s string) int {
	valueStr := c.GetString(s)
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	panic("The value of the key " + s + " in the .env file should be Integer")
}

func (c *DefaultConfig) GetBool(s string) bool {
	valStr := c.GetString(s)
	if val, err := strconv.ParseBool(valStr); err == nil {
		return val
	}

	panic("The value of the key " + s + " in the .env file should be Boolean")
}

func (c *DefaultConfig) GetFloat(s string) float64 {
	valStr := c.GetString(s)
	if val, err := strconv.ParseFloat(valStr, 64); err == nil {
		return val
	}

	panic("The value of the key " + s + " in the .env file should be float")
}

func (c *DefaultConfig) GetStringSlice(s string) []string {
	valStr := c.GetString(s)

	val := strings.Split(valStr, ",")

	return val
}

func (c *DefaultConfig) GetStringMap(s string) map[string]interface{} {
	//TODO implement me
	panic("implement me")
}
