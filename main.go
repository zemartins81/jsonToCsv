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

	//Abre o arquivo de origem JSON.
	sourceFile, err := os.Open(source)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	// Decodifica o arquivo JSON para a estrutura FruitAndVegetableRank.
	var ranking []FruitAndVegetableRank
	if err := json.NewDecoder(sourceFile).Decode(&ranking); err != nil {
		return err
	}

	// Cria o arquivo de destino CSV.
	outputFile, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	// Cria um escritor CSV.
	writer := csv.NewWriter(outputFile)
	defer writer.Flush()

	// Define o cabeçalho das colunas do arquivo CSV.
	header := []string{"vegetable", "fruit", "rank"}
	if err := writer.Write(header); err != nil {
		return err
	}

	// Escreve cada linha do arquivo CSV com os dados da estrutura FruitAndVegetableRank.
	for _, r := range ranking {
		var csvRow []string
		csvRow = append(csvRow, r.Vegetable, r.Fruit, fmt.Sprint(r.Rank))
		if err := writer.Write(csvRow); err != nil {
			return nil
		}
	}
	return nil
}
