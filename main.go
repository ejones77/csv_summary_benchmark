package main

import (
	"os"
)

func main() {
	file, _ := os.Open("given_files/housesInput.csv")
	data, _ := csv2float(file)

	outputFile, _ := os.Create("housesOutputGo.txt")
	defer outputFile.Close()

	for i := 0; i < 100; i++ {
		summary := summarizeData(data)
		outputFile.WriteString(summary + "\n")
	}
}
