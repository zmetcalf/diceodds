package main

import (
	"github.com/zmetcalf/diceodds/dice"
)

func main() {
	rolls := 1000000
	var modifier int64 = 0

	for _, die := range []int{4, 6, 8, 10, 12, 20} {
		toughness := [22]dice.SWDamageResult{}

		for i := 0; i < rolls; i++ {
			die_total := dice.ExplodingRoll(die, 1) + dice.ExplodingRoll(die, 1)

			if dice.RaisedAttack(die) {
				die_total += dice.ExplodingRoll(6, 1)
			}

			for k, v := range toughness {
				toughness[k] = dice.SetHitMiss(&v, die_total, modifier, k)
			}
		}

		for _, v := range toughness {
			v.Display(die, rolls)
		}
	}
}
