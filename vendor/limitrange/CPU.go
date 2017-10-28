package limitrange

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
	"lib"
)

// CPU func
func CPU(prefixNamespaceLimit string, environment string, item K8sLimitRange, i int) {
	text := fmt.Sprintf("%s_limits_cpu", prefixNamespaceLimit)

	limitRangeGaugeDefaultCPU := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: text,
		Help: text,
		ConstLabels: prometheus.Labels{
			"namespace":   item.Metadata.Namespace,
			"environment": environment,
			"id":          fmt.Sprintf("%d", i),
			"config":      "default",
		},
	})
	limitRangeGaugeDefaultCPU.Set(lib.CalculateMetric(item.Spec.Limits[i].Default.CPU))
	prometheus.MustRegister(limitRangeGaugeDefaultCPU)

	limitRangeGaugeDefaultRequestCPU := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: text,
		Help: text,
		ConstLabels: prometheus.Labels{
			"namespace":   item.Metadata.Namespace,
			"environment": environment,
			"id":          fmt.Sprintf("%d", i),
			"config":      "defaultRequest",
		},
	})
	limitRangeGaugeDefaultRequestCPU.Set(lib.CalculateMetric(item.Spec.Limits[i].DefaultRequest.CPU))
	prometheus.MustRegister(limitRangeGaugeDefaultRequestCPU)
}
