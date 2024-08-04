package main

import (
	"github.com/zmetcalf/diceodds/dice"
)

func main() {
	rolls := 1000000
	for _, die := range []int{4, 6, 8, 10, 12, 20} {
		result := dice.CheckResult{}

		for i := 0; i < rolls; i++ {
			dieRoll := dice.ExplodingRoll(die, 1)
			aceRoll := dice.ExplodingRoll(6, 1)

			if dieRoll >= 12 || aceRoll >= 12 {
				result.MultiRaise++
			} else if dieRoll >= 8 || aceRoll >= 8 {
				result.Raised++
			} else if dieRoll >= 4 || aceRoll >= 4 {
				result.Hit++
			} else {
				result.Miss++
			}
		}

		result.Display(die, rolls)
	}
}
