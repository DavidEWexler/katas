package avgpx

import (
	"fmt"
	"testing"
)

func TestAveragePrice(t *testing.T) {
	var tests = []struct {
		input []Trade
		want  string
	}{
		{[]Trade{Trade{1, 10.0}}, "10.0000"},
	}
	for _, test := range tests {
		if got := fmt.Sprintf("%.4f", AveragePrice(test.input...)); got != test.want {
			t.Errorf("AveragePrice(%v) = %s, got %s", test.input, test.want, got)
		}
	}
}
