package resourcequota

import (
	"encoding/json"
	"fmt"

	"lib"
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
	command := []string{
		"--kubeconfig",
		fmt.Sprintf("/Users/wasilp01/kubernetes/%s-kube.config", environment),
		"--all-namespaces=true",
		"get",
		"quota",
		"-o",
		"json",
	}

	stdout := lib.RunKubectl(command)

	var jsonOutput K8sResourceQuotaItems

	json.Unmarshal(stdout, &jsonOutput)

	for _, item := range jsonOutput.Items {

		LimitsCPU(prefixNamespaceLimit, environment, item)
		LimitsMemory(prefixNamespaceLimit, environment, item)
		RequestsCPU(prefixNamespaceLimit, environment, item)
		RequestsMemory(prefixNamespaceLimit, environment, item)
		Pods(prefixNamespaceLimit, environment, item)

	}
}
