package syllogism

import (
	"testing"
)

func TestSyllogism(t *testing.T) {
	testcases := []struct {
		syllogism  Syllogism
		conclusion Conclusion
	}{
		{
			Syllogism{
				Proposition{
					Subject:   []Term{"rational", "animal"},
					Verb:      "be",
					Predicate: []Term{"mortal"},
					Type:      0,
				},
				Proposition{
					Subject:   []Term{"socrates"},
					Verb:      "be",
					Predicate: []Term{"rational", "animal"},
					Type:      1,
				},
				[]Term{"rational", "animal"},
			},
			Conclusion{
				Terms: []Term{"socrates", "mortal"},
				Verb:  "be",
			},
		},
		{
			Syllogism{
				Proposition{
					Subject:   []Term{"rational", "animal"},
					Verb:      "be",
					Negative:  true,
					Predicate: []Term{"mortal"},
					Type:      0,
					False:     true,
				},
				Proposition{
					Subject:   []Term{"socrates"},
					Verb:      "be",
					Predicate: []Term{"rational", "animal"},
					Type:      1,
				},
				[]Term{"rational", "animal"},
			},
			Conclusion{
				Terms: []Term{"socrates", "mortal"},
				Verb:  "be",
			},
		},
		{
			Syllogism{
				Proposition{
					Subject:   []Term{"rational", "animal"},
					Verb:      "be",
					Predicate: []Term{"mortal"},
					Type:      0,
				},
				Proposition{
					Subject:   []Term{"socrates"},
					Verb:      "be",
					Negative:  true,
					Predicate: []Term{"rational", "animal"},
					Type:      1,
					False:     true,
				},
				[]Term{"rational", "animal"},
			},
			Conclusion{
				Terms: []Term{"socrates", "mortal"},
				Verb:  "be",
			},
		},
		{
			Syllogism{
				Proposition{
					Subject:   []Term{"rational", "animal"},
					Verb:      "be",
					Negative:  true,
					Predicate: []Term{"mortal"},
					Type:      0,
					False:     true,
				},
				Proposition{
					Subject:   []Term{"socrates"},
					Verb:      "be",
					Negative:  true,
					Predicate: []Term{"rational", "animal"},
					Type:      1,
					False:     true,
				},
				[]Term{"rational", "animal"},
			},
			Conclusion{
				Terms: []Term{"socrates", "mortal"},
				Verb:  "be",
			},
		},
	}
	for _, tc := range testcases {
		conclusion, err := tc.syllogism.Conclude()
		if err != nil {
			t.Fatal(err)
		}
		if !eq(conclusion.Terms, tc.conclusion.Terms) {
			t.Fatalf("got conclusion terms: %v\nexpected conclusion terms: %v", conclusion.Terms, tc.conclusion.Terms)
		}
		if conclusion.Verb != tc.conclusion.Verb {
			t.Fatalf("got conclusion verb: %v\nexpected conclusion verb: %v", conclusion.Verb, tc.conclusion.Verb)
		}
		if conclusion.Negative != tc.conclusion.Negative {
			t.Fatalf("got conclusion negative: %v\nexpected conclusion negative: %v", conclusion.Negative, tc.conclusion.Negative)
		}
	}
}
