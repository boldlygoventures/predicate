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
