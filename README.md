# rbkr-ops-notification-service

Notification Service for Order Processing System (OPS)

## About

OPS is a hobby microservice architecture project for my own experiments in Go.

## Structure

The home repository that acts as the documentation for this project is available at:

- [rbkr-order-processing-system-microservices](https://github.com/ruanbekker/rbkr-order-processing-system-microservices)

Directory structure:

```bash
├── cmd
│   └── api
│       └── main.go
├── docker-compose.yaml
├── Dockerfile
├── go.mod
├── go.sum
├── internal
│   ├── kafka
│   │   └── consumer.go
│   └── service
│       └── notification.go
├── Makefile
├── pkg
│   └── config
│       └── config.go
└── README.md
```

## Notification Service

- Kafka consumer
- Logs / mock notifications

## Build

```bash
go mod tidy
go build -o app ./cmd/api
```
