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

package predicate

import "testing"

func TestTrue(t *testing.T) {
	if True().P(nil) != true {
		t.Fail()
	}
}

func TestFalse(t *testing.T) {
	if False().P(nil) != false {
		t.Fail()
	}
}

func TestAnd(t *testing.T) {
	var p Predicate

	// true && true -> true
	p = And([]Predicate{
		True(),
		True(),
	})
	if p.P(nil) != true {
		t.Fail()
	}

	// true && false -> false
	p = And([]Predicate{
		True(),
		False(),
	})
	if p.P(nil) != false {
		t.Fail()
	}

	// false && false -> false
	p = And([]Predicate{
		False(),
		False(),
	})
	if p.P(nil) != false {
		t.Fail()
	}
}

func TestOr(t *testing.T) {
	var p Predicate

	// true || true -> true
	p = Or([]Predicate{
		True(),
		True(),
	})
	if p.P(nil) != true {
		t.Fail()
	}

	// true || false -> true
	p = Or([]Predicate{
		True(),
		False(),
	})
	if p.P(nil) != true {
		t.Fail()
	}

	// false || false -> false
	p = Or([]Predicate{
		False(),
		False(),
	})
	if p.P(nil) != false {
		t.Fail()
	}
}

func TestXor(t *testing.T) {
	var p Predicate

	// XOR [] -> false
	p = Xor([]Predicate{})
	if p.P(nil) != false {
		t.Fail()
	}

	// XOR nil -> false
	p = Xor(nil)
	if p.P(nil) != false {
		t.Fail()
	}

	// XOR true -> true
	p = Xor([]Predicate{
		True(),
	})
	if p.P(nil) != true {
		t.Fail()
	}

	// XOR false -> false
	p = Xor([]Predicate{
		False(),
	})
	if p.P(nil) != false {
		t.Fail()
	}

	// true XOR true -> false
	p = Xor([]Predicate{
		True(),
		True(),
	})
	if p.P(nil) != false {
		t.Fail()
	}

	// true XOR false -> true
	p = Xor([]Predicate{
		True(),
		False(),
	})
	if p.P(nil) != true {
		t.Fail()
	}

	// false XOR false -> false
	p = Xor([]Predicate{
		False(),
		False(),
	})
	if p.P(nil) != false {
		t.Fail()
	}
}

func TestNot(t *testing.T) {
	var p Predicate

	// !true -> false
	p = Not(True())
	if p.P(nil) != false {
		t.Fail()
	}

	// !false -> true
	p = Not(False())
	if p.P(nil) != true {
		t.Fail()
	}

	// !true && !true -> false
	p = Not(True(), True())
	if p.P(nil) != false {
		t.Fail()
	}

	// !true && !false -> false
	p = Not(True(), False())
	if p.P(nil) != false {
		t.Fail()
	}

	// !false && !false -> true
	p = Not(False(), False())
	if p.P(nil) != true {
		t.Fail()
	}
}

func TestExists(t *testing.T) {
	var p Predicate

	p = Exists("a", 1)
	if p.P(map[string]interface{}{"a": 1}) != true {
		t.Fail()
	}
}
