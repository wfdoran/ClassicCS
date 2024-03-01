package main

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
)

func dot_product(a []float64, b []float64) (float64, error) {
	if len(a) != len(b) {
		return 0.0, errors.New("lengths do not match")
	}
	n := len(a)
	sum := 0.0
	for i := range n {
		sum += a[i] * b[i]
	}
	return sum, nil
}

func sigmoid(x float64) float64 {
	return 1.0 / (1.0 + math.Exp(x))
}

func d_sigmoid(x float64) float64 {
	sig := sigmoid(x)
	return sig * (1.0 - sig)
}

type ActivationFunc func(float64) float64

type Neuron struct {
	weights       []float64
	bias          float64
	learning_rate float64
	output_cache  float64
	delta         float64

	activation ActivationFunc
	derivative ActivationFunc
}

func NewNeuron(num_inputs int) *Neuron {
	n := Neuron{
		weights:       make([]float64, num_inputs),
		bias:          -1.0 + 2.0*rand.Float64(),
		learning_rate: 0.5,
		output_cache:  0.0,
		delta:         0.0,

		activation: sigmoid,
		derivative: d_sigmoid,
	}

	for i := range num_inputs {
		n.weights[i] = -1.0 + 2.0*rand.Float64()
	}

	return &n
}

type Layer struct {
	neurons  []*Neuron
	previous *Layer
}

func NewFirstLayer(num_neurons int, num_inputs int) *Layer {
	x := Layer{
		neurons:  make([]*Neuron, num_neurons),
		previous: nil,
	}

	for i := range num_neurons {
		x.neurons[i] = NewNeuron(num_inputs)
	}

	return &x
}

func NewLayer(num_neurons int, previous *Layer) *Layer {
	x := Layer{
		neurons:  make([]*Neuron, num_neurons),
		previous: previous,
	}

	num_inputs := len(previous.neurons)

	for i := range num_neurons {
		x.neurons[i] = NewNeuron((num_inputs))
	}

	return &x
}

func main() {
	a := []float64{1.0, 2.0}
	b := []float64{3.0, 4.0}

	c, _ := dot_product(a, b)
	fmt.Println(c)

	for i := -6; i <= 6; i++ {
		x := float64(i) / 2.0
		fmt.Printf("%6.2f %8.4f %8.4f\n", x, sigmoid(x), d_sigmoid(x))
	}

}
