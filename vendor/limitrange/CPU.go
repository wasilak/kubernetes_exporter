package limitrange

import (
	"fmt"

	"lib"

	"github.com/prometheus/client_golang/prometheus"
)

// CPU func
func CPU(desc *prometheus.Desc, ch chan<- prometheus.Metric, environment string, item K8sLimitRange, i int) {

	ch <- prometheus.MustNewConstMetric(
		desc,
		prometheus.GaugeValue,
		lib.CalculateMetric(item.Spec.Limits[i].Default.CPU),
		item.Metadata.Namespace,
		environment,
		fmt.Sprintf("%d", i),
		"default",
	)

	ch <- prometheus.MustNewConstMetric(
		desc,
		prometheus.GaugeValue,
		lib.CalculateMetric(item.Spec.Limits[i].DefaultRequest.CPU),
		item.Metadata.Namespace,
		environment,
		fmt.Sprintf("%d", i),
		"defaultRequest",
	)
}
