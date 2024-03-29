package nn

import (
	"fmt"
	"math"
	"math/rand/v2"
)

func dot_product(a []float64, b []float64) float64 {
	n := len(a)
	sum := 0.0
	for i := range n {
		sum += a[i] * b[i]
	}
	return sum
}

func sigmoid(x float64) float64 {
	return 1.0 / (1.0 + math.Exp(-x))
}

func d_sigmoid(x float64) float64 {
	sig := sigmoid(x)
	return sig * (1.0 - sig)
}

type ActivationFunc func(float64) float64

type Neuron struct {
	weights       []float64
	wt_update     []float64
        update_count  int
	bias          float64
	bias_update   float64
	save_dot_prod float64
	delta         float64

	activation ActivationFunc
	derivative ActivationFunc
}

func NewNeuron(num_inputs int) *Neuron {
	n := Neuron{
		weights:   make([]float64, num_inputs),
		wt_update: make([]float64, num_inputs),
                update_count:  0,
		bias:          -1.0 + 2.0*rand.Float64(),
		bias_update:   0.0,
		save_dot_prod: 0.0,
		delta:         0.0,

		activation: sigmoid,
		derivative: d_sigmoid,
	}

	for i := range num_inputs {
		n.weights[i] = -1.0 + 2.0*rand.Float64()
		n.wt_update[i] = 0.0
	}

	return &n
}

func (n Neuron) String() string {
	rv := ""

	for _, w := range n.weights {
		rv += fmt.Sprintf("%8.4f ", w)
	}
	rv += fmt.Sprintf(" : %8.4f", n.bias)
	return rv
}

func (n *Neuron) Forward(inputs []float64) float64 {
	n.save_dot_prod = dot_product(inputs, n.weights) + n.bias
	return n.activation(n.save_dot_prod)
}

func (n *Neuron) BackProp(e float64, inputs []float64) {
	n.delta = n.derivative(n.save_dot_prod) * e
	for i := range n.weights {
		n.wt_update[i] += n.delta * inputs[i]
	}
        n.update_count++
	n.bias_update = n.delta
}

func (n *Neuron) UpdateWeights(learning_rate float64) float64 {
        if n.update_count == 0 {
           return 0.0
        }
	total := 0.0
	for i, change := range n.wt_update {
                change /= float64(n.update_count)
		change *= learning_rate
		n.weights[i] -= change
		total += math.Abs(change)
		n.wt_update[i] = 0.0
	}
        n.update_count = 0
	n.bias -= n.bias_update
	total += math.Abs(n.bias_update)
	n.bias_update = 0.0

	return total
}

type Layer struct {
	neurons     []*Neuron
	num_inputs  int
	save_inputs []float64
}

func NewLayer(num_neurons int, num_inputs int) *Layer {
	x := Layer{
		neurons:     make([]*Neuron, num_neurons),
		num_inputs:  num_inputs,
		save_inputs: make([]float64, num_inputs),
	}

	for i := range num_neurons {
		x.neurons[i] = NewNeuron(num_inputs)
	}

	return &x
}

func (x *Layer) Forward(input []float64) []float64 {
	output := make([]float64, len(x.neurons))
	for i, v := range input {
		x.save_inputs[i] = v
	}

	for j, n := range x.neurons {
		output[j] = n.Forward(input)
	}

	return output
}

func (x *Layer) BackProp(e []float64) []float64 {
	back := make([]float64, x.num_inputs)

	for i, n := range x.neurons {
		n.BackProp(e[i], x.save_inputs)
		for j, wt := range n.weights {
			back[j] += wt * n.delta
		}
	}
	return back
}

func (x *Layer) UpdateWeights(learning_rate float64) float64 {
	total := 0.0
	for _, n := range x.neurons {
		total += n.UpdateWeights(learning_rate)
	}
	return total
}

func (x Layer) String() string {
	rv := ""
	for _, n := range x.neurons {
		rv += n.String() + "\n"
	}
	return rv
}

type Network struct {
	num_inputs int
	layers     []*Layer
}

func NewNetwork(num_inputs int, num_neurons ...int) *Network {
	nn := Network{
		num_inputs: num_inputs,
		layers:     nil,
	}

	prev := num_inputs

	for _, a := range num_neurons {
		x := NewLayer(a, prev)
		nn.layers = append(nn.layers, x)
		prev = a
	}

	return &nn
}

func (nn *Network) Forward(input []float64) []float64 {
	vec := make([]float64, nn.num_inputs)
	copy(vec, input)
	num_layers := len(nn.layers)
	for i := 0; i < num_layers; i++ {
		vec = nn.layers[i].Forward(vec)
	}
	return vec
}

func (nn *Network) BackProp(e []float64) {
	vec := make([]float64, len(e))
	copy(vec, e)
	num_layers := len(nn.layers)
	for i := num_layers - 1; i >= 0; i-- {
		vec = nn.layers[i].BackProp(vec)
	}
}

func (nn *Network) UpdateWeights(learning_rate float64) float64 {
	total := 0.0
	for _, x := range nn.layers {
		total += x.UpdateWeights(learning_rate)
	}
	return total
}

func (nn *Network) TrainOneData(input []float64, expect []float64) float64 {
	predict := nn.Forward(input)

	num_outputs := len(expect)
	e := make([]float64, num_outputs)
	total_error := 0.0
	for i := range num_outputs {
		e[i] = predict[i] - expect[i]
		total_error += e[i] * e[i]
	}

	nn.BackProp(e)
	return total_error
}

func (nn Network) String() string {
	rv := ""
	for i, x := range nn.layers {
		rv += fmt.Sprintf("Layer %d\n", i)
		rv += x.String()
	}
	return rv
}

func gcd(a, b int) int {
	if a < b {
		a, b = b, a
	}
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func rand_step(n int) int {
	for {
		x := rand.IntN(n)
		if gcd(n, x) == 1 {
			return x
		}
	}
}

type NNData struct {
	Input  []float64
	Output []float64
}

func (nn *Network) Train(data []NNData, num_epochs int, update_freq int) {
        learning_rate := 0.5
	for epoch := range num_epochs {
		total_error := 0.0
		change := 0.0
		num_data := len(data)
		idx := 0
		idx_step := rand_step(num_data)
		for i := range data {
			total_error += nn.TrainOneData(data[idx].Input, data[idx].Output)
			idx = (idx + idx_step) % num_data
			if i%update_freq == update_freq-1 {
				change += nn.UpdateWeights(learning_rate)
			}
		}
		change += nn.UpdateWeights(learning_rate)
		fmt.Printf("%5d %20.10f %20.10f\n", epoch, total_error, change)
	}
}
