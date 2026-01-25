package generics

type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

// Push adds a value to the end of the list.
func (lst *List[T]) Push(v T) {
	newElement := &element[T]{val: v}

	if lst.head == nil {
		lst.head = newElement
		lst.tail = newElement
	} else {
		lst.tail.next = newElement
		lst.tail = newElement
	}
}

// Pop removes the last element and returns it.
func (lst *List[T]) Pop() (T, bool) {
	var zero T

	if lst.head == nil {
		return zero, false
	}

	if lst.head == lst.tail {
		val := lst.head.val
		lst.head = nil
		lst.tail = nil
		return val, true
	}

	current := lst.head
	for current.next != lst.tail {
		current = current.next
	}
	val := lst.tail.val
	current.next = nil
	lst.tail = current
	return val, true
}

// AllElements returns a slice of all elements in the list.
func (lst *List[T]) AllElements() []T {
	var result []T
	for current := lst.head; current != nil; current = current.next {
		result = append(result, current.val)
	}
	return result
}
