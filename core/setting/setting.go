package setting

import (
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	"log"
)

type Database struct {
	User               string `env:"MYSQL_USER"`
	Password           string `env:"MYSQL_PASSWORD"`
	Host               string `env:"MYSQL_HOST"`
	Name               string `env:"MYSQL_NAME"`
	MaxIdleConnections int    `env:"MYSQL_MAXIDLECONNECTIONS"`
	MaxOpenConnections int    `env:"MYSQL_MAXOPENCONNECTIONS"`
}

type RabbitMQ struct {
	User     string `env:"RABBITMQ_USER"`
	Password string `env:"RABBITMQ_PASS"`
	Host     string `env:"RABBITMQ_HOST"`
}

func NewRabbitMQSetting(configFilePath string) *RabbitMQ {
	err := godotenv.Load(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load RabbitMQ configuration")
	}

	rabbitMQ := RabbitMQ{}

	err = env.Parse(&rabbitMQ)
	if err != nil {
		log.Fatalf("Failed to parse Rabbit: %v", err)
	}
	return &rabbitMQ
}

func NewDatabaseSetting(configFilePath string) *Database {
	err := godotenv.Load(configFilePath)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}
	database := Database{}
	err = env.Parse(&database)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}
	return &database
}
