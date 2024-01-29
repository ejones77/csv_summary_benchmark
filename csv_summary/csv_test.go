package main

import (
	"strings"
	"testing"
)

func TestSummarizeData(t *testing.T) {
	data := map[string][]float64{
		"A": {11, 22, 33, 44, 55},
		"B": {66, 77, 55, 33, 22},
		"C": {6.6, 7.7, 5.5, 3.3, 2.2},
	}

	summary := summarizeData(data)

	if !strings.Contains(summary, "A") {
		t.Fatalf("expected summary to contain 'A'")
	}

}

func TestCsv2Float(t *testing.T) {
	csvData := `A,B,C,D
1.1,2,Cat,1/1/2024
4.4,5,Dog,1/2/2024
3.3,3,Bird,1/3/2024
`
	r := strings.NewReader(csvData)
	data, err := csv2float(r)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(data) != 2 {
		t.Fatalf("expected 2 headers, got %d", len(data))
	}

	if len(data["A"]) != 3 {
		t.Fatalf("expected 3 values for header A, got %d", len(data["A"]))
	}
}
