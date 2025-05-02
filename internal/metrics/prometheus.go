package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	HttpRequests = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "ducks_requests_total",
		Help: "total http requests",
	}, []string{"method", "path", "status"})

	RequestDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name: "ducks_request_duration_seconds",
		Help: "http request duration in seconds",
		Buckets: prometheus.DefBuckets,
	}, []string{"method", "path"})
)
