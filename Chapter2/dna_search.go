package main

import (
	"fmt"
	"sort"
)

type Nucleotide int

const (
	A Nucleotide = iota + 1
	C
	G
	T
)

func (n Nucleotide) String() string {
	switch n {
	case A:
		return "A"
	case C:
		return "C"
	case G:
		return "G"
	case T:
		return "T"
	default:
		return ""
	}
}

func char_to_nucleotide(ch byte) Nucleotide {
	switch ch {
	case 'A':
		return A
	case 'C':
		return C
	case 'G':
		return G
	case 'T':
		return T
	default:
		fmt.Println("err", ch)
		return 0
	}
}

type Codon [3]Nucleotide
type Gene []Codon

func string_to_gene(s string) Gene {
	gene := Gene{}
	n := len(s)
	for i := 0; i+2 < n; i += 3 {
		codon := Codon{char_to_nucleotide(s[i]), char_to_nucleotide(s[i+1]), char_to_nucleotide(s[i+2])}
		gene = append(gene, codon)
	}

	return gene
}

func codonLess(a Codon, b Codon) bool {
	for k := 0; k < 3; k++ {
		if a[k] < b[k] {
			return true
		}
		if a[k] > b[k] {
			return false
		}
	}
	return false

}

func linear_contains[T comparable](a []T, key T) bool {
	for _, c := range a {
		if c == key {
			return true
		}
	}
	return false
}

func binary_contains(gene Gene, key Codon) bool {
	lo := -1
	hi := len(gene)

	for hi-lo > 1 {
		mid := (hi + lo) / 2
		if gene[mid] == key {
			return true
		}
		if codonLess(gene[mid], key) {
			lo = mid
		} else {
			hi = mid
		}
	}
	return false
}

func (gene Gene) Len() int {
	return len(gene)
}

func (gene Gene) Less(i, j int) bool {
	return codonLess(gene[i], gene[j])

}

func (gene Gene) Swap(i, j int) {
	gene[i], gene[j] = gene[j], gene[i]
}

func main() {
	// A nucleotide
	fmt.Println(A)

	// A codon
	x := Codon{A, A, T}
	fmt.Println(x)

	// A gene
	s := "ACGTGGCTCTCTAACGTAGG"
	gene := string_to_gene(s)
	fmt.Println(gene)

	fmt.Println(linear_contains(gene, Codon{G, T, A}))
	fmt.Println(linear_contains(gene, Codon{G, G, G}))

	sort.Sort(gene)
	fmt.Println(gene)
	fmt.Println(binary_contains(gene, Codon{G, T, A}))
	fmt.Println(binary_contains(gene, Codon{G, G, G}))
}
