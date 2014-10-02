package predicate

// X represents a value to be evaluated by a Predicate.
type X interface{}

type Predicate interface {
	P(X) bool
}

type PredicateFunc func(X) bool

func (p PredicateFunc) P(x X) bool {
	return p(x)
}

func True() Predicate {
	return PredicateFunc(func(x X) bool {
		return true
	})
}

func False() Predicate {
	return PredicateFunc(func(x X) bool {
		return false
	})
}

func And(s []Predicate) Predicate {
	return PredicateFunc(func(x X) bool {
		for _, p := range s {
			if !p.P(x) {
				return false
			}
		}

		return true
	})
}
