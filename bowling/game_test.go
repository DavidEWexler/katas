package bowling

import (
	"fmt"
	"testing"
)

func TestGame(t *testing.T) {
	var tests = []struct {
		input []int
		count int
		want  int
	}{
		{[]int{}, 20, 0},
		{[]int{1}, 19, 1},
		{[]int{8, 2, 8}, 17, 26},
		{[]int{10, 7, 2}, 17, 28},
	}
	for _, test := range tests {
		var g Game
		for _, pins := range test.input {
			g.Roll(pins)
		}
		for i := 0; i < test.count; i++ {
			g.Roll(0)
		}
		if got := g.Score(); got != test.want {
			t.Errorf("Rolling %v then all zero should score %d, got %d", test.input, test.want, got)
		}
	}
}

func ExampleGame() {
	/*
		1   | 2   | 3   | 4   | 5   | 6   | 7   | 8   | 9   | 10  |
		X   | X   | 9/  | 72  | X   | 63  | -/  | 21  | --  | X23 |
		29  | 49  | 66  | 75  | 94  | 103 | 115 | 118 | 118 | 133 |
	*/
	var g Game
	g.Roll(10)
	g.Roll(10)
	g.Roll(9)
	g.Roll(1)
	g.Roll(7)
	g.Roll(2)
	g.Roll(10)
	g.Roll(6)
	g.Roll(3)
	g.Roll(0)
	g.Roll(10)
	g.Roll(2)
	g.Roll(1)
	g.Roll(0)
	g.Roll(0)
	g.Roll(10)
	g.Roll(2)
	g.Roll(3)
	fmt.Println(g.Score())
	// output: 133
}
