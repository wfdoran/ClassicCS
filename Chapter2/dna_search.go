package main

import "fmt"

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

func linear_contains(gene Gene, key Codon) bool {
	for _, c := range gene {
		if c == key {
			return true
		}
	}
	return false
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
}
