package nn

import "math"

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
	return n.output_cache
}
