// Package syllogism implements the basic form of deductive argument, a
// syllogism. Syllogism is part of Aristotelian (or common sense) logic.
//
// Syllogism connects the subject and predicate terms of its conclusion through
// its middle term in its two propositions. The first of which states a general
// principle and the second brings a particular case under that principle. The
// conclusion then demonstrates the result of applying the general principle to
// the particular case.
package syllogism

import "fmt"

// Proposition states something about all or some of the the subjects. It can be
// affirmative or negative.
type Proposition struct {
	Some      bool
	Subject   string // term about which something is said in the predicate
	Negative  bool
	Predicate string // term that says something about the subject
}

// Conclude connects its two propositions (through a middle term). Major
// proposition states a general principle, minor proposition brings a particular
// case under that principle.
func Conclude(major, minor Proposition) (Proposition, error) {
	var conclusion Proposition

	switch {
	case major.Subject == minor.Predicate:
		conclusion.Subject = minor.Subject
		conclusion.Predicate = major.Predicate
	case major.Predicate == minor.Predicate:
		conclusion.Subject = minor.Subject
		conclusion.Predicate = major.Subject
	default:
		return conclusion, fmt.Errorf("can't conclude from %v and %v", major, minor)
	}

	conclusion.Some = major.Some || minor.Some
	conclusion.Negative = major.Negative || minor.Negative

	return conclusion, nil
}
