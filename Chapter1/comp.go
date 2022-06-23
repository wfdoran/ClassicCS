package main

import (
	"errors"
	"fmt"
	"strings"
)

type bitField struct {
	num_bits int
	bits     []uint64
}

func (b *bitField) Add(x int) error {
	if x != 0 && x != 1 {
		return errors.New("Invalid x")
	}
	if b.num_bits%64 == 0 {
		b.bits = append(b.bits, uint64(0))
	}
	if x == 1 {
		word := b.num_bits / 64
		shift := b.num_bits % 64
		b.bits[word] ^= uint64(1) << shift
	}
	b.num_bits++
	return nil
}

func (b bitField) String() string {
	s := ""
	for i := 0; i < b.num_bits; i++ {
		word := i / 64
		shift := i % 64
		if (b.bits[word]>>shift)&uint64(1) == uint64(1) {
			s += string('1')
		} else {
			s += string('0')
		}
	}
	return s
}

func CompressGene(s string) (bitField, error) {
	var rv bitField

	for _, c := range strings.ToUpper(s) {
		switch c {
		case 'A':
			rv.Add(0)
			rv.Add(0)
		case 'C':
			rv.Add(0)
			rv.Add(1)
		case 'G':
			rv.Add(1)
			rv.Add(0)
		case 'T':
			rv.Add(1)
			rv.Add(1)
		default:
			err := errors.New("Invalid characeter")
			return rv, err
		}
	}

	return rv, nil
}

func DecompressGene(b bitField) string {
	rv := ""

	h := map[int]string{
		0: "A",
		1: "G",
		2: "C",
		3: "T",
	}

	for i := 0; i < b.num_bits; i += 2 {
		word := i / 64
		shift := i % 64

		v := int(b.bits[word]>>shift) & 3

		rv += h[v]
	}
	return rv
}

func main() {
	s := "acgtACGTaaaaGGGGTTgtacaggggAAAATgacT"

	s_comp, err := CompressGene(s)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(s)
		fmt.Println(s_comp)
		s2 := DecompressGene(s_comp)
		fmt.Println(s2)
	}
}
