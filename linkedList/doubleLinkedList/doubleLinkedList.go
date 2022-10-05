package doubleLinkedList

import "fmt"

type Node[T any] struct {
	Value T
	next  *Node[T]
	prev  *Node[T]
}
type LinkedList[T any] struct {
	root  *Node[T]
	tail  *Node[T]
	count int
}

func (l *LinkedList[T]) Front() *Node[T] {
	return l.root
}
func (l *LinkedList[T]) Back() *Node[T] {
	return l.tail
}
func (l *LinkedList[T]) GetAt(index int) *Node[T] {
	if index >= l.count || index < 0 {
		return nil
	}

	var node *Node[T]

	// node = l.root
	// for i := 0; i < index; i++ {
	// 	node = node.next
	// }
	// node = l.tail
	// for i := l.count - 1; i > index; i-- {
	// 	node = node.prev
	// }
	// // from head
	if index < l.count/2 {
		node = l.root
		for i := 0; i < index; i++ {
			node = node.next
		}
		// from tail
	} else {
		// 5
		// 2

		// 4 3
		node = l.tail
		for i := l.count - 1; i > index; i-- {
			node = node.prev
		}
	}

	return node
}

func (l *LinkedList[T]) Count() int {
	return l.count
}

func (l *LinkedList[T]) PushFront(val T) {
	newNode := &Node[T]{
		Value: val,
	}

	if l.count == 0 {
		l.root = newNode
		l.tail = newNode
		l.count = 1
		return
	}

	l.root.prev = newNode
	newNode.next = l.root
	l.root = newNode
	l.count++
}

func (l *LinkedList[T]) Reverse() {
	if l.count <= 1 {
		return
	}

	// var next *Node[T]

	for node := l.root; node != nil; node = node.prev {
		// next = node.next

		node.prev, node.next = node.next, node.prev

		// node = next
	}

	l.root, l.tail = l.tail, l.root
}

func (l *LinkedList[T]) PushBack(val T) {
	newNode := &Node[T]{
		Value: val,
	}

	if l.count == 0 {
		l.root = newNode
		l.tail = newNode
		l.count = 1
		return
	}

	newNode.prev = l.tail
	l.tail.next = newNode
	l.tail = newNode
	l.count++
}

func (l *LinkedList[T]) Exists(node *Node[T]) bool {
	c := l.root
	for ; c != nil; c = c.next {
		if c == node {
			return true
		}
	}
	return false
}

func (l *LinkedList[T]) Remove(node *Node[T]) {
	if !l.Exists(node) {
		return
	}

	if node == l.root {
		l.PopFront()
		return
	}
	if node == l.tail {
		l.PopBack()
		return
	}

	prev := node.prev
	next := node.next

	prev.next = next
	next.prev = prev

	node.prev = nil
	node.next = nil
	l.count--
}

func (l *LinkedList[T]) PopBack() *Node[T] {
	if l.tail == nil {
		return nil
	}

	if l.count == 1 {
		lastTail := l.tail
		l.root = nil
		l.tail = nil
		l.count = 0
		return lastTail
	}

	originalTail := l.tail
	newTail := originalTail.prev
	newTail.next = nil
	l.tail = newTail

	originalTail.prev = nil
	originalTail.next = nil
	l.count--

	return originalTail
}
func (l *LinkedList[T]) PopFront() *Node[T] {
	if l.root == nil {
		return nil
	}

	originalFront := l.root
	if originalFront.next == nil {
		l.count = 0
		l.root = nil
		l.tail = nil
		return originalFront
	}

	front := l.root.next

	front.prev = nil
	l.root = front
	originalFront.next = nil

	l.count--
	return originalFront
}

func (l *LinkedList[T]) InsertBefore(node *Node[T], val T) {
	if l.count == 0 || !l.Exists(node) {
		return
	}

	if node == l.root {
		l.PushFront(val)
		return
	}

	nextNode := node
	prevNode := node.prev
	newNode := &Node[T]{
		Value: val,
		prev:  prevNode,
		next:  nextNode,
	}
	prevNode.next = newNode
	nextNode.prev = newNode
	l.count++
}
func (l *LinkedList[T]) InsertAfter(node *Node[T], val T) {
	if l.count == 0 || !l.Exists(node) {
		return
	}

	if node == l.tail {
		l.PushBack(val)
		return
	}

	nextNode := node.next

	newNode := &Node[T]{
		Value: val,
		prev:  node,
		next:  nextNode,
	}
	node.next = newNode

	nextNode.prev = newNode

	l.count++
}

func (l *LinkedList[T]) PrintNodes() {
	node := l.root
	for ; node != nil; node = node.next {
		p := node.prev
		n := node.next

		var pv T
		var nv T
		if p != nil {
			pv = p.Value
		}
		if n != nil {
			nv = n.Value
		}
		fmt.Printf("%v, prev: %v, next: %v\n", node.Value, pv, nv)
	}
	fmt.Println()
}
