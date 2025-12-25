package deque

import "container/list"

type Deque[T comparable] struct {
	list *list.List
}

func New[T comparable]() *Deque[T] {
	return &Deque[T]{
		list: list.New(),
	}
}

func (d *Deque[T]) IsEmpty() bool {
	return d.list.Len() == 0
}

func (d *Deque[T]) PeekLeft() (T, bool) {
	if d.IsEmpty() {
		var zero T
		return zero, false
	}
	return d.list.Front().Value.(T), true
}

func (d *Deque[T]) PopLeft() (T, bool) {
	if d.IsEmpty() {
		var zero T
		return zero, false
	}

	head := d.list.Front()
	d.list.Remove(head)
	return head.Value.(T), true
}

func (d *Deque[T]) PushLeft(elem T) {
	d.list.PushFront(elem)
}

func (d *Deque[T]) PeekRight() (T, bool) {
	if d.IsEmpty() {
		var zero T
		return zero, false
	}
	return d.list.Back().Value.(T), true
}

func (d *Deque[T]) PopRight() (T, bool) {
	if d.IsEmpty() {
		var zero T
		return zero, false
	}

	tail := d.list.Back()
	d.list.Remove(tail)
	return tail.Value.(T), true
}

func (d *Deque[T]) PushRight(elem T) {
	d.list.PushBack(elem)
}
