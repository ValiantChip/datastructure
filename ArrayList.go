package datastructure

// ArrayList is a list of type V backed by a slice
type arrayList[E any] struct {
	length   int32
	capacity int32
	backing  []E
}

func New[E any]() arrayList[E] {
	capacity := int32(2)
	out := arrayList[E]{
		length:   0,
		capacity: capacity,
		backing:  make([]E, capacity),
	}

	return out
}

func (l arrayList[any]) Size() int32 { return l.length }

func (l *arrayList[E]) Add(e E) {
	if l.length+1 > l.capacity {
		l.capacity *= 2
	}
}
