package main

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
