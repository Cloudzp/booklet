// The gobinarytest package is used in testing serialization and
// deserialization. It is particularly useful when some sub-sequences of bytes
// can acceptably be written in any order (e.g if they are generated from
// representations that do not provide order guarantees like maps).
package gobinarytest

import (
	"bytes"
	"fmt"
	"strings"
)

type MatchError struct {
	FailedMatcher Matcher
	NextBytes     []byte
	Context       []string
}

func (err *MatchError) Error() string {
	return fmt.Sprintf(
		"failed to match <%s> %v on remaining bytes: % x",
		strings.Join(err.Context, " within "),
		err.FailedMatcher, err.NextBytes)
}

func (err *MatchError) AddContext(ctx string) {
	err.Context = append(err.Context, ctx)
}

type TrailingBytesError struct {
	FailedMatcher Matcher
	TrailingBytes []byte
}

func (err *TrailingBytesError) Error() string {
	return fmt.Sprintf("%d trailing bytes when matching %v: % x",
		len(err.TrailingBytes), err.FailedMatcher, err.TrailingBytes)
}

// Matcher is the interface that tests if a sequence of bytes matches
// expectations.
type Matcher interface {
	// Match tests if the given sequence of bytes matches. It returns
	// matches=true if it matches correctly, and returns the number of bytes
	// matched as n.
	Match(b []byte) (n int, err error)

	// Write writes a sequence of bytes to the buffer that is an acceptable match
	// for the matcher.
	Write(writer *bytes.Buffer)

	String() string
}

// Named is a wrapper for a Matcher that simply provides a name annotation on
// its String method and add its name to any MatchErrors.
type Named struct {
	Name string
	Matcher
}

func (bm Named) Match(b []byte) (n int, err error) {
	n, err = bm.Matcher.Match(b)
	if err != nil {
		if matchErr, ok := err.(*MatchError); ok {
			matchErr.AddContext(bm.Name)
		}
	}
	return
}

func (bm Named) String() string {
	return fmt.Sprintf("(%s) %s", bm.Name, bm.Matcher)
}

// Literal matches a literal sequence of bytes.
type Literal []byte

func LiteralString(s string) Literal {
	return Literal(s)
}

func (bm Literal) Match(b []byte) (n int, err error) {
	if len(b) < len(bm) {
		return 0, &MatchError{bm, b, nil}
	}

	for i, v := range bm {
		if v != b[i] {
			return 0, &MatchError{bm, b, nil}
		}
	}

	return len(bm), nil
}

func (bm Literal) Write(writer *bytes.Buffer) {
	writer.Write(bm)
}

func (bm Literal) String() string {
	return fmt.Sprintf("Literal{% x}", []byte(bm))
}

// InOrder matches a sequence of Matchers that must match in
// the order given.
type InOrder []Matcher

func (bm InOrder) Match(b []byte) (n int, err error) {
	var consumed int

	for _, matcher := range bm {
		remainder := b[n:]

		if consumed, err = matcher.Match(remainder); err != nil {
			return 0, err
		}

		n += consumed
	}

	return n, nil
}

func (bm InOrder) Write(writer *bytes.Buffer) {
	for _, matcher := range bm {
		matcher.Write(writer)
	}
}

func (bm InOrder) String() string {
	parts := make([]string, len(bm))
	for i, matcher := range bm {
		parts[i] = matcher.String()
	}
	s := strings.Join(parts, ", ")
	return fmt.Sprintf("InOrder{%s}", s)
}

// AnyOrder matches a set of Matchers, but they do not have to
// match in any particular order.
type AnyOrder []Matcher

func (bm AnyOrder) Match(b []byte) (n int, err error) {
	toMatch := make([]Matcher, len(bm))
	copy(toMatch, bm)
	var consumed int

	for {
		if len(toMatch) == 0 {
			break
		}

		remainder := b[n:]

		foundMatch := false
		for i, matcher := range toMatch {
			if consumed, err = matcher.Match(remainder); err == nil {
				foundMatch = true
				n += consumed

				// Remove matcher from toMatch.
				if len(toMatch) > 1 {
					toMatch[i] = toMatch[len(toMatch)-1]
				}
				toMatch = toMatch[:len(toMatch)-1]

				break
			}
		}

		if !foundMatch {
			return 0, &MatchError{AnyOrder(toMatch), remainder, nil}
		}
	}

	return n, nil
}

func (bm AnyOrder) Write(writer *bytes.Buffer) {
	for _, matcher := range bm {
		matcher.Write(writer)
	}
}

func (bm AnyOrder) String() string {
	parts := make([]string, len(bm))
	for i, matcher := range bm {
		parts[i] = matcher.String()
	}
	s := strings.Join(parts, ", ")
	return fmt.Sprintf("&AnyOrder{%s}", s)
}

func Matches(matcher Matcher, b []byte) error {
	n, err := matcher.Match(b)
	if err != nil {
		return err
	}
	if n != len(b) {
		return &TrailingBytesError{matcher, b}
	}
	return nil
}
