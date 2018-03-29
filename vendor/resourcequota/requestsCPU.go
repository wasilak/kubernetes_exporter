package resourcequota

import (
	"lib"

	"github.com/prometheus/client_golang/prometheus"
)

// RequestsCPU func
func RequestsCPU(desc *prometheus.Desc, ch chan<- prometheus.Metric, environment string, item K8sResourceQuota) {
	ch <- prometheus.MustNewConstMetric(
		desc,
		prometheus.GaugeValue,
		lib.CalculateMetric(item.Status.Used.RequestsCPU),
		item.Metadata.Namespace,
		environment,
		"used",
	)
	ch <- prometheus.MustNewConstMetric(
		desc,
		prometheus.GaugeValue,
		lib.CalculateMetric(item.Status.Hard.RequestsCPU),
		item.Metadata.Namespace,
		environment,
		"hard",
	)
}
