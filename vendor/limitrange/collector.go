package limitrange

import (
	"encoding/json"
	"fmt"
	"lib"

	"github.com/prometheus/client_golang/prometheus"
)

// Collector struct
type Collector struct {
	limitsCPU    *prometheus.Desc
	limitsMemory *prometheus.Desc
}

var (
	environment string
)

// NewLimitRangeCollector func
func NewLimitRangeCollector(prefixNamespaceLimit string, env string) *Collector {
	environment = env

	text := map[string]string{
		"limitsCPU":    fmt.Sprintf("%s_limits_cpu", prefixNamespaceLimit),
		"limitsMemory": fmt.Sprintf("%s_limits_memory", prefixNamespaceLimit),
	}

	variableLabels := []string{
		"namespace",
		"environment",
		"id",
		"kind",
	}
	return &Collector{
		limitsCPU:    prometheus.NewDesc(text["limitsCPU"], text["limitsCPU"], variableLabels, nil),
		limitsMemory: prometheus.NewDesc(text["limitsMemory"], text["limitsMemory"], variableLabels, nil),
	}
}

// Describe func
func (collector *Collector) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.limitsCPU
	ch <- collector.limitsMemory
}

//Collect implements required collect function for all promehteus collectors
func (collector *Collector) Collect(ch chan<- prometheus.Metric) {

	command := []string{
		"--kubeconfig",
		fmt.Sprintf("/Users/wasilp01/kubernetes/%s-kube.config", environment),
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
			CPU(collector.limitsCPU, ch, environment, item, i)
			Memory(collector.limitsMemory, ch, environment, item, i)
		}
	}

}
