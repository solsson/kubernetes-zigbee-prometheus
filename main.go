package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/solsson/go-conbee/sensors"

	"go.uber.org/zap"
)

var (
	conbeeHost string
	conbeeKey string
	metricsEndpoint string
	metricsListen string
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: kubernetes-zigbee-prometheus -host=[string] -key=[string]\n")
	flag.PrintDefaults()
	os.Exit(2)
}

func init() {
	on := new(bool)
	*on = true
	flag.StringVar(&conbeeHost, "host", "127.0.0.1:80", "conbee host addr")
	flag.StringVar(&conbeeKey, "key", "", "conbee api key")
	flag.StringVar(&metricsEndpoint, "endpoint", "/metrics", "Metrics endpoint path")
	flag.StringVar(&metricsListen, "listen", ":8080", "Metrics http listen")
	flag.Parse()
	flag.Usage = usage
}

func main() {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()
	if conbeeKey != "" {
		logger.Info("failed to fetch URL",
			zap.String("host", conbeeHost),
			zap.Int("keylength", len(conbeeKey)),
		)

		ss = sensors.New(conbeeHost, conbeeKey)
		foo := newDeconzCollector(logger, ss)
		prometheus.MustRegister(foo)
		logger.Info("Listening",
			zap.String("endpoint", metricsEndpoint),
			zap.String("on", metricsListen),
		)

		http.Handle(metricsEndpoint, promhttp.Handler())
		
		http.ListenAndServe(metricsListen, nil)
	} else {
		usage()
	}
}
