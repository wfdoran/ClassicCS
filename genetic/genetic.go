package genetic

import "math/rand"

type Chromosome interface {
	Fitness() float64
	RandomInstance() Chromosome
	Crossover(Chromosome) Chromosome
	Mutate()
}

type SelectionType int32

const (
	ROULETTE = iota
	TOURNAMENT
)

type GeneticAlgorithm struct {
	population       []Chromosome
	threshold        float64
	max_generations  int
	mutation_chance  float64
	crossover_chance float64
	selection_type   SelectionType
}

func New(C Chromosome, pop_size int, threshold float64, max_generations int, mutation_chance float64,
	crossover_chance float64, selection_type SelectionType) *GeneticAlgorithm {

	var population []Chromosome
	for i := 0; i < pop_size; i++ {
		population = append(population, C.RandomInstance())
	}

	return &GeneticAlgorithm{population: population,
		threshold:        threshold,
		max_generations:  max_generations,
		mutation_chance:  mutation_chance,
		crossover_chance: crossover_chance,
		selection_type:   selection_type,
	}
}

func (ga GeneticAlgorithm) SelectRoulette() (Chromosome, Chromosome) {
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
func (ga GeneticAlgorithm) SelectTournment(k int) (Chromosome, Chromosome) {
	return ga.population[0], ga.population[1]
}

func (ga *GeneticAlgorithm) ReproduceAndReplace() {
	var new_population []Chromosome

	pop_size := len(ga.population)
	for len(new_population) < pop_size {
		var parent [2]Chromosome

		if ga.selection_type == ROULETTE {
			parent[0], parent[1] = ga.SelectRoulette()
		} else {
			parent[0], parent[1] = ga.SelectTournment(pop_size / 2)
		}

		if rand.Float64() < ga.crossover_chance {
			child := parent[0].Crossover(parent[1])
			new_population = append(new_population, child)
		} else {
			new_population = append(new_population, parent[0])
			new_population = append(new_population, parent[1])
		}
	}

	ga.population = new_population[:pop_size]
}

func (ga *GeneticAlgorithm) Mutate() {
	for _, c := range ga.population {
		if rand.Float64() < ga.mutation_chance {
			c.Mutate()
		}
	}
}

func (ga *GeneticAlgorithm) Run() Chromosome {
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
