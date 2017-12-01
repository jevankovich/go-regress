package regress

import "gonum.org/v1/gonum/mat"

func Regress(x, y []float64, degree int) []float64 {
	xx := vandermonde(x, degree)
	yy := mat.NewVecDense(len(y), y)
	c := mat.NewVecDense(degree+1, nil)

	qr := new(mat.QR)
	qr.Factorize(xx)

	qr.SolveVec(c, false, yy)

	return c.RawVector().Data
}

func vandermonde(x []float64, degree int) *mat.Dense {
	mat := mat.NewDense(len(x), degree+1, nil)
	for i, xx := range x {
		for j, p := 0, 1.0; j < degree+1; j, p = j+1, p*xx {
			mat.Set(i, j, p)
		}
	}

	return mat
}
