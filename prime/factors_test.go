package prime

import "testing"

func TestPrimeFactorsOfOneIsEmpty(t *testing.T) {
	if got := Factors(1); len(got) != 0 {
		t.Errorf("Factors(1) = [], got %v", got)
	}
}
