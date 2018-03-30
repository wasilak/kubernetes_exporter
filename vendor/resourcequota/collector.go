package resourcequota

import (
	"encoding/json"
	"fmt"
	"lib"

	"github.com/prometheus/client_golang/prometheus"
)

// Collector struct
type Collector struct {
	limitsCPU      *prometheus.Desc
	limitsMemory   *prometheus.Desc
	RequestsCPU    *prometheus.Desc
	RequestsMemory *prometheus.Desc
	Pods           *prometheus.Desc
}

var (
	environment string
	kubeconfig  string
)

// NewResourceQuotaCollector func
func NewResourceQuotaCollector(prefixNamespaceLimit string, env string, kubeconf string) *Collector {
	environment = env
	kubeconfig = kubeconf

	variableLabels := []string{
		"namespace",
		"environment",
		"kind",
	}

	text := map[string]string{
		"limitsCPU":      fmt.Sprintf("%s_quota_limits_cpu", prefixNamespaceLimit),
		"limitsMemory":   fmt.Sprintf("%s_quota_limits_memory", prefixNamespaceLimit),
		"RequestsCPU":    fmt.Sprintf("%s_quota_requests_cpu", prefixNamespaceLimit),
		"RequestsMemory": fmt.Sprintf("%s_quota_requests_memory", prefixNamespaceLimit),
		"Pods":           fmt.Sprintf("%s_quota_pods", prefixNamespaceLimit),
	}

	return &Collector{
		limitsCPU:      prometheus.NewDesc(text["limitsCPU"], text["limitsCPU"], variableLabels, nil),
		limitsMemory:   prometheus.NewDesc(text["limitsMemory"], text["limitsMemory"], variableLabels, nil),
		RequestsCPU:    prometheus.NewDesc(text["RequestsCPU"], text["RequestsCPU"], variableLabels, nil),
		RequestsMemory: prometheus.NewDesc(text["RequestsMemory"], text["RequestsMemory"], variableLabels, nil),
		Pods:           prometheus.NewDesc(text["Pods"], text["Pods"], variableLabels, nil),
	}
}

// Describe func
func (collector *Collector) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.limitsCPU
	ch <- collector.limitsMemory
	ch <- collector.RequestsCPU
	ch <- collector.RequestsMemory
	ch <- collector.Pods
}

//Collect implements required collect function for all promehteus collectors
func (collector *Collector) Collect(ch chan<- prometheus.Metric) {

	command := []string{
		"--all-namespaces=true",
		"get",
		"quota",
		"-o",
		"json",
	}

	if len(kubeconfig) > 0 {
		command = append(command, "--kubeconfig="+kubeconfig)
	}

	stdout := lib.RunKubectl(command)

	var jsonOutput K8sResourceQuotaItems

	json.Unmarshal(stdout, &jsonOutput)

	for _, item := range jsonOutput.Items {
		LimitsCPU(collector.limitsCPU, ch, environment, item)
		LimitsMemory(collector.limitsMemory, ch, environment, item)
		RequestsCPU(collector.RequestsCPU, ch, environment, item)
		RequestsMemory(collector.RequestsMemory, ch, environment, item)
		Pods(collector.Pods, ch, environment, item)
	}

}
