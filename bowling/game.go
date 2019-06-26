package bowling

type Game struct {
	rolls [21]int
	at    int
}

func (g *Game) Roll(pins int) {
	g.rolls[g.at] = pins
	g.at++
}

func (g *Game) Score() (score int) {
	at := 0
	for frame := 0; frame < 10; frame++ {
		if g.isStrike(at) {
			score += 10 + g.scoreOfNextTwoBalls(at+1)
			at++
		} else if g.isSpare(at) {
			// Spare
			score += 10 + g.rolls[at+2]
			at += 2
		} else {
			score += g.scoreOfNextTwoBalls(at)
			at += 2
		}
	}
	return score
}

func (g *Game) scoreOfNextTwoBalls(at int) int {
	return g.rolls[at] + g.rolls[at+1]
}

func (g *Game) isStrike(at int) bool {
	return g.rolls[at] == 10
}

func (g *Game) isSpare(at int) bool {
	return g.scoreOfNextTwoBalls(at) == 10
}
