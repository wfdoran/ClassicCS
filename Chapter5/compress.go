package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
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

	fmt.Println(b.Bytes())
	return float64(len(b.Bytes()))
}

func main() {
	people := []string{"Michael", "Sarah", "Joshua", "Narine", "David",
		"Sajid", "Melanie", "Daniel", "Wei", "Dean", "Brian", "Murat", "Lisa"}

	n := data{names: people}

	fmt.Println(CFitness(n))

	n.names[0], n.names[3] = n.names[3], n.names[0]

	fmt.Println(CFitness(n))
}
