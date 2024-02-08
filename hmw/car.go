package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Car struct {
	Brand  string `json:"brand"`
	Year   int    `json:"year"`
	Model  string `json:"model"`
	Weight int    `json:"weight"`
}

func main() {
	filePath := "cars.json"

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var cars []Car
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&cars); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
	for _, car := range cars {
		fmt.Printf("Brand: %s, Year: %d, Model: %s, Weight: %d\n", car.Brand, car.Year, car.Model, car.Weight)
	}
}
