package regress_test

import (
	"testing"

	regress "github.com/jevankovich/go-regress"
)

func TestPolyval(t *testing.T) {
	if regress.Polyval(regress.Polynomial{1}, 0) != 1 {
		t.Fail()
	}

	if regress.Polyval(regress.Polynomial{1}, 1) != 1 {
		t.Fail()
	}

	if regress.Polyval(regress.Polynomial{1, 2}, 1) != 3 {
		t.Fail()
	}
}
