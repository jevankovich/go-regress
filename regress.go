package regress

type Polynomial []float64

func Polyval(p Polynomial, x float64) float64 {
	pslice := []float64(p)

	xx := 1.0
	y := 0.0
	for _, a := range pslice {
		y += a * xx
		xx *= x
	}

	return y
}
