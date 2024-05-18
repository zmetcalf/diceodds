package main

import (
	"github.com/zmetcalf/diceodds/dice"
)

func main() {
	results := dice.PbtaResult{Strong: 0, Weak: 0, Miss: 0}
	rolls := 1000000
	params := dice.PbtaParams{
		Modifier:   0,
		Thresholds: []int64{6, 9},
	}

	for i := 0; i < rolls; i++ {
		die1 := dice.Roll(6, 1)
		die2 := dice.Roll(6, 1)

		results = results.SetHitMiss(die1, die2, &params)
	}

	results.Display(rolls)
}
