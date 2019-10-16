package queue

// TODO: Implement Queues
// INSERT(S, x) inserts the element x into the set S. This
// operation could be written as S ← S ∪ {x}.
// MAXIMUM(S) returns the element of S with the largest key.
// EXTRACT-MAX(S) removes and returns the element of S with the
// largest key.
// INCREASE-KEY (S, x, k) increases the value of element x’s
// key to the new value k, which is assumed to be at least as
// large as x’s current key value.

type MaxPriorityQueue interface {
	Insert()
	Maximum()
	ExtractMax()
	IncreaseKey()
}

type MinPriorityQueue interface {
	Insert()
	Minimum()
	ExtractMin()
	IncreaseKey()
}
