package main

import (
	"classic_sc/nn"
	"fmt"
	"math/rand"
)

func main() {
	rand.Seed(1)
	num_inputs := 2
	num_hidden := 3
	num_outputs := 1

	nnet1 := nn.NewNetwork(num_inputs, num_hidden, num_outputs)

	var data [](nn.NNData)

	num_data := 10000
	for range num_data {
		in := make([]float64, num_inputs)
		for j := range num_inputs {
			in[j] = -1.0 + 2.0*rand.Float64()
		}
		out := nnet1.Forward(in)

		x := nn.NNData{
			Input:  in,
			Output: out,
		}
		data = append(data, x)
	}

	fmt.Println(data[0])

	nnet2 := nn.NewNetwork(num_inputs, num_hidden, num_outputs)

	fmt.Println("Original Values:")
	fmt.Println(nnet2)

	nnet2.Train(data, 10000, 25)

	// nnet2.TrainOneData(data[0].Input, data[0].Output)
	// nnet2.UpdateWeights()

	fmt.Println("True Values:")
	fmt.Println(nnet1)

	fmt.Println("Learned Values:")
	fmt.Println(nnet2)
}
