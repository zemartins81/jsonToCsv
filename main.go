package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type FruitAndVegetableRank struct {
	Vegetable string `json:"vegetable"`
	Fruit     string `json:"fruit"`
	Rank      int64  `json:"rank"`
}

func main() {
	if err := convertJsonToCsv("fruits.json", "fruits.csv"); err != nil {
		log.Fatal(err)
	}
}

func convertJsonToCsv(source, destination string) error {

	sourceFile, err := os.Open(source)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	var ranking []FruitAndVegetableRank
	if err := json.NewDecoder(sourceFile).Decode(&ranking); err != nil {
		return err
	}

	outputFile, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	writer := csv.NewWriter(outputFile)
	defer writer.Flush()

	header := []string{"vegetable", "fruit", "rank"}
	if err := writer.Write(header); err != nil {
		return err
	}

	for _, r := range ranking {
		var csvRow []string
		csvRow = append(csvRow, r.Vegetable, r.Fruit, fmt.Sprint(r.Rank))
		if err := writer.Write(csvRow); err != nil {
			return nil
		}
	}
	return nil
}
