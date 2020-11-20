// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package triangle should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package triangle

import "math"

// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind string

const (
	NaT = "Not a Triangle"
	Equ = "Equilateral"
	Iso = "Isosceles"
	Sca = "Scalene"
)

func KindFromSides(a, b, c float64) Kind {
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!
	var k Kind

	if math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) {
		return NaT
	}
	if math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0) {
		return NaT
	}
	if a <= 0 || b <= 0 || c <= 0 || a+b < c || a+c < b || b+c < a {
		k = NaT
	} else {
		if a == b && b == c {
			k = Equ
		} else if a == b || b == c || a == c {
			k = Iso
		} else {
			k = Sca
		}
	}
	return k
}
