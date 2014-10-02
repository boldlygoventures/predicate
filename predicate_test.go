package predicate

import "testing"

func TestTrue(t *testing.T) {
	if !True().P(nil) {
		t.Fail()
	}
}

func TestFalse(t *testing.T) {
	if False().P(nil) {
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
	if !p.P(nil) {
		t.Fail()
	}

	// true && false -> false
	p = And([]Predicate{
		True(),
		False(),
	})
	if p.P(nil) {
		t.Fail()
	}

	// false && false -> false
	p = And([]Predicate{
		False(),
		False(),
	})
	if p.P(nil) {
		t.Fail()
	}
}
