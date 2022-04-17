// Package syllogism implements syllogism, the basic form of deductive argument.
// Syllogism is a part of Aristotelian (or common sense) logic.
package syllogism

import (
	"fmt"
)

const (
	PropositionMajor PropositionType = iota // general principle
	PropositionMinor                        // particular case
)

type PropositionType int32

// Syllogism connects the subject and predicate terms in its two propositions.
type Syllogism struct {
	PropositionMajor Proposition
	PropositionMinor Proposition
	Middle           Term
}

// Proposition is a declarative sentence, affirmative or negative, that can be
// either true or false.
type Proposition struct {
	Type      PropositionType
	Subject   Term
	Verb      string // e.g., be
	Negative  bool   // e.g., not be
	Predicate Term
	False     bool
}

// Conclusion is the application of a major proposition to a minor proposition.
type Conclusion struct {
	Term     Term
	Verb     string
	Negative bool
}

// Term is the subject or predicate of a proposition expressing a concept.
type Term []string

// Conclude generates a conclusion.
func (s Syllogism) Conclude() (Conclusion, error) {
	if !eq(s.PropositionMajor.Subject, s.PropositionMinor.Predicate) {
		return Conclusion{}, fmt.Errorf("Subject of major proposition (%v) doesn't match predicate of minor proposition (%v)", s.PropositionMajor.Subject, s.PropositionMinor.Predicate)
	}
	if s.PropositionMajor.Verb != s.PropositionMinor.Verb {
		return Conclusion{}, fmt.Errorf("verbs (%v, %v) of propositions don't match", s.PropositionMajor.Verb, s.PropositionMinor.Verb)
	}

	var conclusion Conclusion
	conclusion.Verb = s.PropositionMajor.Verb

	switch {
	case !s.PropositionMajor.Negative && !s.PropositionMinor.Negative:
		conclusion.Negative = false
	case s.PropositionMajor.Negative && !s.PropositionMinor.Negative:
		conclusion.Negative = true
	case !s.PropositionMajor.Negative && s.PropositionMinor.Negative:
		conclusion.Negative = true
	case s.PropositionMajor.Negative && s.PropositionMinor.Negative:
		conclusion.Negative = true
	}

	if s.PropositionMajor.False || s.PropositionMinor.False {
		conclusion.Negative = !conclusion.Negative
	}

	var t Term
	t = append(t, s.PropositionMinor.Subject...)
	t = append(t, s.PropositionMajor.Predicate...)
	conclusion.Term = append(conclusion.Term, t...)
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
