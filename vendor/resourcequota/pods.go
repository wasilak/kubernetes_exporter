package resourcequota

import (
	"lib"

	"github.com/prometheus/client_golang/prometheus"
)

// Pods func
func Pods(desc *prometheus.Desc, ch chan<- prometheus.Metric, environment string, item K8sResourceQuota) {
	ch <- prometheus.MustNewConstMetric(
		desc,
		prometheus.GaugeValue,
		lib.CalculateMetric(item.Status.Used.Pods),
		item.Metadata.Namespace,
		environment,
		"used",
	)
	ch <- prometheus.MustNewConstMetric(
		desc,
		prometheus.GaugeValue,
		lib.CalculateMetric(item.Status.Hard.Pods),
		item.Metadata.Namespace,
		environment,
		"hard",
	)
}
