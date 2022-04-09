package datastructure

// ArrayList is a list of type V backed by a slice
type arrayList[E interface{}] struct {
	length   int32
	capacity int32
	backing  []E
}

func New[E any]() arrayList {
	out := arrayList{
		length:   0,
		capacity: 2,
		backing:  make([]E, capacity),
	}

	return out
}

func (l arrayList) Size() int32 { return l.length }

func (l *arrayList[E]) Add(e E) {
	if length+1 > capacity {
		capacity *= 2
	}
}
