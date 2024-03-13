package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"classic_sc/nn"
)

const num_features int = 4

type IrisDataLines struct {
	feature    [num_features]float64
	norm_input [num_features]float64
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
		for i := 0; i < num_features; i++ {
			d.feature[i], _ = strconv.ParseFloat(line[i], 64)
		}
		d.class = line[4]
		data = append(data, d)
	}

	for i := 0; i < num_features; i++ {
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

	// fmt.Println(data)

	var irus_data []nn.NNData

	for _, d := range data {
		var nn_d nn.NNData
		nn_d.Input = d.norm_input[:]
		nn_d.Output = make([]float64, num_classes)
		nn_d.Output[d.class_int] = 1.0

		irus_data = append(irus_data, nn_d)
	}

	nnet := nn.NewNetwork(num_features, 6, 6, num_classes)
	nnet.Train(irus_data, 10000, 25)

	for _, d := range irus_data {
		model := nnet.Forward(d.Input)
		fmt.Printf("%8.4f %8.4f %8.4f | %8.4f %8.4f %8.4f\n",
			d.Output[0], d.Output[1], d.Output[2],
			model[0], model[1], model[2])
	}
}
