package main

import (
	"github.com/zmetcalf/diceodds/dice"
)

func main() {
	results := dice.PbtaResult{Strong: 0, Weak: 0, Miss: 0, StrongMatch: 0, MissMatch: 0}
	rolls := 1000000
	var modifier int64 = 2

	for i := 0; i < rolls; i++ {
		challenge1 := dice.Roll(10, 1)
		challenge2 := dice.Roll(10, 1)
		action := dice.Roll(6, 1)

		if challenge1 < action+modifier && challenge2 < action+modifier {
			results.Strong++
			if challenge1 == challenge2 {
				results.StrongMatch++
			}
		} else if challenge1 < action+modifier || challenge2 < action+modifier {
			results.Weak++
		} else {
			results.Miss++
			if challenge1 == challenge2 {
				results.MissMatch++
			}
		}
	}

	results.Display(rolls)
	results.DisplayMatch(rolls)
}
