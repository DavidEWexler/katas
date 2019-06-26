package lcm

import "testing"

func TestLeastCommonMultiple(t *testing.T) {
	var tests = []struct {
		n1, n2 int
		want   int
	}{
		{5, 7, 35},
		{2, 4, 4},
		{4, 2, 4},
		{4, 6, 12},
		{12, 20, 60},
	}
	for _, test := range tests {
		if got := LeastCommonMultiple(test.n1, test.n2); got != test.want {
			t.Errorf("LeastCummonMultiple(%d, %d) = %d, got %d",
				test.n1, test.n2, test.want, got)
		}
	}
}
