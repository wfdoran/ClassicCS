package main

import (
	"classic_sc/kmeans"
	"fmt"
	"math/rand"
	"time"
)

type Album struct {
	name   string
	year   int
	length float64
	tracks int
}

func main() {
	rand.Seed(time.Now().UnixNano())

	albums := []Album{
		Album{"Got to be there", 1972, 35.45, 10},
		Album{"Ben", 1972, 31.31, 10},
		Album{"Music & Me", 1973, 32.09, 10},
		Album{"Forever, Michael", 1975, 33.36, 10},
		Album{"Off the Wall", 1979, 42.28, 10},
		Album{"Thriller", 1982, 42.19, 9},
		Album{"Bad", 1987, 48.16, 10},
		Album{"Dangerous", 1991, 77.03, 14},
		Album{"HIStory", 1995, 148.58, 30},
		Album{"Invincible", 2001, 77.05, 16},
	}

	// fmt.Println(governors)

	var points []kmeans.DataPoint
	for _, x := range albums {
		a := []float64{x.length, float64(x.tracks)}
		pt := kmeans.DataPointInit(a)
		points = append(points, *pt)
	}

	k := 2
	km := kmeans.New(k, points)
	results := km.Run(100)

	// fmt.Println(results)

	min_x := results[0].Point[0]
	max_x := results[0].Point[0]
	min_y := results[0].Point[1]
	max_y := results[0].Point[1]

	for _, r := range results {
		if r.Point[0] < min_x {
			min_x = r.Point[0]
		}
		if r.Point[0] > max_x {
			max_x = r.Point[0]
		}
		if r.Point[1] < min_y {
			min_y = r.Point[1]
		}
		if r.Point[1] > max_y {
			max_y = r.Point[1]
		}
	}
	n := 50
	t := make(map[[2]int]int)
	for _, r := range results {
		x := int((r.Point[0] - min_x) / (max_x - min_x) * float64(n))
		y := int((r.Point[1] - min_y) / (max_y - min_y) * float64(n))
		t[[2]int{x, y}] = r.Assignment
	}

	// Strange looping to match picture in book (page 123)
	for x := n; x >= 0; x-- {
		for y := 0; y <= n; y++ {
			v, ok := t[[2]int{y, x}]

			if ok {
				fmt.Printf("%d", v)
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Println()
	}

}
