package main

import (
	"fmt"
	"math"
	"math/rand"
)

func calculating_pi(num_terms int) float64 {
	numerator := 4.0
	denominator := 1.0
	operation := 1.0
	pi := 0.0

	for i := 0; i < num_terms; i++ {
		pi += operation * (numerator / denominator)
		denominator += 2.0
		operation *= -1.0
	}

	return pi
}

func calculate_f(x float64, num_terms int) float64 {
	rv := 0.0
	operation := 1.0
	term := 1.0
	denominator := x

	for i := 0; i < num_terms; i++ {
		rv += operation / (term * denominator)
		operation *= -1.0
		term += 2.0
		denominator *= (x * x)
	}
	return rv
}

func calculating_pi2(num_terms int) float64 {
	return 4.0 * (calculate_f(2.0, num_terms) + calculate_f(3.0, num_terms))
}

func calculating_pi3(num_trials int) float64 {
	hits := 0
	for i := 0; i < num_trials; i++ {
		x := rand.Float64()
		y := rand.Float64()
		d := x*x + y*y
		if d <= 1.0 {
			hits++
		}
	}
	return 4.0 * float64(hits) / float64(num_trials)
}

func main() {
	for num_terms := 1; num_terms <= 100; num_terms++ {
		pi := calculating_pi(num_terms)
		diff := math.Abs(pi - math.Pi)
		fmt.Printf("%3d %20.12f %20.12f\n", num_terms, pi, diff)
	}
	fmt.Println()
	for num_terms := 1; num_terms <= 100; num_terms++ {
		pi := calculating_pi2(num_terms)
		diff := math.Abs(pi - math.Pi)
		fmt.Printf("%3d %20.12f %20.12f\n", num_terms, pi, diff)
	}
	fmt.Println()
	for num_terms := 1; num_terms <= 100; num_terms++ {
		pi := calculating_pi3(num_terms * num_terms)
		diff := math.Abs(pi - math.Pi)
		fmt.Printf("%3d %20.12f %20.12f\n", num_terms, pi, diff)
	}

}
