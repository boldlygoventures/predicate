/*
The MIT License (MIT)

Copyright (c) 2014 Boldly Go Ventures

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

// Package predicate defines the Predicate interface and provides basic predicates.
package predicate

// X represents a value to be evaluated by a Predicate.
type X interface{}

type Predicate interface {
	P(X) bool
}

// PredicateFunc maps X to bool
type PredicateFunc func(X) bool

// P satisfies the Predicate interface.
func (p PredicateFunc) P(x X) bool {
	return p(x)
}

// True returns a PredicateFunc that will always return true, for any x.
func True() Predicate {
	return PredicateFunc(func(x X) bool {
		return true
	})
}

// True returns a PredicateFunc that will always return false, for any x.
func False() Predicate {
	return PredicateFunc(func(x X) bool {
		return false
	})
}

type Set []Predicate

type And Set

// P returns true if and only all of its member Predicates is true for x. The logic is short circuited,
// returning false when a member Predicate is false.
func (p And) P(x X) bool {
	for _, p := range p {
		if !p.P(x) {
			return false
		}
	}

	return true
}

type Or Set

// P returns true if any of its member Predicates is true for x. The logic is short circuited,
// returning true when a member Predicate is true.
func (p Or) P(x X) bool {
	for _, p := range p {
		if p.P(x) {
			return true
		}
	}

	return false
}

type Xor Set

// P returns true if and only if one of its member Predicates is true for x. The logic is short circuited, returning
// false when a second member Predicate is true.
func (p Xor) P(x X) bool {
	var n int
	for _, p := range p {
		if p.P(x) {
			n++
		}

		// short circuit
		if n > 1 {
			return false
		}
	}

	return n == 1
}

// Not returns a PredicateFunc that will return false if any of the passed Predicates is true for x.
func Not(p ...Predicate) Predicate {
	return PredicateFunc(func(x X) bool {
		return !Or(p).P(x)
	})
}

// Exists returns a PredicateFunc that will return true only if x[k] is in set s.
func Exists(k string, s interface{}) Predicate {
	return PredicateFunc(func(x X) bool {
		if y, ok := x.(map[string]interface{}); ok {
			x = y[k]
		}

		switch s := s.(type) {
		case map[string]interface{}:
			for _, v := range s {
				if v == x {
					return true
				}
			}
		case []interface{}:
			for _, v := range s {
				if v == x {
					return true
				}
			}
		default: //json: string, float64, bool, nil
			return s == x
		}

		return false
	})
}
