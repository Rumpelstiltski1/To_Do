package metricks

import "github.com/prometheus/client_golang/prometheus"

var (
	HttpRequestTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_request_total",
			Help: "Общее количество запросов",
		},
		[]string{"method", "path"},
	)

	HttpRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Длительность HTTP-запросов",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "path"},
	)

	HttpResponseStatus = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_response_status_total",
			Help: "Общее количество HTTP-ответов по статусам",
		},
		[]string{"status"},
	)
)

func InitMetriks() {
	prometheus.MustRegister(HttpRequestTotal)
	prometheus.MustRegister(HttpRequestDuration)
	prometheus.MustRegister(HttpResponseStatus)

}
