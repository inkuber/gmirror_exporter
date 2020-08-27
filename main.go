package main

import (
	"net/http"

	"fmt"
	"github.com/inkuber/gmirror_exporter/src/gmirror"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"time"
)

type Metrics struct {
	MirrorStatus *prometheus.GaugeVec
}

var metrics Metrics

func main() {

	InitMetrics()
	go GMirrorStatus()

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":9112", nil)
}

func InitMetrics() {
	metrics.MirrorStatus = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "gmirror_mirror_status",
			Help: "The status of mirror",
		},
		[]string{"mirror"},
	)
}

func GMirrorStatus() {
	for {
		gm := gmirror.NewGMirror()
		status, err := gm.Status()

		if err != nil {
			panic(err)
		}

		for _, mirror := range status.Mirrors {
			state := 0.0
			if mirror.State != "COMPLETE" {
				state = 1.0
			}

			fmt.Printf("gmirror_mirror_status{%s}=%.2f\n", mirror.Name, state)
			metrics.MirrorStatus.WithLabelValues(mirror.Name).Set(state)
		}

		time.Sleep(60 * time.Second)
	}
}
