package dice

import "fmt"

type PbtaResult struct {
	Strong      int
	Weak        int
	Miss        int
	StrongMatch int
	WeakMatch   int
	MissMatch   int
}

type PbtaParams struct {
	Modifier   int64
	Thresholds []int64
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
		"Strong Match: %f Weak Match: %f Miss Match: %f\n",
		float64(r.StrongMatch)/float64(rolls),
		float64(r.WeakMatch)/float64(rolls),
		float64(r.MissMatch)/float64(rolls))
}

func (results *PbtaResult) SetHitMiss(die1 int64, die2 int64, params *PbtaParams) PbtaResult {
	res := PbtaResult{
		Strong:      results.Strong,
		Weak:        results.Weak,
		Miss:        results.Miss,
		StrongMatch: results.StrongMatch,
		WeakMatch:   results.WeakMatch,
		MissMatch:   results.MissMatch,
	}
	if die1+die2+params.Modifier >= params.Thresholds[1] {
		res.Strong++
		if die1 == die2 {
			res.StrongMatch++
		}
	} else if die1+die2+params.Modifier >= params.Thresholds[0] {
		res.Weak++
		if die1 == die2 {
			res.WeakMatch++
		}
	} else {
		res.Miss++
		if die1 == die2 {
			res.MissMatch++
		}
	}

	return res
}
