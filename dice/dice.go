package dice

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

type PbtaResult struct {
	Strong      int
	Weak        int
	Miss        int
	StrongMatch int
	MissMatch   int
}

type SWDamageResult struct {
	Hit       int
	Miss      int
	Toughness int
}

func Roll(sides int, starting int64) int64 {
	die, err := rand.Int(rand.Reader, big.NewInt(int64(sides)))

	if err != nil {
		panic(err)
	}

	return die.Int64() + starting
}

func ExplodingRoll(sides int, starting int64) int64 {
	die := Roll(sides, starting)
	if die == int64(sides) {
		die += ExplodingRoll(sides, starting)
	}

	return die
}

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

func (r PbtaResult) Display(rolls int) {
	fmt.Printf(
		"Strong: %f Weak: %f Miss: %f\n",
		float64(r.Strong)/float64(rolls),
		float64(r.Weak)/float64(rolls),
		float64(r.Miss)/float64(rolls))
}

func (r PbtaResult) DisplayMatch(rolls int) {
	fmt.Printf(
		"Strong Match: %f Miss Match: %f\n",
		float64(r.StrongMatch)/float64(rolls),
		float64(r.MissMatch)/float64(rolls))
}

func RaisedAttack(die int) bool {
	action_die := ExplodingRoll(die, 1)
	ace_die := ExplodingRoll(6, 1)

	return action_die >= 8 || ace_die >= 8
}
