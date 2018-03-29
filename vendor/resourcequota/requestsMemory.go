package resourcequota

import (
	"lib"

	"github.com/prometheus/client_golang/prometheus"
)

// RequestsMemory func
func RequestsMemory(desc *prometheus.Desc, ch chan<- prometheus.Metric, environment string, item K8sResourceQuota) {
	ch <- prometheus.MustNewConstMetric(
		desc,
		prometheus.GaugeValue,
		lib.CalculateMetric(item.Status.Used.RequestsMemory),
		item.Metadata.Namespace,
		environment,
		"used",
	)
	ch <- prometheus.MustNewConstMetric(
		desc,
		prometheus.GaugeValue,
		lib.CalculateMetric(item.Status.Hard.RequestsMemory),
		item.Metadata.Namespace,
		environment,
		"hard",
	)
}
