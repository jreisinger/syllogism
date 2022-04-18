package syllogism

import (
	"reflect"
	"testing"
)

func TestSyllogism(t *testing.T) {
	testcases := []struct {
		major      Proposition
		minor      Proposition
		conclusion Proposition
	}{
		// Classical syllogism example.
		{
			Proposition{false, "rational animals", false, "mortal"},
			Proposition{false, "socrates", false, "rational animals"},
			Proposition{false, "socrates", false, "mortal"},
		},

		// http://ceur-ws.org/Vol-1412/3o.pdf
		{
			Proposition{false, "police dogs", true, "vicious"},
			Proposition{true, "highly trained dogs", false, "vicious"},
			Proposition{true, "highly trained dogs", true, "police dogs"},
		},
		{
			Proposition{false, "nutritional things", false, "expensive"},
			Proposition{true, "vitamin tablets", true, "expensive"},
			Proposition{true, "vitamin tablets", true, "nutritional things"},
		},
	}
	for _, tc := range testcases {
		conclusion, err := Conclude(tc.major, tc.minor)
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(conclusion, tc.conclusion) {
			t.Fatalf("\nGot: %+v\nExp: %+v", conclusion, tc.conclusion)
		}
	}
}
