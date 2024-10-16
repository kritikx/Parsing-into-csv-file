package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Parsing data into csv file")

	file, err := os.OpenFile("PCBCOMP21012022_210422103228.csv", os.O_RDONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}
	fp := csv.NewReader(file)

	data, err := fp.ReadAll()
	if err != nil {
		panic(err)
	}

	for _, line := range data {
		fmt.Println(line)
	}
	defer file.Close()
}
