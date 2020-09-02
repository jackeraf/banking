package main

import (
	"banking/pkg/controllers"
	"banking/pkg/monitoring"
	"net/http"

	customRabbitmq "github.com/jackeraf/go-rabbitmq-library"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
)

// Docker file
// https://stackoverflow.com/questions/47837149/build-docker-with-go-app-cannot-find-package

func main() {
	logrus.Info("Starting banking project")
	metrics := monitoring.NewMetrics()
	monitoring.MetricsToRegister(metrics)
	logrus.Info("before creating rabbitmq client")
	client, err := customRabbitmq.NewRabbitmqClient()
	if err != nil {
		metrics.NumErrors.With(prometheus.Labels{"banking": "rabbitmq errors"}).Inc()
		logrus.Fatalf("Error getting rabbit client %v\n", err)
	}
	err = client.CreateQueues([]string{
		"Withdraw",
		"Transfer",
		"Deposit",
	})
	if err != nil {
		logrus.Fatalf("Error creating rabbit queues %v\n", err)
	}
	logrus.Info("queues created successfully")

	http.Handle("/", controllers.HomeHandler(client, metrics))
	http.Handle("/metrics", promhttp.Handler())
	logrus.Fatal(http.ListenAndServe(":8000", nil))
}
