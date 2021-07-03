package main

import (
	"encoding/csv"
	"os"
)

func main() {
	file, err := os.Create("test.csv")
	if err != nil {
		panic(err)
	}
	writer := csv.NewWriter(file)
	writer.Write([]string{"a", "b", "some loooong text"})
	writer.Flush()
}
