package dice

import "fmt"

func (results *SWDamageResult) SetHitMiss(die_total int64, modifier int64, toughness int) SWDamageResult {
	res := SWDamageResult{Hit: results.Hit, Miss: results.Miss, Toughness: toughness}
	if die_total+modifier >= int64(toughness) {
		res.Hit++
	} else {
		res.Miss++
	}

	return res
}

func (r SWDamageResult) Display(die int, rolls int) {
	fmt.Printf(
		"Die: %d Toughness: %d Hit: %f Miss: %f\n",
		die,
		r.Toughness,
		float64(r.Hit)/float64(rolls),
		float64(r.Miss)/float64(rolls))
}
