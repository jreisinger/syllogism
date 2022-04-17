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
					Type:      0,
					Subject:   Term{"rational", "animal"},
					Verb:      "be",
					Predicate: Term{"mortal"},
				},
				Proposition{
					Type:      1,
					Subject:   Term{"socrates"},
					Verb:      "be",
					Predicate: Term{"rational", "animal"},
				},
				Term{"rational", "animal"},
			},
			Conclusion{
				Term: Term{"socrates", "mortal"},
				Verb: "be",
			},
		},
		{
			Syllogism{
				Proposition{
					Type:      0,
					Subject:   Term{"rational", "animal"},
					Verb:      "be",
					Negative:  true,
					Predicate: Term{"mortal"},
					False:     true,
				},
				Proposition{
					Type:      1,
					Subject:   Term{"socrates"},
					Verb:      "be",
					Predicate: Term{"rational", "animal"},
				},
				Term{"rational", "animal"},
			},
			Conclusion{
				Term: Term{"socrates", "mortal"},
				Verb: "be",
			},
		},
		{
			Syllogism{
				Proposition{
					Type:      0,
					Subject:   Term{"rational", "animal"},
					Verb:      "be",
					Predicate: Term{"mortal"},
				},
				Proposition{
					Type:      1,
					Subject:   Term{"socrates"},
					Verb:      "be",
					Negative:  true,
					Predicate: Term{"rational", "animal"},
					False:     true,
				},
				Term{"rational", "animal"},
			},
			Conclusion{
				Term: Term{"socrates", "mortal"},
				Verb: "be",
			},
		},
		{
			Syllogism{
				Proposition{
					Type:      0,
					Subject:   Term{"rational", "animal"},
					Verb:      "be",
					Negative:  true,
					Predicate: Term{"mortal"},
					False:     true,
				},
				Proposition{
					Type:      1,
					Subject:   Term{"socrates"},
					Verb:      "be",
					Negative:  true,
					Predicate: Term{"rational", "animal"},
					False:     true,
				},
				Term{"rational", "animal"},
			},
			Conclusion{
				Term: Term{"socrates", "mortal"},
				Verb: "be",
			},
		},
		{
			Syllogism{
				Proposition{
					Type:      0,
					Subject:   Term{"bird"},
					Verb:      "lay",
					Predicate: Term{"eggs"},
				},
				Proposition{
					Type:      1,
					Subject:   Term{"elephant"},
					Verb:      "be",
					Negative:  true,
					Predicate: Term{"bird"},
				},
				Term{"bird"},
			},
			Conclusion{
				Term:     Term{"elephant", "eggs"},
				Verb:     "lay",
				Negative: true,
			},
		},
	}
	for _, tc := range testcases {
		conclusion, err := tc.syllogism.Conclude()
		if err != nil {
			t.Fatal(err)
		}
		if !eq(conclusion.Term, tc.conclusion.Term) {
			t.Fatalf("got conclusion terms: %v\nexpected conclusion terms: %v", conclusion.Term, tc.conclusion.Term)
		}
		if conclusion.Verb != tc.conclusion.Verb {
			t.Fatalf("got conclusion verb: %v\nexpected conclusion verb: %v", conclusion.Verb, tc.conclusion.Verb)
		}
		if conclusion.Negative != tc.conclusion.Negative {
			t.Fatalf("got conclusion negative: %v\nexpected conclusion negative: %v", conclusion.Negative, tc.conclusion.Negative)
		}
	}
}
