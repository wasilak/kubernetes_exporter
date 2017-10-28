package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"limitrange"
	"resourcequota"
)

var (
	addr           = flag.String("listen-address", ":8080", "The address to listen on for HTTP requests")
	environment    = flag.String("environment", "dev", "Cluster environment")
	metricPrefix   = flag.String("prefix", "k8s_namespace", "Metrics prefix")
	metricEndpoint = flag.String("endpoint", "/metrics", "Metrics prefix")
)

func main() {
	flag.Parse()

	fmt.Println("Started Prometheus Exporter for Kubernetes")

	fmt.Println("Bootstrapping LimitRange stats")
	limitrange.GetK8sLimitRangeItems(*metricPrefix, *environment)

	fmt.Println("Bootstrapping ResourceQuota stats")
	resourcequota.GetK8sResourceQuotaItems(*metricPrefix, *environment)

	fmt.Printf("Starting web server at %s\n", *addr)
	fmt.Printf("Stats exposed at %s\n", *metricEndpoint)
	fmt.Printf("Can be accessed at %s%s\n", *addr, *metricEndpoint)

	http.Handle(*metricEndpoint, promhttp.Handler())
	log.Fatal(http.ListenAndServe(*addr, nil))
}
