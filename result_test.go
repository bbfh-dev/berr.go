package berr_test

import (
	"errors"
	"testing"

	"github.com/bbfh-dev/berr.go"
)

func divide(a float64, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("invalid operation: division by zero")
	}

	return a / b, nil
}

func TestValue(t *testing.T) {
	if berr.Get(divide(5, 0)).ValueOr(0) != 0 {
		t.Fatal()
	}

	if berr.Get(divide(10, 5)).ValueOrFatal() != 2 {
		t.Fatal()
	}

	if berr.Get(divide(10, 5)).ValueOrPanic() != 2 {
		t.Fatal()
	}
}
