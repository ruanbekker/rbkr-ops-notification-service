package main

import (
	"log"

	"github.com/ruanbekker/rbkr-ops-notification-service/internal/kafka"
	"github.com/ruanbekker/rbkr-ops-notification-service/internal/service"
	"github.com/ruanbekker/rbkr-ops-notification-service/pkg/config"
)

func main() {
	cfg := config.Load()

	consumer := kafka.NewConsumer(cfg.KafkaBrokers)
	svc := service.New()

	log.Println("notification-service started...")

	consumer.Start(svc.Handle)
}
