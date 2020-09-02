package controllers

import (
	"banking/pkg/monitoring"
	"fmt"
	"net/http"
	"strconv"

	customRabbitmq "github.com/jackeraf/go-rabbitmq-library"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
)

// HomeHandler .
func HomeHandler(rabbitmq customRabbitmq.RabbitMqClient, metrics *monitoring.BankingMetrics) http.HandlerFunc {
	amount := 5
	return func(w http.ResponseWriter, r *http.Request) {
		metrics.NumRequests.With(prometheus.Labels{"banking": "home/requests"}).Inc()
		err := rabbitmq.Publish("Withdraw", strconv.Itoa(amount))
		if err != nil {
			logrus.WithError(err).Errorf("could not publish message to withdraw %v", err)
		}
		metrics.NumSuccess.With(prometheus.Labels{"banking": "home/success"}).Inc()
		w.WriteHeader(http.StatusOK)
		_, err = w.Write([]byte(fmt.Sprintf("Money withdrawn %v €", strconv.Itoa(amount))))
		if err != nil {
			logrus.WithError(err).Fatalf("could not Write response %v", err)
		}
	}
}
