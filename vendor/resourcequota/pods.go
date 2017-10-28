package resourcequota

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
	"lib"
)

// Pods func
func Pods(prefixNamespaceLimit string, environment string, item K8sResourceQuota) {
	text := fmt.Sprintf("%s_quota_%s_pods", prefixNamespaceLimit, "used")
	gauge := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: text,
		Help: text,
		ConstLabels: prometheus.Labels{
			"namespace":   item.Metadata.Namespace,
			"environment": environment,
		},
	})
	gauge.Set(lib.CalculateMetric(item.Status.Used.Pods))
	prometheus.MustRegister(gauge)

	text = fmt.Sprintf("%s_quota_%s_pods", prefixNamespaceLimit, "hard")
	gauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: text,
		Help: text,
		ConstLabels: prometheus.Labels{
			"namespace":   item.Metadata.Namespace,
			"environment": environment,
		},
	})
	gauge.Set(lib.CalculateMetric(item.Status.Hard.Pods))
	prometheus.MustRegister(gauge)

}
