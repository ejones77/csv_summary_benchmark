package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
)

func csv2float(r io.Reader) (map[string][]float64, error) {
	cr := csv.NewReader(r)

	records, err := cr.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to read data from file: %w", ErrInvalidFormat)
	}

	if len(records) == 0 {
		return nil, fmt.Errorf("file is empty")
	}

	headers := records[0]
	data := make(map[string][]float64)

	// Iterate through csv, skip header row
	for _, row := range records[1:] {
		for i, val := range row {
			value, err := strconv.ParseFloat(val, 64)
			if err != nil {
				// Ignore any non-numeric columns
				continue
			}

			// Use header as key in map
			header := headers[i]
			data[header] = append(data[header], value)
		}
	}

	return data, nil
}
