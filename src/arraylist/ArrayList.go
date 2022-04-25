package arraylist

import (
	"fmt"

	"github.com/ValiantChip/datastructure/derr"
)

// ArrayList is a list of type E backed by a slice
type ArrayList[E any] struct {
	length   int
	capacity int
	backing  []E
}

// NewArrayList returns a new ArrayList of type E
func NewArrayList[E any]() *ArrayList[E] {
	capacity := 2
	return &ArrayList[E]{
		length:   0,
		capacity: capacity,
		backing:  make([]E, capacity),
	}
}

func (l *ArrayList[any]) Size() int { return l.length }

func (l *ArrayList[E]) Add(e E) {
	l.checkCapacity()
	l.backing[l.length] = e
	l.length++
}

func (l *ArrayList[E]) Set(e E, index int) {
	if index > l.length {
		panic(derr.IndexOutOfBoundsError{Bounds: l.length, RecievedIndex: index})
	}

	l.checkCapacity()
	for i := l.length; i > index; i-- {
		l.backing[i] = l.backing[i-1]
	}

	l.backing[index] = e
	l.length++
}

func (l *ArrayList[E]) Get(index int) E {
	if index >= l.length {
		panic(derr.IndexOutOfBoundsError{Bounds: l.length, RecievedIndex: index})
	}

	return l.backing[index]
}

func (l *ArrayList[E]) Remove(index int) E {
	if index >= l.length {
		panic(derr.IndexOutOfBoundsError{Bounds: l.length, RecievedIndex: index})
	}

	out := l.backing[index]

	for i := index; i < l.length-1; i++ {
		l.backing[i] = l.backing[i+1]
	}

	var empty E
	l.backing[l.length-1] = empty

	return out
}

func (l *ArrayList[E]) AsSlice() []E {
	temp := make([]E, l.length)

	for i, e := range l.backing {
		temp[i] = e
	}

	return temp
}

func (l *ArrayList[E]) checkCapacity() {
	if l.length+1 > l.capacity {
		l.capacity *= 2
		temp := make([]E, l.capacity)
		_, err := copyInto(l.backing, temp)
		if err != nil {
			panic(err)
		}
		l.backing = temp
		return
	}

	if l.length < l.capacity/3 {
		l.capacity /= 2
		temp := make([]E, l.capacity)
		_, err := copyInto(l.backing, temp)
		if err != nil {
			panic(err)
		}
		l.backing = temp
		return
	}
}

func copyInto[E any](origin []E, target []E) ([]E, error) {
	if cap(target) < len(origin) {
		return target, fmt.Errorf("origin slice to large to be copied to target; Size: %d, Required: %d", cap(target), len(origin))
	}

	for i, t := range origin {
		target[i] = t
	}

	return target, nil
}
