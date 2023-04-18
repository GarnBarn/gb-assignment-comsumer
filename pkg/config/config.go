package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Env                              string
	HTTP_SERVER_PORT                 string `envconfig:"HTTP_SERVER_PORT" default:"3003"`
	MYSQL_CONNECTION_STRING          string `envconfig:"MYSQL_CONNECTION_STRING"`
	RABBITMQ_CONNECTION              string `envconfig:"RABBITMQ_CONNECTION" default:"amqp://guest:guest@localhost:5672"`
	RABBITMQ_MAXIMUM_RETRY           int    `envconfig:"RABBITMQ_MAXIMUM_RETRY" default:"3"`
	RABBITMQ_DEAD_LETTERING_EXCHANGE string `envconfig:"RABBITMQ_DEAD_LETTERING_EXCHANGE" default:"assignment.fail"`
	RABBITMQ_ASSIGNMENT_EXCHANGE     string `envconfig:"RABBITMQ_ASSIGNMENT_EXCHANGE" default:"assignment"`
	RABBITMQ_ASSIGNMENT_CREATE_QUEUE string `envconfig:"RABBITMQ_ASSIGNMENT_CREATE_QUEUE" default:"assignmen.create.update"`
}

func Load() Config {
	var config Config
	ENV, ok := os.LookupEnv("ENV")
	if !ok {
		// Default value for ENV.
		ENV = "dev"
	}
	// Load the .env file only for dev env.
	ENV_CONFIG, ok := os.LookupEnv("ENV_CONFIG")
	if !ok {
		ENV_CONFIG = "./.env"
	}

	err := godotenv.Load(ENV_CONFIG)
	if err != nil {
		logrus.Warn("Can't load env file")
	}

	envconfig.MustProcess("", &config)
	config.Env = ENV

	return config
}
