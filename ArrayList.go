package datastructure

import "fmt"

// ArrayList is a list of type V backed by a slice
type ArrayList[E any] struct {
	values inner[E]
}

type inner[E any] struct {
	length   int
	capacity int
	backing  []E
}

// NewArrayList returns a new ArrayList of type E
func NewArrayList[E any]() ArrayList[E] {
	capacity := 2
	out := ArrayList[E]{
		values: inner[E]{
			length:   0,
			capacity: capacity,
			backing:  make([]E, capacity),
		},
	}

	return out
}

func test() {
	list := NewArrayList[int]()
	_ = list
}

func (l *ArrayList[any]) Size() int { return l.values.length }

func (l *ArrayList[E]) Add(e E) {
	l.checkCapacity()
	l.values.backing[l.values.length] = e
	l.values.length++
}

func (l *ArrayList[E]) Set(e E, index int) {
	if index > l.values.length {
		panic(fmt.Errorf("Index %d out of range %d", index, l.values.length-1))
	}

	l.checkCapacity()
	for i := l.values.length; i > index; i-- {
		l.values.backing[i] = l.values.backing[i-1]
	}

	l.values.backing[index] = e
	l.values.length++
}

func (l *ArrayList[E]) Get(index int) E {
	if index >= l.values.length {
		panic(fmt.Errorf("Index %d out of range %d", index, l.values.length-1))
	}

	return l.values.backing[index]
}

func (l *ArrayList[E]) Remove(index int) E {
	if index >= l.values.length {
		panic(fmt.Errorf("Index %d out of range %d", index, l.values.length-1))
	}

	out := l.values.backing[index]

	for i := index; i < l.values.length-1; i++ {
		l.values.backing[i] = l.values.backing[i+1]
	}

	var empty E
	l.values.backing[l.values.length-1] = empty

	return out
}

func (l *ArrayList[E]) checkCapacity() {
	if l.values.length+1 > l.values.capacity {
		l.values.capacity *= 2
		temp := make([]E, l.values.capacity)
		_, err := copyInto(l.values.backing, temp)
		if err != nil {
			panic(err)
		}
		l.values.backing = temp
		return
	}

	if l.values.length < l.values.capacity/3 {
		l.values.capacity /= 2
		temp := make([]E, l.values.capacity)
		_, err := copyInto(l.values.backing, temp)
		if err != nil {
			panic(err)
		}
		l.values.backing = temp
		return
	}
}

func copyInto[E any](origin []E, target []E) ([]E, error) {
	if cap(target) < len(origin) {
		return target, fmt.Errorf("Origin slice to large to be copied to target; Size: %d, Required: %d", cap(target), len(origin))
	}

	for i, t := range origin {
		target[i] = t
	}

	return target, nil
}
