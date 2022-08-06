package main

import (
	"classic_sc/genetic"
	"fmt"
	"math/rand"
	"time"
)

type assignment struct {
	letters [10]rune
}

func (a *assignment) normalize() {
	var empty []int
	for i, v := range a.letters {
		if v == ' ' {
			empty = append(empty, i)
		}
	}

	letters := []rune{'S', 'E', 'N', 'D', 'M', 'O', 'R', 'Y'}
	for _, target := range letters {
		var used []int

		for i, v := range a.letters {
			if v == target {
				used = append(used, i)
			}
		}

		if len(used) > 1 {
			keep := rand.Intn(len(used))

			for i := 0; i < len(used); i++ {
				if i != keep {
					idx := used[i]
					a.letters[idx] = ' '
					empty = append(empty, idx)
				}
			}
		}
	}

	for _, target := range letters {
		var used []int

		for i, v := range a.letters {
			if v == target {
				used = append(used, i)
			}
		}

		if len(used) == 0 {
			n := len(empty)
			i := rand.Intn(n)
			idx := empty[i]
			empty[i] = empty[n-1]
			empty = empty[:n-1]

			a.letters[idx] = target
		}
	}
}

func RandomAssignment() assignment {
	var x assignment
	x.letters = [10]rune{'S', 'E', 'N', 'D', 'M', 'O', 'R', 'Y', ' ', ' '}

	for i, _ := range x.letters {
		j := rand.Intn(i + 1)
		x.letters[i], x.letters[j] = x.letters[j], x.letters[i]
	}
	return x
}

func (a assignment) index(ch rune) int {
	for i, v := range a.letters {
		if v == ch {
			return i
		}
	}
	panic("nooo")
	return -1
}

func SMFitness(a assignment) float64 {
	s := a.index('S')
	e := a.index('E')
	n := a.index('N')
	d := a.index('D')
	m := a.index('M')
	o := a.index('O')
	r := a.index('R')
	y := a.index('Y')

	send := 1000*s + 100*e + 10*n + d
	more := 1000*m + 100*o + 10*r + e
	money := 10000*m + 1000*o + 100*n + 10*e + y

	diff := send + more - money
	if diff < 0 {
		diff = -diff
	}

	if s == 0 {
		diff += 10000
	}
	if m == 0 {
		diff += 10000
	}

	return 1.0 / float64(1+diff)
}

func (a assignment) String() string {
	s := a.index('S')
	e := a.index('E')
	n := a.index('N')
	d := a.index('D')
	m := a.index('M')
	o := a.index('O')
	r := a.index('R')
	y := a.index('Y')

	send := 1000*s + 100*e + 10*n + d
	more := 1000*m + 100*o + 10*r + e
	money := 10000*m + 1000*o + 100*n + 10*e + y

	return fmt.Sprintf("%d + %d = %d (%d)", send, more, money, send+more-money)
}

func SMCrossover(a0 assignment, a1 assignment) (assignment, assignment) {
	var out0 assignment
	var out1 assignment
	for i := 0; i < 10; i++ {
		if rand.Float64() < .5 {
			out0.letters[i] = a0.letters[i]
			out1.letters[i] = a1.letters[i]
		} else {
			out0.letters[i] = a1.letters[i]
			out1.letters[i] = a0.letters[i]

		}
	}
	out0.normalize()
	out1.normalize()

	return out0, out1
}

func SMMutate(a assignment) assignment {
	idx0 := rand.Intn(10)
	idx1 := (idx0 + 1 + rand.Intn(9)) % 10

	var out assignment
	for i := 0; i < 10; i++ {
		out.letters[i] = a.letters[i]
	}
	out.letters[idx0], out.letters[idx1] = out.letters[idx1], out.letters[idx0]
	return out
}

func main() {
	rand.Seed(time.Now().UnixNano())
	pop_size := 32
	thresh := 1.0
	max_generations := 1000
	mutation_prob := 0.1
	crossover_prob := 0.7
	ga := genetic.New[assignment](
		pop_size,
		thresh,
		max_generations,
		mutation_prob,
		crossover_prob,
		genetic.ROULETTE,
		SMFitness,
		SMCrossover,
		SMMutate,
		RandomAssignment)
	best := ga.Run()
	fmt.Println(best, SMFitness(best))
	fmt.Println(best.letters)
}
