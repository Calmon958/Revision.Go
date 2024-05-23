package main

import "fmt"

func Chunk(slice []int, size int) {
	if size == 0 {
		return
	}
	var result [][]int
	// length := 0
	for i := 0; i < len(slice); i += size {
		endRune := i + size
		if endRune > len(slice) {
			endRune = len(slice)
		}
		result = append(result, slice[i:endRune])
	}

	fmt.Println(result)
}

func main() {
	Chunk([]int{}, 10)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 0)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 3)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 5)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 4)
}
