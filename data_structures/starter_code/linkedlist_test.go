package starter_code

import (
	"testing"
)

func Test_newLlSizeIsZero(t *testing.T) {
	ll := LinkedList[int]{}

	size := ll.size
	if size != 0 {
		t.Error("incorrect result: expected 0, but got", size)
	}
}

func Test_newLlHeadIsNil(t *testing.T) {
	ll := LinkedList[int]{}

	if ll.head != nil {
		t.Error("incorrect result: head is not nil")
	}
}

func Test_addVal(t *testing.T) {
	ll := LinkedList[int]{}
	ll.Add(1)

	if ll.head == nil {
		t.Error("incorrect result: head is nil")
	} else if ll.size != 1 {
		t.Error("incorrect result: expected size of 1, got", ll.size)
	} else if ll.head.Next != nil {
		t.Error("incorrect result: head.next is not nil")
	}
}
