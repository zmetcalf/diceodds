package main

import (
	"github.com/zmetcalf/diceodds/dice"
)

func main() {
	toughness := [22]dice.SWDamageResult{}

	rolls := 1000000
	var modifier int64 = 0

	for i := 0; i < rolls; i++ {
		die_total := dice.ExplodingRoll(6, 1) + dice.ExplodingRoll(6, 1)

		if dice.RaisedAttack(6) {
			die_total += dice.ExplodingRoll(6, 1)
		}

		for k, v := range toughness {
			toughness[k] = v.SetHitMiss(die_total, modifier, k)
		}
	}

	for _, v := range toughness {
		v.Display(6, rolls)
	}
}
