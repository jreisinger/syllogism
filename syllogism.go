// Package syllogism implements syllogism, the basic form of deductive argument.
// Syllogism is a part of Aristotelian (or common sense) logic.
package syllogism

import "fmt"

// Proposition is a declarative sentence.
type Proposition struct {
	Subject   Term
	Predicate Term
}

// Term is the subject or predicate of a proposition expressing a concept.
type Term []string

// Conclude connects the subject and predicate terms in its two propositions.
// Major proposition states a general principle, minor proposition brings a
// particalur case under that priciple.
func Conclude(major, minor Proposition) (Proposition, error) {
	var conclusion Proposition
	switch {
	case eq(major.Subject, minor.Predicate):
		conclusion.Subject = minor.Subject
		conclusion.Predicate = major.Predicate
	default:
		return conclusion, fmt.Errorf("can't conclude from %v and %v", major, minor)
	}
	return conclusion, nil
}

func eq(a, b Term) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
