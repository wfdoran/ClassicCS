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
	weights      []float64
	activation   ActivationFunc
	derivative   ActivationFunc
	output_cache float64
	delta        float64
}

func NewNeuron(weights []float64) *Neuron {
	return &Neuron{
		weights:      weights,
		activation:   Sigmoid,
		derivative:   DerivativeSigmoid,
		output_cache: 0.0,
		delta:        0.0,
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

func NewInputLayer(num_neurons int, num_inputs int) *Layer {
	var layer Layer

	layer.previous = nil

	for i := 0; i < num_neurons; i++ {
		wts := make([]float64, num_inputs)
		for j := 0; j < num_inputs; j++ {
			wts[j] = rand.Float64()
		}

		layer.neurons = append(layer.neurons, NewNeuron(wts))
	}

	return &layer
}

func NewLayer(num_neurons int, previous *Layer) *Layer {
	var layer Layer

	layer.previous = previous

	num_inputs := len(previous.neurons)
	for i := 0; i < num_neurons; i++ {
		wts := make([]float64, num_inputs)
		for j := 0; j < num_inputs; j++ {
			wts[j] = -1.0 + 2.0*rand.Float64()
		}

		layer.neurons = append(layer.neurons, NewNeuron(wts))
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

func (layer *Layer) CalculateDeltas(err []float64) {
	k := len(layer.neurons[0].weights)
	propagation_err := make([]float64, k)
	for i, n := range layer.neurons {
		n.delta = n.derivative(n.output_cache) * err[i]
		for j, wt := range layer.neurons[i].weights {
			propagation_err[j] += n.delta * wt
		}
	}

	if layer.previous != nil {
		layer.previous.CalculateDeltas(propagation_err)
	}
}

type Network struct {
	num_inputs    int
	layers        []*Layer
	learning_rate float64
}

func NewNetwork(num_inputs int, learning_rate float64, layer_size []int) *Network {
	var nn Network
	nn.num_inputs = num_inputs
	nn.learning_rate = learning_rate

	{
		layer := NewInputLayer(layer_size[0], num_inputs)
		nn.layers = append(nn.layers, layer)
	}

	for i := 1; i < len(layer_size); i++ {
		layer := NewLayer(layer_size[i], nn.layers[i-1])
		nn.layers = append(nn.layers, layer)
	}
	return &nn
}

func (n *Network) UpdateWeights() {
}
