package kmeans

import (
	"fmt"
	"math"
	"math/rand"
)

func MeanStddev(a []float64) (float64, float64) {
	sum := 0.0
	sumsqr := 0.0

	for _, v := range a {
		sum += v
		sumsqr += v * v
	}

	n := float64(len(a))

	mean := sum / n
	stddev := math.Sqrt(sumsqr/n - mean*mean)
	return mean, stddev
}

func Zscores(a []float64) []float64 {
	rv := make([]float64, len(a))
	mean, stddev := MeanStddev(a)
	for i, v := range a {
		rv[i] = (v - mean) / stddev
	}

	return rv
}

type DataPoint struct {
	original   []float64
	dimensions []float64
}

func DataPointInit(a []float64) *DataPoint {
	n := len(a)

	rv := DataPoint{
		original:   make([]float64, n),
		dimensions: make([]float64, n),
	}

	for i := 0; i < n; i++ {
		rv.original[i] = a[i]
		rv.dimensions[i] = a[i]
	}

	return &rv
}

func (d DataPoint) NumDimensions() int {
	return len(d.dimensions)
}

func (d1 DataPoint) Distance(d2 *DataPoint) float64 {
	if d1.NumDimensions() != d2.NumDimensions() {
		return -1.0
	}

	sum := 0.0
	for i := 0; i < d1.NumDimensions(); i++ {
		diff := d1.dimensions[i] - d2.dimensions[i]
		sum += diff * diff
	}
	return math.Sqrt(sum)
}

func (d1 DataPoint) Equal(d2 *DataPoint) bool {
	if d1.NumDimensions() != d2.NumDimensions() {
		return false
	}

	for i := 0; i < d1.NumDimensions(); i++ {
		if d1.dimensions[i] != d2.dimensions[i] {
			return false
		}
	}
	return true
}

func (d DataPoint) String() string {
	return fmt.Sprint(d.original)
}

type Cluster struct {
	point_idx []int
	centroid  DataPoint
}

type KMeans struct {
	k        int
	clusters []Cluster
	points   []DataPoint
}

func New(k int, data []DataPoint) *KMeans {
	var km KMeans
	km.k = k
	for _, pt := range data {
		km.points = append(km.points, pt)
	}
	km.clusters = make([]Cluster, k)
	return &km
}

func (km KMeans) NumPoints() int {
	return len(km.points)
}

func (km KMeans) NumDimensions() int {
	if km.NumPoints() == 0 {
		return -1
	}
	return km.points[0].NumDimensions()
}

func (km *KMeans) Normalize() {
	dim := km.NumDimensions()
	num_pts := km.NumPoints()
	orig := make([]float64, num_pts)

	for i := 0; i < dim; i++ {
		for j := 0; j < num_pts; j++ {
			orig[j] = km.points[j].original[i]
		}
		norm := Zscores(orig)
		for j := 0; j < num_pts; j++ {
			km.points[j].dimensions[i] = norm[i]
		}
	}
}

func (km *KMeans) ClearClusters() {
	for j := 0; j < km.k; j++ {
		km.clusters[j].point_idx = []int{}
	}
}

func (km *KMeans) AssignCluster() {
	km.ClearClusters()

	num_points := km.NumPoints()
	for j := 0; j < num_points; j++ {
		assignment := -1
		min_dist := 0.0

		for i := 0; i < km.k; i++ {
			dist := km.points[j].Distance(&km.clusters[i].centroid)
			if assignment == -1 || dist < min_dist {
				assignment = i
				min_dist = dist
			}
		}

		km.clusters[assignment].point_idx = append(km.clusters[assignment].point_idx, j)
	}
}

func (km *KMeans) GenerateCentroids() bool {
	dim := len(km.points[0].original)
	change := false
	for j := 0; j < km.k; j++ {
		num_pts := len(km.clusters[j].point_idx)

		a := make([]float64, dim)

		if num_pts == 0 {
			idx := rand.Intn(len(km.points))
			for i := 0; i < dim; i++ {
				a[i] = km.points[idx].dimensions[i]
			}
			change = true
		} else {
			for _, pt_idx := range km.clusters[j].point_idx {
				for i, v := range km.points[pt_idx].dimensions {
					a[i] += v
				}
			}

			for i := 0; i < dim; i++ {
				a[i] /= float64(len(km.clusters[j].point_idx))
				if a[i] != km.clusters[j].centroid.dimensions[i] {
					change = true
				}
			}
		}

		km.clusters[j].centroid.dimensions = a
	}
	return change
}

func (km *KMeans) Run(max_iters int) {
	km.ClearClusters()

	for iter := 0; iter < max_iters; iter++ {
		if !km.GenerateCentroids() {
			break
		}
		km.AssignCluster()
	}
}
