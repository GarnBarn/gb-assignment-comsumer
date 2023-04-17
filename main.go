package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/GarnBarn/common-go/httpserver"
	"github.com/GarnBarn/gb-assignment-consumer/config"
	"github.com/GarnBarn/gb-assignment-consumer/processor"
	"github.com/sirupsen/logrus"
	"github.com/wagslane/go-rabbitmq"
)

var (
	appConfig config.Config
)

func init() {
	appConfig = config.Load()
}

func main() {
	// Connect RabbitMQ
	conn, err := rabbitmq.NewConn(
		appConfig.RABBITMQ_CONNECTION,
		rabbitmq.WithConnectionOptionsLogging,
	)
	if err != nil {
		logrus.Fatal(err)
	}

	// Start HealthChecking Server
	go func() {
		httpServer := httpserver.NewHttpServer()
		logrus.Info("Listening and serving HTTP on :", appConfig.HTTP_SERVER_PORT)
		httpServer.Run(fmt.Sprint(":", appConfig.HTTP_SERVER_PORT))
	}()

	// Create Processor
	processor := processor.NewProcessor()

	consumer, err := rabbitmq.NewConsumer(
		conn,
		processor.Process,
		appConfig.RABBITMQ_ASSIGNMENT_QUEUE,
		rabbitmq.WithConsumerOptionsQueueDurable,
	)
	if err != nil {
		logrus.Fatal(err)
	}

	gracefulStop := make(chan os.Signal, 1)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)
	<-gracefulStop

	logrus.Info("Gracefully shutting down.")
	consumer.Close()
	conn.Close()

	logrus.Info("Successfully shutting down the amqp. Bye!!")
}
