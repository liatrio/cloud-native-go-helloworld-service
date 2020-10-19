package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("<h1>Hello Liatrio!</h1>"))

	log.Debugf("Request received on %s", r.RequestURI)

	if err != nil {
		log.Errorf("error sending response: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	hellosGiven.Inc()
}

var (
	hellosGiven = promauto.NewCounter(prometheus.CounterOpts{
		Name: "cloud_native_go_helloworld_service_hellos_given_total",
		Help: "The total number of hellos returned by the service",
	})
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.SetLevel(log.DebugLevel)

	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.Handle("/metrics", promhttp.Handler())

	log.Infof("Server starting on port: %s", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
