package main

import (
	"classic_sc/genetic"
	"fmt"
	"math/rand"
	"time"
)

type pair struct {
	x int
	y int
}

func Fitness(p pair) float64 {
	v := 6*p.x - p.x*p.x + 4*p.y - p.y*p.y
	if v >= 1 {
		return float64(v)
	} else {
		return 1.0 / (2.0 - float64(v))
	}
}

func RandomPair() pair {
	return pair{x: rand.Intn(100),
		y: rand.Intn(100)}
}

func Crossover(p0 pair, p1 pair) (pair, pair) {
	out0 := pair{x: p0.x, y: p1.y}
	out1 := pair{x: p1.x, y: p0.y}
	return out0, out1
}

func Mutate(p pair) pair {
	out := pair{x: p.x, y: p.y}
	if rand.Float64() > .5 {
		if rand.Float64() > .5 {
			out.x += 1
		} else {
			out.x -= 1
		}
	} else {
		if rand.Float64() > .5 {
			out.y += 1
		} else {
			out.y -= 1
		}
	}
	return out
}

func main() {
	rand.Seed(time.Now().UnixNano())
	ga := genetic.New[pair](20, 13.0, 100, 0.1, 0.7, genetic.ROULETTE,
		Fitness, Crossover, Mutate, RandomPair)
	best := ga.Run()
	fmt.Println(best, Fitness(best))
}
