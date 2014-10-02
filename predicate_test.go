package predicate

import "testing"

func TestTrue(t *testing.T) {
	if !True().P(nil) {
		t.Fail()
	}
}
