package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	conbeeHost = "10.0.0.18"
	conbeeKey  = "0A498B9909"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: get-all-sensors -host=[string] -key=[string]\n")
	flag.PrintDefaults()
	os.Exit(2)
}

func init() {
	on := new(bool)
	*on = true
	flag.StringVar(&conbeeHost, "host", os.Getenv("DECONZ_CONBEE_HOST"), "conbee host addr")
	flag.StringVar(&conbeeKey, "key", os.Getenv("DECONZ_CONBEE_APIKEY"), "conbee api key")
	flag.Parse()
	flag.Usage = usage
}

func main() {
	if conbeeKey != "" {
		foo := newDeconzCollector(conbeeHost, conbeeKey)
		prometheus.MustRegister(foo)
	
		http.Handle("/metrics", promhttp.Handler())
		fmt.Println("Listening on port 8080")
		http.ListenAndServe(":8080", nil)
	} else {
		usage()
	}
}
