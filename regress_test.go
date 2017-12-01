package regress_test

import (
	"math"
	"testing"

	regress "github.com/jevankovich/go-regress"
)

func similar(x, y []float64, eps float64) bool {
	if len(x) != len(y) {
		return false
	}

	for i := range x {
		if math.Abs(x[i]-y[i]) > eps {
			return false
		}
	}

	return true
}

func TestRegress(t *testing.T) {
	x := []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	y := []float64{1, 6, 17, 34, 57, 86, 121, 162, 209, 262, 321}

	p := regress.Regress(x, y, 2)

	if !similar(p, []float64{1, 2, 3}, 0.000001) {
		t.Error("Got:", p, "Expected:", []float64{1, 2, 3})
	}
}
