package resourcequota

import (
	"github.com/prometheus/client_golang/prometheus"
)

// K8sResourceQuota type
type K8sResourceQuota struct {
	Kind     string
	Metadata struct {
		Name      string
		Namespace string
	}
	Status struct {
		Hard struct {
			LimitsCPU      string `json:"limits.cpu"`
			LimitsMemory   string `json:"limits.memory"`
			Pods           string
			RequestsCPU    string `json:"requests.cpu"`
			RequestsMemory string `json:"requests.memory"`
		}
		Used struct {
			LimitsCPU      string `json:"limits.cpu"`
			LimitsMemory   string `json:"limits.memory"`
			Pods           string
			RequestsCPU    string `json:"requests.cpu"`
			RequestsMemory string `json:"requests.memory"`
		}
	}
}

// K8sResourceQuotaItems type
type K8sResourceQuotaItems struct {
	Items []K8sResourceQuota
}

// GetK8sResourceQuotaItems func
func GetK8sResourceQuotaItems(prefixNamespaceLimit string, environment string) {
	collector := NewResourceQuotaCollector(prefixNamespaceLimit, environment)
	prometheus.MustRegister(collector)
}
