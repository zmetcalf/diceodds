package main

import (
	"github.com/zmetcalf/diceodds/dice"
)

func main() {
	results := dice.PbtaResult{Strong: 0, Weak: 0, Miss: 0}
	rolls := 1000000
	var modifier int64 = 1

	for i := 0; i < rolls; i++ {
		die1 := dice.Roll(6, 1)
		die2 := dice.Roll(6, 1)

		if die1+die2+modifier >= 10 {
			results.Strong++
		} else if die1+die2+modifier >= 7 {
			results.Weak++
		} else {
			results.Miss++
		}
	}

	results.Display(rolls)
}
