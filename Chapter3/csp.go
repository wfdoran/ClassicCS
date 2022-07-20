package main

import "fmt"

type constraint_checker[V comparable, D comparable] func(map[V]D) bool

type Constraint[V comparable, D comparable] struct {
	inputs []V
	check  constraint_checker[V, D]
}

type CSP[V comparable, D comparable] struct {
	variables   map[V]bool
	domains     map[D]bool
	constraints []Constraint[V, D]
}

func New[V comparable, D comparable]() *CSP[V, D] {
	return &CSP[V, D]{variables: make(map[V]bool), domains: make(map[D]bool), constraints: nil}
}

func New2[V comparable, D comparable](variables []V, domains []D) *CSP[V, D] {
	csp := New[V, D]()
	for _, v := range variables {
		csp.variables[v] = true
	}
	for _, d := range domains {
		csp.domains[d] = true
	}
	return csp
}

func (csp *CSP[V, D]) AddVariable(vs ...V) {
	for _, v := range vs {
		csp.variables[v] = true
	}
}

func (csp *CSP[V, D]) AddDomain(ds ...D) {
	for _, d := range ds {
		csp.domains[d] = true
	}
}

func (csp *CSP[V, D]) AddConstraint(c Constraint[V, D]) {
	for _, v := range c.inputs {
		_, ok := csp.variables[v]
		if !ok {
			fmt.Println("Warning variable ", v, " not in csp")
		}
	}
	csp.constraints = append(csp.constraints, c)
}

func (csp CSP[V, D]) consistent(assignment map[V]D) bool {
	for _, c := range csp.constraints {
		if !c.check(assignment) {
			return false
		}
	}
	return true
}

func (csp *CSP[V, D]) AddConstraintNotEqual(v1 V, v2 V) {
	var c Constraint[V, D]

	c.inputs = append(c.inputs, v1, v2)

	c.check = func(assignment map[V]D) bool {
		d1, ok1 := assignment[v1]
		d2, ok2 := assignment[v2]

		if !ok1 || !ok2 {
			return true
		}
		return d1 != d2
	}

	csp.AddConstraint(c)
}

func (csp CSP[V, D]) backtrack_search(assignment map[V]D) map[V]D {
	if len(assignment) == len(csp.variables) {
		return assignment
	}

	var branch_var V

	for v, _ := range csp.variables {
		_, ok := assignment[v]
		if !ok {
			branch_var = v
			break
		}
	}

	for set_d, _ := range csp.domains {
		sub_assignment := make(map[V]D)
		for v, d := range assignment {
			sub_assignment[v] = d
		}
		sub_assignment[branch_var] = set_d

		if csp.consistent(sub_assignment) {
			result := csp.backtrack_search(sub_assignment)
			if result != nil {
				return result
			}
		}
	}
	return nil
}

func main() {
	csp := New[string, string]()

	csp.AddVariable("Western Australia",
		"Northern Territory",
		"Queensland",
		"South Australia",
		"New South Wales",
		"Victoria",
		"Tasmania")

	csp.AddDomain("red", "green", "blue")

	csp.AddConstraintNotEqual("Western Australia", "Northern Territory")
	csp.AddConstraintNotEqual("Western Australia", "South Australia")
	csp.AddConstraintNotEqual("South Australia", "Northern Territory")
	csp.AddConstraintNotEqual("Queensland", "Northern Territory")
	csp.AddConstraintNotEqual("Queensland", "South Australia")
	csp.AddConstraintNotEqual("Queensland", "New South Wales")
	csp.AddConstraintNotEqual("New South Wales", "South Australia")
	csp.AddConstraintNotEqual("Victoria", "South Australia")
	csp.AddConstraintNotEqual("Victoria", "New South Wales")
	csp.AddConstraintNotEqual("Victoria", "Tasmania")

	blank := make(map[string]string)
	soln := csp.backtrack_search(blank)

	for v, d := range soln {
		fmt.Println(v, " -> ", d)
	}

}
