package bowling

import (
	"testing"
)

func TestGutterGameScoresZero(t *testing.T) {
	var game Game
	for i := 0; i < 20; i++ {
		game.Roll(0)
	}
	if got := game.Score(); got != 0 {
		t.Errorf("Gutter game should score zero, got %d", got)
	}
}

func ExampleGame() {
	/*
		1   | 2   | 3   | 4   | 5   | 6   | 7   | 8   | 9   | 10  |
		X   | X   | 9/  | 72  | X   | 63  | -/  | 21  | --  | X23 |
		29  | 49  | 66  | 75  | 94  | 103 | 115 | 118 | 118 | 133 |
	*/
	// var g Game
	// g.Roll(10)
	// g.Roll(10)
	// g.Roll(9)
	// g.Roll(1)
	// g.Roll(7)
	// g.Roll(2)
	// g.Roll(10)
	// g.Roll(6)
	// g.Roll(3)
	// g.Roll(0)
	// g.Roll(10)
	// g.Roll(2)
	// g.Roll(1)
	// g.Roll(0)
	// g.Roll(0)
	// g.Roll(10)
	// g.Roll(2)
	// g.Roll(3)
	// fmt.Println(g.Score())
	// output: 133
}
