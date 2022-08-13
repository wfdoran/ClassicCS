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
