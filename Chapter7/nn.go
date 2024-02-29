package main

import (
	"errors"
	"fmt"
	"math"
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
