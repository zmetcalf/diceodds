package main

import (
	"github.com/zmetcalf/diceodds/dice"
)

func main() {
	results := dice.PbtaResult{Params: dice.PbtaParams{
		Modifier: 3,
		Rolls:    1000000,
	}}

	for i := 0; i < results.Params.Rolls; i++ {
		challenge1 := dice.Roll(10, 1)
		challenge2 := dice.Roll(10, 1)
		action := dice.Roll(6, 1)

		if challenge1 < action+results.Params.Modifier && challenge2 < action+results.Params.Modifier {
			results.Strong++
			if challenge1 == challenge2 {
				results.StrongMatch++
			}
		} else if challenge1 < action+results.Params.Modifier || challenge2 < action+results.Params.Modifier {
			results.Weak++
		} else {
			results.Miss++
			if challenge1 == challenge2 {
				results.MissMatch++
			}
		}
	}

	results.Display()
	results.DisplayMatch()
}
