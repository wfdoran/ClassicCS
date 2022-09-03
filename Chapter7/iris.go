package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type IrisDataLines struct {
	feature [4]float64
	class   string
}

func main() {
	filename := "iris.csv"
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	lines, err := csv.NewReader(file).ReadAll()
	if err != nil {
		panic(err)
	}

	var data []IrisDataLines
	for _, line := range lines {
		var d IrisDataLines
		for i := 0; i < 4; i++ {
			d.feature[i], _ = strconv.ParseFloat(line[i], 64)
		}
		d.class = line[4]
		data = append(data, d)
	}

	fmt.Println(data)
}
