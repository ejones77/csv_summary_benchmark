package main

import (
	"fmt"
	"strings"

	"github.com/montanaflynn/stats"
)

func summarizeData(data map[string][]float64) string {
	var summary strings.Builder
	for col, values := range data {
		if len(values) == 0 {
			continue
		}
		summary.WriteString(fmt.Sprintf("%s\n", col))
		percentiles := []float64{25, 50, 75}
		descriptiveStats, _ := stats.Describe(stats.Float64Data(values), false, &percentiles)
		summary.WriteString(descriptiveStats.String(2))
		summary.WriteString("\n------------------------\n")
	}
	return summary.String()
}
