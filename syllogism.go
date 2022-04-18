// Package syllogism implements syllogism, the basic form of deductive argument.
// Syllogism is a part of Aristotelian (or common sense) logic.
package syllogism

import "fmt"

// Proposition states something about all or some of the the subjects. It can be
// affirmative or negative.
type Proposition struct {
	Some      bool
	Subject   string
	Negative  bool
	Predicate string
}

// Conclude connects its two propositions through a middle term. Major
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
