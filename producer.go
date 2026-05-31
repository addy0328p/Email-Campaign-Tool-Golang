package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func loadRecipient(filePath string) error {
	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	for _, record := range records[1:] {
		fmt.Println(record)
	}
	return nil
}
