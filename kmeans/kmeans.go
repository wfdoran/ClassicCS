package kmeans

import (
	"fmt"
	"math"
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

func (d DataPoint) NumDimenstions() int {
	return len(d.dimensions)
}

func (d1 DataPoint) Distance(d2 *DataPoint) float64 {
	if len(d1.dimensions) != len(d2.dimensions) {
		return -1.0
	}

	sum := 0.0
	for i := 0; i < len(d1.dimensions); i++ {
		diff := d1.dimensions[i] - d2.dimensions[i]
		sum += diff * diff
	}
	return math.Sqrt(sum)
}

func (d1 DataPoint) Equal(d2 *DataPoint) bool {
	if len(d1.dimensions) != len(d2.dimensions) {
		return false
	}

	for i := 0; i < len(d1.dimensions); i++ {
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

func (km *KMeans) Normalize() {
	dim := len(km.points[0].original)
	num_pts := len(km.points)
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

func (km *KMeans) AssignCluster() {
	for i := 0; i < km.k; i++ {
		km.clusters[i].point_idx = []int{}
	}

	num_points := len(km.points)
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

func (km *KMeans) GenerateCentroids() {
	dim := len(km.points[0].original)
	for j := 0; j < km.k; j++ {
		a := make([]float64, dim)

		for _, pt_idx := range km.clusters[j].point_idx {
			for i, v := range km.points[pt_idx].dimensions {
				a[i] += v
			}
		}

		for i := 0; i < dim; i++ {
			a[i] /= float64(len(km.clusters[i].point_idx))
		}

		km.clusters[j].centroid.dimensions = a
	}

}
