package linkedList

import (
	"testing"

	"github.com/undervars/playground/linkedList/doubleLinkedList"
	"github.com/undervars/playground/linkedList/singleLinkedList"
)

func BenchmarkSingleLinkedListReverse(b *testing.B) {
	var sl singleLinkedList.LinkedList[int]

	for i := 0; i < b.N; i++ {
		sl.PushBack(i)
	}

	sl.Reverse()
}

func BenchmarkSingleLinkedListReverseSwap(b *testing.B) {
	var sl singleLinkedList.LinkedList[int]

	for i := 0; i < b.N; i++ {
		sl.PushBack(i)
	}

	sl.ReverseSwap()
}

func BenchmarkDoubleLinkedListReverse(b *testing.B) {
	var dl doubleLinkedList.LinkedList[int]

	for i := 0; i < b.N; i++ {
		dl.PushBack(i)
	}

	dl.Reverse()
}
