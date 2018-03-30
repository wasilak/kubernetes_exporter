package limitrange

import (
	"github.com/prometheus/client_golang/prometheus"
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
func GetK8sLimitRangeItems(prefixNamespaceLimit string, environment string, kubeconfig string) {
	collector := NewLimitRangeCollector(prefixNamespaceLimit, environment, kubeconfig)
	prometheus.MustRegister(collector)
}
