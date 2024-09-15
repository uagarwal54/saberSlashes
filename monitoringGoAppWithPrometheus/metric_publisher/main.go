package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type (
	Device struct {
		ID       int    `json:"id'`
		Mac      string `json:"mac"`
		Firmware string `json:"firmware"`
	}
	metrics struct {
		// The gauge type of metric has 2 types one is Gauge And GaugeVec (Gauge Vector).
		// This metric type will have a single numerical value
		devices prometheus.Gauge

		// GaugeVec metric type bundles a set of gauges with the same name but with different lables.
		// Like if we want the metric to represent the number of services running on a server then we use Gauge but if we want them to be further classified by their
		// owners then we should use GaugeVec
		info *prometheus.GaugeVec
	}
)

// contain all the connected devices
var dvs []Device
var version string

func init() {
	// Dumy data
	dvs = []Device{
		{1, "5F-33-CC-1F-43-82", "2.1.6"},
		{2, "EF-2B-C4-F5-D6-34", "2.1.6"},
	}
	version = "2.10.5"
}

func main() {
	// Initialize a new registry without any collectors
	reg := prometheus.NewRegistry()

	m := NewMetrics(reg)

	m.devices.Set(float64(len(dvs)))
	m.info.With(prometheus.Labels{"version": version}).Set(1)
	// We can register the built-ion register also with prometheus which will allow us to get all the metrics along with out custom metrics
	// reg.MustRegister(collectors.NewGoCollector())
	// promHandler := promhttp.HandlerFor(reg, promhttp.HandlerOpts{Registry: reg})

	promHandler := promhttp.HandlerFor(reg, promhttp.HandlerOpts{})

	dMux := http.NewServeMux()
	dMux.HandleFunc("/devices", getDevice)

	mMux := http.NewServeMux()
	mMux.Handle("/metrics", promHandler)
	go func() {
		log.Fatal(http.ListenAndServe(":8080", dMux))
	}()
	go func() {
		log.Fatal(http.ListenAndServe(":8082", mMux))
	}()

	select {}

}

func getDevice(w http.ResponseWriter, r *http.Request) {
	b, err := json.Marshal(dvs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func NewMetrics(reg prometheus.Registerer) *metrics {
	m := &metrics{
		devices: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: "kbrm",
			Name:      "Consumers",
			Help:      "The number of consumers running. ",
		}),
		info: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "kbrm",
			Name:      "info",
			Help:      "Information about my devices",
		}, []string{"version"}),
	}
	reg.MustRegister(m.devices, m.info)
	return m
}
