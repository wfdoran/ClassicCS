package main

import (
	"classic_sc/kmeans"
	"fmt"
	"math/rand"
	"time"
)

type Governor struct {
	longitude float64
	age       int
	state     string
}

func main() {
	rand.Seed(time.Now().UnixNano())

	governors := []Governor{Governor{-86.79113, 72, "Alabama"},
		Governor{-152.404419, 66, "Alaska"},
		Governor{-111.431221, 53, "Arizona"},
		Governor{-92.373123, 66, "Arkansas"},
		Governor{-119.681564, 79, "California"},
		Governor{-105.311104, 65, "Colorado"},
		Governor{-72.755371, 61, "Connecticut"},
		Governor{-75.507141, 61, "Delaware"},
		Governor{-81.686783, 64, "Florida"},
		Governor{-83.643074, 74, "Georgia"},
		Governor{-157.498337, 60, "Hawaii"},
		Governor{-114.478828, 75, "Idaho"},
		Governor{-88.986137, 60, "Illinois"},
		Governor{-86.258278, 49, "Indiana"},
		Governor{-93.210526, 57, "Iowa"},
		Governor{-96.726486, 60, "Kansas"},
		Governor{-84.670067, 50, "Kentucky"},
		Governor{-91.867805, 50, "Louisiana"},
		Governor{-69.381927, 68, "Maine"},
		Governor{-76.802101, 61, "Maryland"},
		Governor{-71.530106, 60, "Massachusetts"},
		Governor{-84.536095, 58, "Michigan"},
		Governor{-93.900192, 70, "Minnesota"},
		Governor{-89.678696, 62, "Mississippi"},
		Governor{-92.288368, 43, "Missouri"},
		Governor{-110.454353, 51, "Montana"},
		Governor{-98.268082, 52, "Nebraska"},
		Governor{-117.055374, 53, "Nevada"},
		Governor{-71.563896, 42, "New Hampshire"},
		Governor{-74.521011, 54, "New Jersey"},
		Governor{-106.248482, 57, "New Mexico"},
		Governor{-74.948051, 59, "New York"},
		Governor{-79.806419, 60, "North Carolina"},
		Governor{-99.784012, 60, "North Dakota"},
		Governor{-82.764915, 65, "Ohio"},
		Governor{-96.928917, 62, "Oklahoma"},
		Governor{-122.070938, 56, "Oregon"},
		Governor{-77.209755, 68, "Pennsylvania"},
		Governor{-71.51178, 46, "Rhode Island"},
		Governor{-80.945007, 70, "South Carolina"},
		Governor{-99.438828, 64, "South Dakota"},
		Governor{-86.692345, 58, "Tennessee"},
		Governor{-97.563461, 59, "Texas"},
		Governor{-111.862434, 70, "Utah"},
		Governor{-72.710686, 58, "Vermont"},
		Governor{-78.169968, 60, "Virginia"},
		Governor{-121.490494, 66, "Washington"},
		Governor{-80.954453, 66, "West Virginia"},
		Governor{-89.616508, 49, "Wisconsin"},
		Governor{-107.30249, 55, "Wyoming"}}

	// fmt.Println(governors)

	var points []kmeans.DataPoint
	for _, g := range governors {
		a := []float64{g.longitude, float64(g.age)}
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
	for x := n - 1; x >= 0; x-- {
		for y := 0; y < n; y++ {
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
