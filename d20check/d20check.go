package main

import (
	"github.com/zmetcalf/diceodds/dice"
)

func main() {
	rolls := 1000000
	result := dice.CheckResult{}
	die := 20

	for i := 0; i < rolls; i++ {
		roll := dice.Roll(die, 1)

		if roll >= 20 {
			result.MultiRaise++
		} else if roll >= 16 {
			result.Raised++
		} else if roll >= 6 {
			result.Hit++
		} else {
			result.Miss++
		}
	}

	result.Display(die, rolls)
}
