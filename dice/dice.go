package dice

import (
	"crypto/rand"
	"math/big"
)

type CheckResult struct {
	Hit        int
	Raised     int
	MultiRaise int
	Miss       int
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

func RaisedAttack(die int) bool {
	action_die := ExplodingRoll(die, 1)
	ace_die := ExplodingRoll(6, 1)

	return action_die >= 8 || ace_die >= 8
}
