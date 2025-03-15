package tools

import (
	"encoding/csv"
	"os"
)

func LoadCSV(filename string) [][]string {

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	r := csv.NewReader(file)
	rows, err := r.ReadAll()
	if err != nil {
		panic(err)
	}

	return rows
}
