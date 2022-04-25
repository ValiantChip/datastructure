package linkedlist

import (
	"github.com/ValiantChip/datastructure/derr"
)

type LinkedList[E any] struct {
	head *section[E]
	tail *section[E]

	len int
}

type section[E any] struct {
	Elem E

	Next *section[E]
}

func NewLinkedList[E any]() *LinkedList[E] {
	init := &section[E]{}

	return &LinkedList[E]{
		head: init,
		tail: init,
		len:  0,
	}
}

func (l *LinkedList[any]) Size() int { return l.len }

func (l *LinkedList[E]) Add(e E) {
	l.tail.Next = &section[E]{
		Elem: e,
	}

	l.tail = l.tail.Next

	l.len++
}

func (l *LinkedList[E]) Get(index int) E {
	if index >= l.len {
		panic(derr.IndexOutOfBoundsError{Bounds: l.len, RecievedIndex: index})
	}

	if index == l.len-1 {
		return l.tail.Elem
	}

	e := l.head.Next
	for i := 0; i < index; i++ {
		e = e.Next
	}

	return e.Elem
}

func (l *LinkedList[E]) Set(e E, index int) {
	if index >= l.len {
		panic(derr.IndexOutOfBoundsError{Bounds: l.len, RecievedIndex: index})
	}

	if index == l.len-1 {
		l.tail.Elem = e
	}

	s := l.head.Next
	for i := 0; i < index-1; i++ {
		s = s.Next
	}

	s.Next.Elem = e
}

func (l *LinkedList[E]) Remove(index int) E {
	if index >= l.len {
		panic(derr.IndexOutOfBoundsError{Bounds: l.len, RecievedIndex: index})
	}

	e := l.head.Next
	for i := 0; i < index-1; i++ {
		e = e.Next
	}

	out := e.Next.Elem
	e.Next = e.Next.Next

	return out
}

func (l *LinkedList[E]) AsSlice() []E {
	out := make([]E, l.len)

	cur := l.head.Next
	for i := 0; i < l.len; i++ {
		out[i] = cur.Elem
		cur = cur.Next
	}

	return out
}
