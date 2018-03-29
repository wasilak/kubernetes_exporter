package resourcequota

import (
	"lib"

	"github.com/prometheus/client_golang/prometheus"
)

// LimitsMemory func
func LimitsMemory(desc *prometheus.Desc, ch chan<- prometheus.Metric, environment string, item K8sResourceQuota) {
	ch <- prometheus.MustNewConstMetric(
		desc,
		prometheus.GaugeValue,
		lib.CalculateMetric(item.Status.Used.LimitsMemory),
		item.Metadata.Namespace,
		environment,
		"used",
	)
	ch <- prometheus.MustNewConstMetric(
		desc,
		prometheus.GaugeValue,
		lib.CalculateMetric(item.Status.Hard.LimitsMemory),
		item.Metadata.Namespace,
		environment,
		"hard",
	)
}
