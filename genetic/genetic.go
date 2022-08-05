package genetic

import (
	"fmt"
	"math/rand"
)

type SelectionType int32

const (
	ROULETTE = iota
	TOURNAMENT
)

type FitnessFunc[T any] func(T) float64
type CrossoverFunc[T any] func(T, T) (T, T)
type MutateFunc[T any] func(T) T
type RandomInitFunc[T any] func() T

type GeneticAlgorithm[T any] struct {
	population       []T
	threshold        float64
	max_generations  int
	mutation_chance  float64
	crossover_chance float64
	selection_type   SelectionType
	fitness          FitnessFunc[T]
	crossover        CrossoverFunc[T]
	mutate           MutateFunc[T]
	random_init      RandomInitFunc[T]
}

func New[T any](pop_size int, threshold float64, max_generations int, mutation_chance float64,
	crossover_chance float64, selection_type SelectionType, fitness FitnessFunc[T],
	crossover CrossoverFunc[T], mutate MutateFunc[T], random_init RandomInitFunc[T]) *GeneticAlgorithm[T] {

	var population []T

	for i := 0; i < pop_size; i++ {
		population = append(population, random_init())
	}

	return &GeneticAlgorithm[T]{
		population:       population,
		threshold:        threshold,
		max_generations:  max_generations,
		mutation_chance:  mutation_chance,
		crossover_chance: crossover_chance,
		selection_type:   selection_type,
		fitness:          fitness,
		crossover:        crossover,
		mutate:           mutate,
		random_init:      random_init,
	}
}

func (ga GeneticAlgorithm[T]) SelectRoulette() (T, T) {
	total := 0.0
	for _, c := range ga.population {
		total += ga.fitness(c)
	}

	x0 := total * rand.Float64()
	rv0 := -1

	for i, c := range ga.population {
		x0 -= ga.fitness(c)
		if x0 <= 0.0 {
			rv0 = i
			break
		}
	}

	x1 := (total - ga.fitness(ga.population[rv0])) * rand.Float64()
	rv1 := -1

	for i, c := range ga.population {
		if i == rv0 {
			continue
		}
		x1 -= ga.fitness(c)
		if x1 <= 0.0 {
			rv1 = i
			break
		}
	}

	return ga.population[rv0], ga.population[rv1]
}

// FixMe
func (ga GeneticAlgorithm[T]) SelectTournment(k int) (T, T) {
	return ga.population[0], ga.population[1]
}

func (ga *GeneticAlgorithm[T]) ReproduceAndReplace() {
	var new_population []T

	pop_size := len(ga.population)
	for len(new_population) < pop_size {
		var parent [2]T

		if ga.selection_type == ROULETTE {
			parent[0], parent[1] = ga.SelectRoulette()
		} else {
			parent[0], parent[1] = ga.SelectTournment(pop_size / 2)
		}

		if rand.Float64() < ga.crossover_chance {
			child0, child1 := ga.crossover(parent[0], parent[1])
			new_population = append(new_population, child0, child1)
		} else {
			new_population = append(new_population, parent[0], parent[1])
		}
	}

	ga.population = new_population[:pop_size]
}

func (ga *GeneticAlgorithm[T]) Mutate() {
	for i, c := range ga.population {
		if rand.Float64() < ga.mutation_chance {
			ga.population[i] = ga.mutate(c)
		}
	}
}

func (ga *GeneticAlgorithm[T]) Run() T {
	best := ga.population[0]
	for _, c := range ga.population[1:] {
		if ga.fitness(c) > ga.fitness(best) {
			best = c
		}
	}

	for iter := 0; iter < ga.max_generations; iter++ {
		fmt.Println(ga.population)
		if ga.fitness(best) > ga.threshold {
			return best
		}

		ga.ReproduceAndReplace()
		ga.Mutate()

		for _, c := range ga.population {
			if ga.fitness(c) > ga.fitness(best) {
				best = c
			}
		}
	}

	return best
}
