package starter_code

type LinkedList[T comparable] struct {
	head *Node[T]
	size int
}

type Node[T comparable] struct {
	Val  T
	Next *Node[T]
}

func (ll *LinkedList[T]) Size() int {
	if ll != nil {
		return ll.size
	}
	return 0
}

func (ll *LinkedList[T]) Add(newVal T) {
	// TODO: Implement
}

func (ll *LinkedList[T]) Remove(val T) {
	// TODO: Implement
}

func (ll *LinkedList[T]) Contains(val T) bool {
	// TODO: Implement
	return false
}
