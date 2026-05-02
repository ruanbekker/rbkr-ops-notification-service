package service

import (
	"encoding/json"
	"log"
)

type NotificationService struct{}

func New() *NotificationService {
	return &NotificationService{}
}

type OrderReservedEvent struct {
	OrderID string `json:"order_id"`
	Status  string `json:"status"`
}

type OrderFailedEvent struct {
	OrderID string `json:"order_id"`
	Reason  string `json:"reason"`
}

func (s *NotificationService) Handle(topic string, msg []byte) {
	switch topic {

	case "order_reserved":
		var e OrderReservedEvent
		if err := json.Unmarshal(msg, &e); err != nil {
			log.Println("parse error:", err)
			return
		}

		log.Printf("📦 Order %s reserved successfully!\n", e.OrderID)

	case "order_failed":
		var e OrderFailedEvent
		if err := json.Unmarshal(msg, &e); err != nil {
			log.Println("parse error:", err)
			return
		}

		log.Printf("❌ Order %s failed: %s\n", e.OrderID, e.Reason)
	}
}
