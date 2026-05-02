package main

import (
	"log"
	"net/http"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/ruanbekker/rbkr-ops-notification-service/internal/kafka"
	"github.com/ruanbekker/rbkr-ops-notification-service/internal/service"
	"github.com/ruanbekker/rbkr-ops-notification-service/internal/metrics"
	"github.com/ruanbekker/rbkr-ops-notification-service/pkg/config"
)

func main() {
	cfg := config.Load()

	consumer := kafka.NewConsumer(cfg.KafkaBrokers)
	svc := service.New()

	go func() {
        http.Handle("/metrics", promhttp.Handler())
        log.Println("metrics server running on :2112")
        if err := http.ListenAndServe(":2112", nil); err != nil {
            log.Fatalf("metrics server failed: %v", err)
        }
    }()

    metrics.Init()

	log.Println("notification-service started...")

	consumer.Start(svc.Handle)
}
