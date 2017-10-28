package lib

import (
	"log"
	"os/exec"
	"regexp"
	"strconv"
)

// RunKubectl func
func RunKubectl(command []string) []byte {
	stdout, err := exec.Command("kubectl", command...).Output()

	if err != nil {
		log.Fatal(err)
	}

	return stdout
}

// CalculateMetric func
func CalculateMetric(value string) float64 {
	r, _ := regexp.Compile("^(\\d+)([\\D]+)$")
	regexMatch := r.FindStringSubmatch(value)

	multiplier := 1.0

	if regexMatch != nil {
		switch regexMatch[2] {
		case "Mi":
			multiplier = 1024.0 * 1024.0
		case "Gi":
			multiplier = 1024.0 * 1024.0 * 1024.0
		case "m":
			multiplier = 0.001
		}

		valueFloat, _ := strconv.ParseFloat(regexMatch[1], 64)
		return multiplier * valueFloat
	}

	valueFloat, _ := strconv.ParseFloat(value, 64)
	return valueFloat
}
