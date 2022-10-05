package countingSort

import "fmt"

func SortLowerCases() {
	str := "dslkjfdslkfjeqwmkljflksdvciuewksdahjsdh"

	var count [26]int
	var maxChar rune
	var maxCount int

	for i := 0; i < len(str); i++ {
		c := rune(str[i])
		count[c-'a'] += 1
	}

	for i := 0; i < len(count); i++ {
		if maxCount < count[i] {
			maxCount = count[i]
			maxChar = rune(i + 'a')
		}
	}

	fmt.Printf("maxChar: %c, maxCount: %d\n", maxChar, maxCount)
	fmt.Println(count)

	for i := 1; i < len(count); i++ {
		count[i] += count[i-1]
	}

	fmt.Println(count)

	sorted := make([]rune, len(str))
	for i := 0; i < len(str); i++ {
		c := rune(str[i])
		countIndex := c - 'a'
		sorted[count[countIndex]-1] = c
		count[countIndex] -= 1
	}
	fmt.Println(sorted)

	for _, n := range sorted {
		fmt.Printf("%c ", n)
	}
}
