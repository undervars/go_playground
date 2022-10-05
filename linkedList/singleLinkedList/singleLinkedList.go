package singleLinkedList

import "fmt"

type Node[T any] struct {
	next  *Node[T]
	value T
}

type LinkedList[T any] struct {
	root  *Node[T]
	tail  *Node[T]
	count int
}

func (l *LinkedList[T]) ReverseSwap() {
	if l.count <= 1 {
		return
	}

	head := l.root
	cursor := head.next

	for cursor != nil {
		n := cursor
		cursor = cursor.next

		n.next = head
		head = n
	}

	l.root, l.tail = l.tail, l.root
	l.tail.next = nil
}

func (l *LinkedList[T]) Reverse() {
	if l.count <= 1 {
		return
	}

	newL := &LinkedList[T]{}

	for node := l.PopFront(); node != nil; node = l.PopFront() {
		newL.PushFront(node.value)
	}

	l.root = newL.root
	l.tail = newL.tail
	l.count = newL.count
}

func (l *LinkedList[T]) PushBack(value T) {
	new := &Node[T]{
		next:  nil,
		value: value,
	}

	if l.root == nil {
		l.root = new
		l.tail = new
		l.count = 1
		return
	}

	l.tail.next = new
	l.tail = new
	l.count++
}

func (l *LinkedList[T]) PushFront(value T) {
	new := &Node[T]{
		value: value,
	}

	if l.root == nil {
		l.root = new
		l.tail = new
		l.count = 1
		return
	}

	new.next = l.root
	l.root = new
	l.count++
}

func (l *LinkedList[T]) PrintElems() {
	node := l.root
	for node != nil {
		fmt.Printf("%v ", node.value)
		node = node.next
	}

	fmt.Println()
}

func (l *LinkedList[T]) Front() *Node[T] {
	return l.root
}

func (l *LinkedList[T]) Back() *Node[T] {
	return l.tail
}

func (l *LinkedList[T]) Count() int {
	return l.count
}

func (l *LinkedList[T]) GetAt(index int) *Node[T] {

	if index >= l.count {
		return nil
	}

	node := l.root
	for i := 0; i < index; i++ {
		node = node.next
	}
	return node
}

func (l *LinkedList[T]) Exists(node *Node[T]) bool {

	for n := l.root; n != nil; n = n.next {
		if n == node {
			return true
		}
	}

	return false
}

func (l *LinkedList[T]) InsertBefore(node *Node[T], value T) {
	if node == l.root {
		l.PushFront(value)
		return
	}

	prevNode := l.findPrevNode(node)
	if prevNode == nil {
		return
	}

	newNode := &Node[T]{
		value: value,
		next:  node,
	}

	prevNode.next = newNode
	l.count++
}

func (l *LinkedList[T]) Remove(node *Node[T]) {
	if node == l.root {
		l.PopFront()
		return
	}

	prevNode := l.findPrevNode(node)
	if prevNode == nil {
		return
	}

	prevNode.next = node.next
	node.next = nil
	if node == l.tail {
		l.tail = prevNode
	}
	l.count--

}

func (l *LinkedList[T]) PopFront() *Node[T] {
	if l.root == nil {
		return nil
	}

	n := l.root
	l.root.next, l.root = nil, l.root.next
	if l.root == nil {
		l.tail = nil
	}
	l.count--

	return n
}

func (l *LinkedList[T]) findPrevNode(nextNode *Node[T]) *Node[T] {
	node := l.root
	for node.next != nil {
		if node.next == nextNode {
			return node
		}
		node = node.next
	}
	return nil
}

func (l *LinkedList[T]) InsertAfter(node *Node[T], value T) {
	if !l.Exists(node) {
		return
	}

	newNode := &Node[T]{
		value: value,
		next:  node.next,
	}
	node.next = newNode

	if node == l.tail {
		l.tail = node
	}
	l.count++
}

func ExecSingleLikedList() {
	linkedList := LinkedList[int]{}

	// linkedList.PushBack(2)
	// linkedList.PushBack(3)
	// linkedList.PushBack(4)
	// linkedList.PushBack(5)

	linkedList.PushFront(2)
	linkedList.PushFront(3)
	linkedList.PushFront(4)
	linkedList.PushFront(5)
	linkedList.PrintElems()

	fmt.Println(linkedList.Front())
	fmt.Println(linkedList.Back())

}
