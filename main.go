package main

import (
	"archive/zip"
	"os"
)

func main() {
	file, err := os.Create("test.zip")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	zipWriter := zip.NewWriter(file)
	defer zipWriter.Close()

	writer, err := zipWriter.Create("newfileInZip.txt")
	if err != nil {
		panic(err)
	}

	writer.Write([]byte("write!!!!!!!!"))
}
