package countingSort

import "fmt"

func SortHeight() {

	structs := []struct {
		Name   string
		Height float64
	}{
		{
			Name:   "Kyle",
			Height: 173.4,
		},
		{
			Name:   "Ken",
			Height: 164.5,
		},
		{
			Name:   "Ryu",
			Height: 178.8,
		},
		{
			Name:   "Honda",
			Height: 154.2,
		},
		{
			Name:   "Hwarang",
			Height: 188.8,
		},
		{
			Name:   "Lebron",
			Height: 209.8,
		},
		{
			Name:   "Hodong",
			Height: 197.7,
		},
		{
			Name:   "Tom",
			Height: 164.8,
		},
		{
			Name:   "Kevin",
			Height: 164.8,
		},
	}

	var heightMap [3000][]string

	for i := 0; i < len(structs); i++ {
		s := structs[i]
		intHeight := int(s.Height * 10)
		name := s.Name

		heightMap[intHeight] = append(heightMap[intHeight], name)
	}

	min := 154.3
	max := 190.0

	intMin := int(min * 10)
	intMax := int(max * 10)

	for i := intMin; i <= intMax; i++ {
		names := heightMap[i]
		for _, name := range names {
			height := float64(i) / 10
			fmt.Printf("%s: %.1f\n", name, height)
		}
	}
}
