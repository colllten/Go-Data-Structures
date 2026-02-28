package solution

type LinkedList[T comparable] struct {
	Head *Node[T]
	Size uint
}

type Node[T comparable] struct {
	Value T
	Next  *Node[T]
}

func (ll *LinkedList[T]) Add(newValue T) {
	if ll == nil {
		return
	}

	newNode := &Node[T]{
		Value: newValue,
		Next:  nil,
	}

	if ll.Head == nil {
		ll.Head = newNode
		ll.Size++
		return
	}

	newNode.Next = ll.Head
	ll.Head = newNode
	ll.Size++
}

func (ll LinkedList[T]) ContainsIter(val T) bool {
	temp := ll.Head
	for temp != nil {
		if temp.Value == val {
			return true
		}
		temp = temp.Next
	}

	return false
}

func (ll LinkedList[T]) ContainsRecur(val T) bool {
	return containsValRecurs(ll.Head, val)
}

func containsValRecurs[T comparable](node *Node[T], val T) bool {
	if node == nil {
		return false
	}

	if node.Value == val {
		return true
	}

	return containsValRecurs(node.Next, val)
}

func (ll *LinkedList[T]) Remove(val T) {
	if ll == nil || ll.Size == 0 {
		return
	}

	if ll.Size == 1 && ll.Head.Value == val {
		ll.Head = nil
		ll.Size--
		return
	}

	prev := ll.Head
	temp := ll.Head.Next

	if prev.Value == val {
		ll.Head = temp
		ll.Size--
		return
	}

	for temp != nil {
		if temp.Value == val {
			prev.Next = temp.Next
			ll.Size--
			return
		}
		prev = temp
		temp = temp.Next
	}
}
