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
