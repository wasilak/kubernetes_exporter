package resourcequota

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"

	"lib"
)

// RequestsCPU func
func RequestsCPU(prefixNamespaceLimit string, environment string, item K8sResourceQuota) {
	text := fmt.Sprintf("%s_quota_%s_requests_cpu", prefixNamespaceLimit, "used")
	gauge := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: text,
		Help: text,
		ConstLabels: prometheus.Labels{
			"namespace":   item.Metadata.Namespace,
			"environment": environment,
		},
	})
	gauge.Set(lib.CalculateMetric(item.Status.Used.RequestsCPU))
	prometheus.MustRegister(gauge)

	text = fmt.Sprintf("%s_quota_%s_requests_cpu", prefixNamespaceLimit, "hard")
	gauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: text,
		Help: text,
		ConstLabels: prometheus.Labels{
			"namespace":   item.Metadata.Namespace,
			"environment": environment,
		},
	})
	gauge.Set(lib.CalculateMetric(item.Status.Hard.RequestsCPU))
	prometheus.MustRegister(gauge)

}
