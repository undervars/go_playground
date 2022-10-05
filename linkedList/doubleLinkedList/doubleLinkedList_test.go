package doubleLinkedList

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPushBack(t *testing.T) {
	linkedList := &LinkedList[int]{}

	linkedList.PushBack(1)
	linkedList.PushBack(2)
	linkedList.PushBack(3)
	linkedList.PushBack(4)
	linkedList.PushBack(5)

	for node := linkedList.root; node != nil; node = node.next {
		prev := node.prev
		next := node.next
		prevValue := -1
		if prev != nil {
			prevValue = prev.Value
		}
		nextValue := -1
		if next != nil {
			nextValue = next.Value
		}
		fmt.Printf("%d prev: %d, next: %d\n", node.Value, prevValue, nextValue)
	}
	fmt.Println()

	assert.Equal(t, 5, linkedList.count)
	assert.Equal(t, 1, linkedList.Front().Value)
	assert.Equal(t, 5, linkedList.Back().Value)

	assert.Nil(t, linkedList.GetAt(5))
	assert.Nil(t, linkedList.GetAt(-1))
	assert.Equal(t, 1, linkedList.GetAt(0).Value)
	assert.Equal(t, 2, linkedList.GetAt(1).Value)
	assert.Equal(t, 3, linkedList.GetAt(2).Value)
	assert.Equal(t, 4, linkedList.GetAt(3).Value)
	assert.Equal(t, 5, linkedList.GetAt(4).Value)

}

func TestPushFront(t *testing.T) {
	linkedList := &LinkedList[int]{}

	linkedList.PushFront(1)
	linkedList.PushFront(2)
	linkedList.PushFront(3)

	assert.Equal(t, 3, linkedList.Front().Value)
	assert.Equal(t, 1, linkedList.Back().Value)
	assert.Equal(t, 2, linkedList.GetAt(1).Value)

	firstNode := linkedList.GetAt(0)
	secondNode := linkedList.GetAt(1)
	thirdNode := linkedList.GetAt(2)
	unknownNode := &Node[int]{Value: 1000}

	assert.True(t, linkedList.Exists(firstNode))
	assert.True(t, linkedList.Exists(secondNode))
	assert.True(t, linkedList.Exists(thirdNode))
	assert.False(t, linkedList.Exists(unknownNode))

	linkedList.PrintNodes()
	linkedList.InsertAfter(firstNode, 1919)
	assert.Equal(t, 1919, linkedList.GetAt(1).Value)
	assert.Equal(t, 4, linkedList.Count())

	secondNode = linkedList.GetAt(1)
	linkedList.PrintNodes()
	linkedList.InsertAfter(secondNode, 2929)
	linkedList.PrintNodes()
	assert.Equal(t, 2929, linkedList.GetAt(2).Value)

	thirdNode = linkedList.GetAt(2)
	linkedList.InsertAfter(thirdNode, 3939)
	linkedList.PrintNodes()
	assert.Equal(t, 3939, linkedList.GetAt(3).Value)

	// linkedList.InsertAfter(firstNode, 1919)
	// linkedList.InsertAfter(firstNode, 1919)
	// assert.Equal(t, 3, linkedList.count)
	// newNode := &Node[int]{Value: 100}
	// linkedList.InsertAfter(newNode, 1000)
	// assert.Equal(t, 3, linkedList.count)
	// linkedList.InsertAfter(linkedList.root, 1000)
	// assert.Equal(t, 4, linkedList.count)
	// for node := linkedList.root; node != nil; node = node.next {
	// 	fmt.Printf("%d ", node.Value)
	// }
	// fmt.Println()

	// assert.Equal(t, 3, linkedList.count)
	// assert.Equal(t, 1, linkedList.Front().Value)
	// assert.Equal(t, 3, linkedList.Back().Value)

	// assert.Nil(t, linkedList.GetAt(3))
	// assert.Nil(t, linkedList.GetAt(-1))
	// assert.Equal(t, 1, linkedList.GetAt(0).Value)
	// assert.Equal(t, 2, linkedList.GetAt(1).Value)
	// assert.Equal(t, 3, linkedList.GetAt(2).Value)
	// fmt.Print(linkedList.root.next)
}

func TestInsertBefore(t *testing.T) {
	linkedList := &LinkedList[int]{}

	linkedList.PushFront(1)
	linkedList.PushFront(2)
	linkedList.PushFront(3)

	firstNode := linkedList.GetAt(0)
	// secondNode := linkedList.GetAt(0)
	// firstNode := linkedList.GetAt(0)
	linkedList.InsertBefore(firstNode, 1919)
	assert.Equal(t, 1919, linkedList.GetAt(0).Value)

	secondNode := linkedList.GetAt(1)
	linkedList.InsertBefore(secondNode, 2929)
	assert.Equal(t, 2929, linkedList.GetAt(1).Value)

	lastNode := linkedList.GetAt(linkedList.count - 1)
	linkedList.InsertBefore(lastNode, 9999)
	linkedList.PrintNodes()
	assert.Equal(t, 9999, linkedList.GetAt(linkedList.count-2).Value)

	linkedList.InsertAfter(linkedList.tail, 10000)
	linkedList.InsertBefore(linkedList.root, 10000)

	second := linkedList.GetAt(1)
	linkedList.PopFront()
	assert.Equal(t, second, linkedList.Front())
}

func TestPopFront(t *testing.T) {
	var l LinkedList[int]

	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)

	assert.Equal(t, 3, l.Count())
	n := l.PopFront()
	assert.Equal(t, 1, n.Value)
	assert.Nil(t, n.next)
	assert.Nil(t, n.prev)
	assert.Equal(t, 2, l.Count())

	l.PopFront()
	l.PopFront()

	assert.Equal(t, 0, l.count)
	assert.Nil(t, l.Front())
	assert.Nil(t, l.Back())
}

func TestRemove(t *testing.T) {
	var l LinkedList[int]

	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)

	firstNode := l.GetAt(0)
	l.Remove(firstNode)
	assert.Equal(t, 2, l.Count())
	l.PopFront()

}

func TestReverse(t *testing.T) {
	var l LinkedList[int]

	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	l.PushBack(4)
	l.Reverse()
	l.PrintNodes()

	// assert.Equal(t, 4, l.Front().Value)
	// assert.Equal(t, 3, l.GetAt(0).Value)
	// assert.Equal(t, 2, l.GetAt(1).Value)
	// assert.Equal(t, 1, l.Back().Value)
}
