package metrics

import "github.com/prometheus/client_golang/prometheus"

var NotificationsSent = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "notifications_sent_total",
		Help: "Total notifications sent",
	},
)

func Init() {
	prometheus.MustRegister(NotificationsSent)
}
