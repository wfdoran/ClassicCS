package main

import (
	"classic_sc/kmeans"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	k := 4
	num_pts := 16

	var points []kmeans.DataPoint
	for i := 0; i < num_pts; i++ {
		a := []float64{rand.Float64(), rand.Float64()}
		pt := kmeans.DataPointInit(a)
		points = append(points, *pt)
	}

	km := kmeans.New(k, points)

	results := km.Run(10)

	// fmt.Println(results)

	n := 50
	t := make(map[[2]int]int)
	for _, r := range results {
		x := int(r.Point[0] * float64(n))
		y := int(r.Point[1] * float64(n))
		t[[2]int{x, y}] = r.Assignment
	}

	for x := 0; x < n; x++ {
		for y := 0; y < n; y++ {
			v, ok := t[[2]int{x, y}]

			if ok {
				fmt.Printf("%d", v)
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Println()
	}

}
