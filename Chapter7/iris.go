package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type IrisDataLines struct {
	feature    [4]float64
	norm_input [4]float64
	class      string
	class_int  int
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

	for i := 0; i < 4; i++ {
		lo := data[0].feature[i]
		hi := data[0].feature[i]

		for j := 1; j < len(data); j++ {
			v := data[j].feature[i]
			if v > hi {
				hi = v
			}
			if v < lo {
				lo = v
			}
		}

		for j := 0; j < len(data); j++ {
			v := data[j].feature[i]
			data[j].norm_input[i] = (v - lo) / (hi - lo)
		}
	}

	classes := make(map[string]int)
	num_classes := 0

	for j := 0; j < len(data); j++ {
		v := data[j].class

		cl, ok := classes[v]
		if !ok {
			cl = num_classes
			num_classes++
			classes[v] = cl
		}
		data[j].class_int = cl
	}

	fmt.Println(data)
}
