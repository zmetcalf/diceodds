package dice

import "testing"

func TestRoll(t *testing.T) {
	var result [7]int

	for i := 0; i < 1000; i++ {
		result[Roll(6, 1)]++
	}

	if result[0] != 0 {
		t.Errorf("Expected 0, got %d", result[0])
	}

	if result[1] == 0 {
		t.Errorf("Expected positive number, got %d", result[1])
	}

	if result[6] == 0 {
		t.Errorf("Expected positive number, got %d", result[1])
	}
}

func TestExplodingRoll(t *testing.T) {
	exploded := 0
	nuked := 0

	for i := 0; i < 10000; i++ {
		roll := ExplodingRoll(6, 1)
		if roll > 6 {
			exploded++
		}
		if roll > 24 {
			nuked++
		}
	}

	if exploded == 0 {
		t.Errorf("Expected positive number for exploded, got %d", exploded)
	}

	if nuked == 0 {
		t.Errorf("Expected positive number for nuked, got %d", nuked)
	}
}
