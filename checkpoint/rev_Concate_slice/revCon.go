package main

import (
	"fmt"
)


func RevConcatAlternate(slice1, slice2 []int) []int {
	len1 := len(slice1)
	len2 := len(slice2)
	fin := []int{}

	if len1 < len2 {
		slice1, slice2 = slice2, slice1
		len1, len2 = len2, len1
	}

	for i:= len1 - 1; i>=0; i-- {
		fin = append(fin, slice1[i])
		if i < len2 {
			fin = append(fin, slice2[i])
		}
	}
	return fin
}

func main() {
	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6, 7, 8, 9}))
	fmt.Println(RevConcatAlternate([]int{1, 2, 3, 9, 8}, []int{4, 5}))
	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{}))
}
