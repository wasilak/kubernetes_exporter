package resourcequota

import (
	"lib"

	"github.com/prometheus/client_golang/prometheus"
)

// LimitsCPU func
func LimitsCPU(desc *prometheus.Desc, ch chan<- prometheus.Metric, environment string, item K8sResourceQuota) {
	ch <- prometheus.MustNewConstMetric(
		desc,
		prometheus.GaugeValue,
		lib.CalculateMetric(item.Status.Used.LimitsCPU),
		item.Metadata.Namespace,
		environment,
		"used",
	)
	ch <- prometheus.MustNewConstMetric(
		desc,
		prometheus.GaugeValue,
		lib.CalculateMetric(item.Status.Hard.LimitsCPU),
		item.Metadata.Namespace,
		environment,
		"hard",
	)
}
