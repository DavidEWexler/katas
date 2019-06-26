package prime

import "testing"
import "reflect"

func TestPrimeFactorsOfOneIsEmpty(t *testing.T) {
	if got := Factors(1); len(got) != 0 {
		t.Errorf("Factors(1) = [], got %v", got)
	}
}

func TestFactors(t *testing.T) {
	var tests = []struct {
		input int
		want  []int
	}{
		{2, []int{2}},
		{3, []int{3}},
		{4, []int{2, 2}},
		{8, []int{2, 2, 2}},
		{9, []int{3, 3}},
		{2 * 3 * 5 * 7 * 7 * 13 * 17 * 29, []int{2, 3, 5, 7, 7, 13, 17, 29}},
		{1300639, []int{1300639}},
	}
	for _, test := range tests {
		if got := Factors(test.input); !reflect.DeepEqual(got, test.want) {
			t.Errorf("Factors(%d) = %v, got %v", test.input, test.want, got)
		}
	}
}

func BenchmarkPrimeFactors(b *testing.B) {
	num := 2 * 3 * 5 * 7 * 7 * 13 * 17 * 29
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		Factors(num)
	}
	b.StopTimer()
}

func BenchmarkLargePrime(b *testing.B) {
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		Factors(1300639)
	}
	b.StopTimer()
}
