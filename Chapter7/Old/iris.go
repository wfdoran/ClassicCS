package main

import (
	"classic_sc/nn"
	"encoding/csv"
	"os"
	"strconv"
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

	var nn_data []nn.NNData

	for _, d := range data {
		var nn_d nn.NNData
		nn_d.Input = d.norm_input[:]
		nn_d.Output = make([]float64, num_classes)
		nn_d.Output[d.class_int] = 1.0
		// fmt.Println(d)

		nn_data = append(nn_data, nn_d)
	}

	// fmt.Println(nn_data)

	learning_rate := 1.0
	nn := nn.NewNetwork(num_features, learning_rate, []int{5, num_classes})
	nn.Train(nn_data, 2)
}
