package limitrange

import (
	"encoding/json"
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
	"lib"
)

// K8sLimitRange type
type K8sLimitRange struct {
	Kind     string
	Metadata struct {
		Name      string
		Namespace string
	}
	Spec struct {
		Limits []struct {
			Default struct {
				CPU    string
				Memory string
			}
			DefaultRequest struct {
				CPU    string
				Memory string
			}
			Type string
		}
	}
}

// K8sLimitRangeItems type
type K8sLimitRangeItems struct {
	Items []K8sLimitRange
}

// GetK8sLimitRangeItems func
func GetK8sLimitRangeItems(prefixNamespaceLimit string, environment string) {
	command := []string{
		"--kubeconfig",
		"/Users/wasilp01/kubernetes/dev-kube.config",
		"--all-namespaces=true",
		"get",
		"limits",
		"-o",
		"json",
	}

	stdout := lib.RunKubectl(command)

	var jsonOutput K8sLimitRangeItems

	json.Unmarshal(stdout, &jsonOutput)

	for _, item := range jsonOutput.Items {

		for i := range item.Spec.Limits {

			CPU(prefixNamespaceLimit, environment, item, i)

			text := fmt.Sprintf("%s_limits_memory", prefixNamespaceLimit)

			limitRangeGaugeDefaultMemory := prometheus.NewGauge(prometheus.GaugeOpts{
				Name: text,
				Help: text,
				ConstLabels: prometheus.Labels{
					"namespace":   item.Metadata.Namespace,
					"environment": environment,
					"id":          fmt.Sprintf("%d", i),
					"config":      "default",
				},
			})
			limitRangeGaugeDefaultMemory.Set(lib.CalculateMetric(item.Spec.Limits[i].Default.Memory))
			prometheus.MustRegister(limitRangeGaugeDefaultMemory)

			limitRangeGaugeDefaultRequestMemory := prometheus.NewGauge(prometheus.GaugeOpts{
				Name: text,
				Help: text,
				ConstLabels: prometheus.Labels{
					"namespace":   item.Metadata.Namespace,
					"environment": environment,
					"id":          fmt.Sprintf("%d", i),
					"config":      "defaultRequest",
				},
			})
			limitRangeGaugeDefaultRequestMemory.Set(lib.CalculateMetric(item.Spec.Limits[i].DefaultRequest.Memory))
			prometheus.MustRegister(limitRangeGaugeDefaultRequestMemory)

		}
	}
}
