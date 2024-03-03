package main

import (
	"classic_sc/nn"
	"fmt"
	"math"
	"math/rand"
)

func sigmoid(x float64) float64 {
	return 1.0 / (1.0 + math.Exp(x))
}

func dot_product(a []float64, b []float64) float64 {
	n := len(a)
	sum := 0.0
	for i := range n {
		sum += a[i] * b[i]
	}
	return sum
}

func main() {
	wt := []float64{0.2, 0.3, 0.5}
	bias := 0.1
	num_inputs := len(wt)

	num_data := 10
	in := make([][]float64, num_data)
	out := make([]float64, num_data)
	for i := range num_data {
		in[i] = make([]float64, num_inputs)
		for j := range num_inputs {
			in[i][j] = -1.0 + 2.0*rand.Float64()
		}
		out[i] = sigmoid(dot_product(in[i], wt) + bias)
	}

	n := nn.NewNeuron(3)

	for epoch := range 1000 {
		error := 0.0
		for i := range num_data {
			actual := n.Forward(in[i])
			e := actual - out[i]
			error += math.Abs(e)
			n.BackProp(e, in[i])
		}
		change := n.UpdateWeights()
		fmt.Printf("%4d %20.10f %20.10f\n", epoch, error, change)
		if error < 1.0e-8 {
			break
		}
	}
	fmt.Println("wts = ", n)
}
