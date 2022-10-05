package singleLinkedList

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPushBack(t *testing.T) {
	var l LinkedList[int]

	assert.Nil(t, l.root)
	assert.Nil(t, l.tail)

	l.PushBack(1)
	assert.NotNil(t, l.root)
	assert.NotNil(t, l.tail)
	assert.Equal(t, l.Front().value, 1)
	assert.Equal(t, l.Back().value, 1)

	l.PushBack(2)
	assert.NotNil(t, l.root)
	assert.NotNil(t, l.tail)
	assert.Equal(t, l.Front().value, 1)
	assert.Equal(t, l.Back().value, 2)

	l.PushBack(3)
	assert.NotNil(t, l.root)
	assert.NotNil(t, l.tail)
	assert.Equal(t, l.Front().value, 1)
	assert.Equal(t, l.Back().value, 3)

	assert.Equal(t, l.Count(), 3)

	l.PushBack(4)
	assert.NotNil(t, l.root)
	assert.NotNil(t, l.tail)
	assert.Equal(t, l.Front().value, 1)
	assert.Equal(t, l.Back().value, 4)

	assert.Equal(t, l.Count(), 4)

	assert.Equal(t, 1, l.GetAt(0).value)
	assert.Equal(t, 2, l.GetAt(1).value)
	assert.Equal(t, 3, l.GetAt(2).value)
	assert.Equal(t, 4, l.GetAt(3).value)
	assert.Nil(t, l.GetAt(4))

	secondNode := l.GetAt(1)
	l.InsertAfter(secondNode, 999)
	assert.Equal(t, 5, l.Count())
	assert.Equal(t, 999, l.GetAt(2).value)

	tmpNode := &Node[int]{value: 10000}
	l.InsertAfter(tmpNode, 1000)
	assert.Equal(t, 5, l.Count())

	assert.Equal(t, 2, l.GetAt(1).value)
	l.InsertBefore(secondNode, 1919)
	assert.Equal(t, 6, l.Count())

	assert.Equal(t, 1919, l.GetAt(1).value)

	l.InsertBefore(l.root, 191919)
	assert.Equal(t, 191919, l.GetAt(0).value)

	l.PrintElems()
}

func TestPushFront(t *testing.T) {
	var l LinkedList[int]

	assert.Nil(t, l.root)
	assert.Nil(t, l.tail)

	l.PushFront(1)
	assert.NotNil(t, l.root)
	assert.NotNil(t, l.tail)
	assert.Equal(t, l.Front().value, 1)
	assert.Equal(t, l.Back().value, 1)

	l.PushFront(2)
	assert.NotNil(t, l.root)
	assert.NotNil(t, l.tail)
	assert.Equal(t, l.Front().value, 2)
	assert.Equal(t, l.Back().value, 1)

	l.PushFront(3)
	assert.NotNil(t, l.root)
	assert.NotNil(t, l.tail)
	assert.Equal(t, l.Front().value, 3)
	assert.Equal(t, l.Back().value, 1)

	assert.Equal(t, l.Count(), 3)

	l.PushFront(4)
	assert.NotNil(t, l.root)
	assert.NotNil(t, l.tail)
	assert.Equal(t, l.Front().value, 4)
	assert.Equal(t, l.Back().value, 1)

	assert.Equal(t, l.Count(), 4)

	assert.Equal(t, 4, l.GetAt(0).value)
	assert.Equal(t, 3, l.GetAt(1).value)
	assert.Equal(t, 2, l.GetAt(2).value)
	assert.Equal(t, 1, l.GetAt(3).value)
	assert.Nil(t, l.GetAt(4))
}

func TestPopFront1(t *testing.T) {
	var l LinkedList[int]

	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	l.PopFront()

	assert.Equal(t, 2, l.Count())
	assert.Equal(t, 2, l.Front().value)
	assert.Equal(t, 3, l.Back().value)

	l.PopFront()
	assert.Equal(t, 1, l.Count())
	assert.Equal(t, 3, l.Front().value)
	assert.Equal(t, 3, l.Back().value)

	l.PopFront()
	assert.Equal(t, 0, l.Count())
	assert.Nil(t, l.Front())
	assert.Nil(t, l.Back())
}

func TestRemove(t *testing.T) {
	var l LinkedList[int]

	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)

	l.Remove(l.GetAt(1))

	assert.Equal(t, 2, l.Count())
	assert.Equal(t, 1, l.Front().value)
	assert.Equal(t, 3, l.Back().value)

	l.Remove(l.GetAt(0))
	fmt.Println(l)
	l.Remove(&Node[int]{value: 100})
	fmt.Println(l)
	assert.Equal(t, 1, l.Count())
	assert.Equal(t, 3, l.Front().value)
	assert.Equal(t, 3, l.Back().value)

	l.Remove(l.GetAt(0))
	assert.Equal(t, 0, l.Count())
	assert.Nil(t, l.root)
	assert.Nil(t, l.tail)
	assert.Nil(t, l.Front())
	assert.Nil(t, l.Back())

	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	l.Remove(l.GetAt(2))
	assert.Equal(t, 2, l.Back().value)

}

func TestPopFront(t *testing.T) {
	var l LinkedList[int]

	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	assert.Equal(t, 3, l.Count())
	p1 := l.PopFront()
	assert.Equal(t, 2, l.Count())
	assert.Equal(t, 1, p1.value)
	p2 := l.PopFront()
	assert.Equal(t, 1, l.Count())
	p3 := l.PopFront()
	assert.Equal(t, 0, l.Count())
	assert.Equal(t, 2, p2.value)
	assert.Equal(t, 3, p3.value)

	assert.Nil(t, p1.next)
	assert.Nil(t, p2.next)
	assert.Nil(t, p3.next)
}

func TestReverse(t *testing.T) {
	var l LinkedList[int]

	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	l.PushBack(4)
	assert.Equal(t, 4, l.Count())
	l.Reverse()
	assert.Equal(t, 4, l.Count())

	l.PopFront()
	assert.Equal(t, 3, l.Count())
	l.PopFront()
	l.PopFront()
	assert.NotNil(t, l.tail)
	l.PopFront()
	assert.Equal(t, 0, l.Count())
	assert.Nil(t, l.tail)
	// assert.Equal(t, 4, l.Front().value)
	// assert.Equal(t, 3, l.GetAt(0).value)
	// assert.Equal(t, 2, l.GetAt(1).value)
	// assert.Equal(t, 1, l.Back().value)
}

func TestReverseSwap(t *testing.T) {
	var l LinkedList[int]

	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	l.PushBack(4)
	assert.Equal(t, 4, l.Count())
	l.ReverseSwap()

	assert.Equal(t, l.root, l.Front())
	assert.Equal(t, l.tail, l.Back())
	assert.Equal(t, l.tail, l.GetAt(3))
	assert.Equal(t, l.root, l.GetAt(0))
	l.PrintElems()
}
