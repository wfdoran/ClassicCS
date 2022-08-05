package genetic

import "math/rand"

type Chromosome[T any] interface {
	Fitness() float64
	Crossover(Chromosome[T]) (Chromosome[T], Chromosome[T])
	Mutate() Chromosome[T]
}

type SelectionType int32

const (
	ROULETTE = iota
	TOURNAMENT
)

type RandomInstanceFunc[T any] func() Chromosome[T]

type GeneticAlgorithm[T any] struct {
	population       []Chromosome[T]
	threshold        float64
	max_generations  int
	mutation_chance  float64
	crossover_chance float64
	selection_type   SelectionType
}

func New[T any](pop_size int, threshold float64, max_generations int, mutation_chance float64,
	crossover_chance float64, selection_type SelectionType,
	RandomInstance RandomInstanceFunc[T]) *GeneticAlgorithm[T] {

	var population []Chromosome[T]
	for i := 0; i < pop_size; i++ {
		population = append(population, RandomInstance())
	}

	return &GeneticAlgorithm[T]{
		population:       population,
		threshold:        threshold,
		max_generations:  max_generations,
		mutation_chance:  mutation_chance,
		crossover_chance: crossover_chance,
		selection_type:   selection_type,
	}
}

func (ga GeneticAlgorithm[T]) SelectRoulette() (Chromosome[T], Chromosome[T]) {
	total := 0.0
	for _, c := range ga.population {
		total += c.Fitness()
	}

	x0 := total * rand.Float64()
	rv0 := -1

	for i, c := range ga.population {
		x0 -= c.Fitness()
		if x0 <= 0.0 {
			rv0 = i
			break
		}
	}

	x1 := (total - ga.population[rv0].Fitness()) * rand.Float64()
	rv1 := -1

	for i, c := range ga.population {
		if i == rv0 {
			continue
		}
		x1 -= c.Fitness()
		if x1 <= 0.0 {
			rv1 = i
			break
		}
	}

	return ga.population[rv0], ga.population[rv1]
}

// FixMe
func (ga GeneticAlgorithm[T]) SelectTournment(k int) (Chromosome[T], Chromosome[T]) {
	return ga.population[0], ga.population[1]
}

func (ga *GeneticAlgorithm[T]) ReproduceAndReplace() {
	var new_population []Chromosome[T]

	pop_size := len(ga.population)
	for len(new_population) < pop_size {
		var parent [2]Chromosome[T]

		if ga.selection_type == ROULETTE {
			parent[0], parent[1] = ga.SelectRoulette()
		} else {
			parent[0], parent[1] = ga.SelectTournment(pop_size / 2)
		}

		if rand.Float64() < ga.crossover_chance {
			child0, child1 := parent[0].Crossover(parent[1])
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
			ga.population[i] = c.Mutate()
		}
	}
}

func (ga *GeneticAlgorithm[T]) Run() Chromosome[T] {
	best := ga.population[0]
	for _, c := range ga.population[1:] {
		if c.Fitness() > best.Fitness() {
			best = c
		}
	}

	for iter := 0; iter < ga.max_generations; iter++ {
		if best.Fitness() > ga.threshold {
			return best
		}

		ga.ReproduceAndReplace()
		ga.Mutate()

		for _, c := range ga.population {
			if c.Fitness() > best.Fitness() {
				best = c
			}
		}
	}

	return best
}
