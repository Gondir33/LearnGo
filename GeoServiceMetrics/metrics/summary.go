package metrics

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	requestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "http_request_duration_seconds",
			Help: "Histogram of the time (in seconds) each request takes.",
		},
		[]string{"method"},
	)

	requestCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_request_count",
			Help: "Count of HTTP requests",
		},
		[]string{"method"},
	)

	cacheDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "cache_duration_seconds",
			Help: "Histogram of cache lookup duration (in seconds) for each method.",
		},
		[]string{"method"},
	)

	externalAPIDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "external_api_duration_seconds",
			Help: "Histogram of external API call duration (in seconds) for each method.",
		},
		[]string{"method"},
	)
)

func init() {
	prometheus.MustRegister(requestDuration)
	prometheus.MustRegister(requestCount)
	prometheus.MustRegister(cacheDuration)
	prometheus.MustRegister(externalAPIDuration)
}

func ObserveRequsetDuration(d time.Duration, method string) {
	requestDuration.WithLabelValues(method).Observe(d.Seconds())
}
func ObserveRequsetCount(method string) {
	requestCount.WithLabelValues(method).Inc()
}
func ObserveCacheDuration(d time.Duration, method string) {
	cacheDuration.WithLabelValues(method).Observe(d.Seconds())
}
func ObserveExternalAPIDuration(d time.Duration, method string) {
	externalAPIDuration.WithLabelValues(method).Observe(d.Seconds())
}
