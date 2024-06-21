package main

import (
	"github.com/zmetcalf/diceodds/dice"
)

func main() {
	rolls := 1000000
	results := dice.PbtaResult{
		Params: dice.PbtaParams{
			Modifier:   3,
			Thresholds: []int64{8, 11},
			Rolls:      rolls,
		},
	}

	for i := 0; i < rolls; i++ {
		die1 := dice.Roll(6, 1)
		die2 := dice.Roll(6, 1)

		results = results.SetHitMiss(die1, die2)
	}

	results.Display()
	results.DisplayMatch()
}
