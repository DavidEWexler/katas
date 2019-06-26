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
		{[]Trade{}, "0.0000"},
		{[]Trade{Trade{1, 10.0}, Trade{2, 20.0}, Trade{4, 5.0}}, "10.0000"},
		{[]Trade{Trade{1, 1.0}, Trade{2, 2.0}}, "1.6667"},
	}
	for _, test := range tests {
		if got := fmt.Sprintf("%.4f", AveragePrice(test.input...)); got != test.want {
			t.Errorf("AveragePrice(%v) = %s, got %s", test.input, test.want, got)
		}
	}
}
