package editd

import (
	"testing"
)

type testCase struct {
	s1       string
	s2       string
	distance int
}

func TestEditDistance(t *testing.T) {
	tests := []testCase{
		testCase{"hello", "hello", 0},
		testCase{"hello", "hell", 1},
		testCase{"hello", "helpo", 1},
		testCase{"llo", "hello", 2},
		testCase{"cubby", "bear", 5},
		testCase{"", "", 0},
		testCase{"foo", "bar", 3}}		

	for _, test := range tests {
		distance := EditDistance(test.s1, test.s2)

		if distance != test.distance {
			t.Error()
		}
	}
}
