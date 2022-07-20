package main

import (
	"classic_sc/csp"
	"fmt"
)

func main() {
	csp := csp.New[string, string]()

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

	soln := csp.BacktrackSearch()

	for v, d := range soln {
		fmt.Println(v, " -> ", d)
	}

}
