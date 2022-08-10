package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"math/rand"
	"time"
)

type data struct {
	names []string
}

func CFitness(d data) float64 {
	var b bytes.Buffer
	gz := gzip.NewWriter(&b)

	for _, name := range d.names {
		gz.Write([]byte(name))
	}
	gz.Flush()
	gz.Close()

	// fmt.Println(b.Bytes())
	return float64(len(b.Bytes()))
}

func main() {
	rand.Seed(time.Now().UnixNano())

	people := []string{"Michael", "Sarah", "Joshua", "Narine", "David",
		"Sajid", "Melanie", "Daniel", "Wei", "Dean", "Brian", "Murat", "Lisa"}

	m := len(people)
	for i := 1; i < m; i++ {
		j := rand.Intn(i + 1)
		people[i], people[j] = people[j], people[i]
	}

	n := data{names: people}

	fmt.Println(people)
	fmt.Println(CFitness(n))
}
