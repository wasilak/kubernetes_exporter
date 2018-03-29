package resourcequota

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
	"lib"
)

// RequestsMemory func
func RequestsMemory(prefixNamespaceLimit string, environment string, item K8sResourceQuota) {
	text := fmt.Sprintf("%s_quota_%s_requests_memory", prefixNamespaceLimit, "used")
	gauge := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: text,
		Help: text,
		ConstLabels: prometheus.Labels{
			"namespace":   item.Metadata.Namespace,
			"environment": environment,
		},
	})
	gauge.Set(lib.CalculateMetric(item.Status.Used.RequestsMemory))
	prometheus.MustRegister(gauge)

	text = fmt.Sprintf("%s_quota_%s_requests_memory", prefixNamespaceLimit, "hard")
	gauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: text,
		Help: text,
		ConstLabels: prometheus.Labels{
			"namespace":   item.Metadata.Namespace,
			"environment": environment,
		},
	})
	gauge.Set(lib.CalculateMetric(item.Status.Hard.RequestsMemory))
	prometheus.MustRegister(gauge)

}