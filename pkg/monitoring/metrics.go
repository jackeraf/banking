package monitoring

import "github.com/prometheus/client_golang/prometheus"

// BankingMetrics .
type BankingMetrics struct {
	NumErrors   prometheus.CounterVec
	NumSuccess  prometheus.CounterVec
	NumRequests prometheus.CounterVec
}

// NewMetrics .
func NewMetrics() *BankingMetrics {
	sucessResponses := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "banking_home_success_responses_total",
			Help: "Number of success responses for home handler endpoint",
		},
		[]string{"banking"},
	)
	errorResponses := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "banking_home_error_responses_total",
			Help: "Number of error responses for home handler endpoint",
		},
		[]string{"banking"},
	)
	numberOfRequests := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "banking_home_requests_total",
			Help: "Number of requests for home handler endpoint",
		},
		[]string{"banking"},
	)
	return &BankingMetrics{
		NumErrors:   *errorResponses,
		NumSuccess:  *sucessResponses,
		NumRequests: *numberOfRequests,
	}
}

// MetricsToRegister .
func MetricsToRegister(metrics *BankingMetrics) {
	prometheus.MustRegister(metrics.NumErrors)
	prometheus.MustRegister(metrics.NumRequests)
	prometheus.MustRegister(metrics.NumSuccess)
}
