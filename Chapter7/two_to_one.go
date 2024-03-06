package main

import (
	"classic_sc/nn"
	"fmt"
	"math/rand"
)

func main() {
	num_inputs := 2
	num_outputs := 1

	nnet1 := nn.NewNetwork(num_inputs, num_outputs)

	var data [](nn.NNData)

	num_data := 10
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

	nnet2 := nn.NewNetwork(num_inputs, num_outputs)
	nnet2.Train(data, 1000)

	fmt.Println("True Values:")
	fmt.Println(nnet1)

	fmt.Println("Learned Values:")
	fmt.Println(nnet2)
}
