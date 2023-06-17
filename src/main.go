package main

import (
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var counter1 = prometheus.NewHistogram(prometheus.HistogramOpts{
	Name: "hello_total",
	Help: "hello",
})

var counter2 = prometheus.NewCounter(prometheus.CounterOpts{
	Name: "test_total",
	Help: "test",
})

func hello(w http.ResponseWriter, r *http.Request) {
	counter1.Observe(time.Since(time.Now()).Seconds())
	w.Write([]byte("hello"))
}

func test1(w http.ResponseWriter, r *http.Request) {
	counter2.Inc()
	w.Write([]byte("test1"))
}

func main() {
	registry := prometheus.DefaultRegisterer
	registry.MustRegister(counter1, counter2)
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/test1", test1)

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe("0.0.0.0:9999", nil)
}
