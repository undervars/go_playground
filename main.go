package main

import "fmt"

func AddArr(arr [3]int) {
	arr[0] = 100
	fmt.Println(arr)
}

func main() {
	a := [3]int{1, 2, 3}

	fmt.Println(a)
	AddArr(a)

	fmt.Println(a)
	// countingSort.SortHeight()
}
