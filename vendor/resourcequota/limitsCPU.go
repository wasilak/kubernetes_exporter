package resourcequota

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
	"lib"
)

// LimitsCPU func
func LimitsCPU(prefixNamespaceLimit string, environment string, item K8sResourceQuota) {
	text := fmt.Sprintf("%s_quota_%s_limits_cpu", prefixNamespaceLimit, "used")
	gauge := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: text,
		Help: text,
		ConstLabels: prometheus.Labels{
			"namespace":   item.Metadata.Namespace,
			"environment": environment,
		},
	})
	gauge.Set(lib.CalculateMetric(item.Status.Used.LimitsCPU))
	prometheus.MustRegister(gauge)

	text = fmt.Sprintf("%s_quota_%s_limits_cpu", prefixNamespaceLimit, "hard")
	gauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: text,
		Help: text,
		ConstLabels: prometheus.Labels{
			"namespace":   item.Metadata.Namespace,
			"environment": environment,
		},
	})
	gauge.Set(lib.CalculateMetric(item.Status.Hard.LimitsCPU))
	prometheus.MustRegister(gauge)

}
