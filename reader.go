package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type LoadOptions struct {
	filepath string
}

func LoadFile(options LoadOptions) ([][]string, error) {
	file, err := os.Open(options.filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file at %s", options.filepath)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	data, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to read file at %s", options.filepath)
	}

	return data, nil
}
