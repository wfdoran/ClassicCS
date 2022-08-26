package nn

import (
	"math"
	"math/rand"
)

func DotProduct(x []float64, y []float64) float64 {
	rv := 0.0
	for i, v := range x {
		rv += v * y[i]
	}
	return rv
}

func Sigmoid(x float64) float64 {
	return 1.0 / (1.0 + math.Exp(-x))
}

func DerivativeSigmoid(x float64) float64 {
	s := Sigmoid(x)
	return s * (1.0 - s)
}

type ActivationFunc func(float64) float64
type Neuron struct {
	weights       []float64
	learning_rate float64
	activation    ActivationFunc
	derivative    ActivationFunc
	output_cache  float64
	delta         float64
}

func NewNeuron(weights []float64, learning_rate float64) *Neuron {
	return &Neuron{
		weights:       weights,
		learning_rate: learning_rate,
		activation:    Sigmoid,
		derivative:    DerivativeSigmoid,
		output_cache:  0.0,
		delta:         0.0,
	}
}

func (n *Neuron) Output(inputs []float64) float64 {
	n.output_cache = DotProduct(n.weights, inputs)
	return n.activation(n.output_cache)
}

type Layer struct {
	neurons  []*Neuron
	previous *Layer
}

func NewInputLayer(num_neurons int, learning_rate float64, num_inputs int) *Layer {
	var layer Layer

	layer.previous = nil

	for i := 0; i < num_neurons; i++ {
		wts := make([]float64, num_inputs)
		for j := 0; j < num_inputs; j++ {
			wts[j] = rand.Float64()
		}

		layer.neurons = append(layer.neurons, NewNeuron(wts, learning_rate))
	}

	return &layer
}

func NewLayer(num_neurons int, learning_rate float64, previous *Layer) *Layer {
	var layer Layer

	layer.previous = previous

	num_inputs := len(previous.neurons)
	for i := 0; i < num_neurons; i++ {
		wts := make([]float64, num_inputs)
		for j := 0; j < num_inputs; j++ {
			wts[j] = -1.0 + 2.0*rand.Float64()
		}

		layer.neurons = append(layer.neurons, NewNeuron(wts, learning_rate))
	}

	return &layer
}

func (layer *Layer) Output(input []float64) []float64 {
	var inner []float64
	if layer.previous == nil {
		inner = input
	} else {
		inner = layer.previous.Output(input)
	}

	num_neurons := len(layer.neurons)
	rv := make([]float64, num_neurons)
	for i := 0; i < num_neurons; i++ {
		rv[i] = layer.neurons[i].Output(inner)
	}
	return rv
}

func (layer *Layer) CalculateDeltas(expected []float64) {
	for i, n := range layer.neurons {
		n.delta = n.derivative(n.output_cache) * (expected[i] - n.activation(n.output_cache))
	}
}

type Network struct {
	num_inputs int
	layers     []*Layer
}

func NewNetwork(num_inputs int, learning_rate float64, layer_size []int) *Network {
	var nn Network
	nn.num_inputs = num_inputs

	{
		layer := NewInputLayer(layer_size[0], learning_rate, num_inputs)
		nn.layers = append(nn.layers, layer)
	}

	for i := 1; i < len(layer_size); i++ {
		layer := NewLayer(layer_size[i], learning_rate, nn.layers[i-1])
		nn.layers = append(nn.layers, layer)
	}
	return &nn
}
