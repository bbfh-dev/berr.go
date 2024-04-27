package berr_test

import (
	"testing"

	"github.com/bbfh-dev/berr.go"
)

func divideErrOnly(a float64, b float64) error {
	_, err := divide(a, b)
	return err
}

func TestError(t *testing.T) {
	berr.Try(divideErrOnly(5, 0), func(err error) {
		if err == nil {
			t.Fatal()
		}
	})
}
