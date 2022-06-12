package syllogism

import (
	"fmt"
	"log"
	"reflect"
	"testing"
)

// Classical syllogism example.
func Example() {
	major := Proposition{
		Subject:   "rational animals",
		Predicate: "mortal",
	}
	minor := Proposition{
		Subject:   "Socrates",
		Predicate: "rational animals",
	}
	conclusion, err := Conclude(major, minor)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s are %s\n", major.Subject, major.Predicate)
	fmt.Printf("%s is one of the %s\n", minor.Subject, minor.Predicate)
	fmt.Printf("therefore %s is %s\n", conclusion.Subject, conclusion.Predicate)
	// Output: rational animals are mortal
	// Socrates is one of the rational animals
	// therefore Socrates is mortal
}

func TestSyllogism(t *testing.T) {
	testcases := []struct {
		major      Proposition
		minor      Proposition
		conclusion Proposition
	}{
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

		// https://en.wikipedia.org/wiki/Syllogism#Examples
		{ // Barbara (AAA-1)
			Proposition{false, "men", false, "mortal"},
			Proposition{false, "Greeks", false, "men"},
			Proposition{false, "Greeks", false, "mortal"},
		},
		{ // Celarent (EAE-1)
			Proposition{false, "reptiles", true, "fur"},
			Proposition{false, "snakes", false, "reptiles"},
			Proposition{false, "snakes", true, "fur"},
		},
		{ // Darii (AII-1)
			Proposition{false, "rabbits", false, "fur"},
			Proposition{true, "pets", false, "rabbits"},
			Proposition{true, "pets", false, "fur"},
		},
		{ // Ferio (EIO-1)
			Proposition{false, "homework", true, "fun"},
			Proposition{true, "reading", false, "homework"},
			Proposition{true, "reading", true, "fun"},
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
