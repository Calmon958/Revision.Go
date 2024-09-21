package main

import "fmt"

func ConcatAlter(str1, str2 []int) []int {
	if len(str1) < len(str2) {
		str1, str2 = str2, str1
	}

	res := []int{}
	for i, nb := range str1 {
		res = append(res, nb)
		if i < len(str2) {
			res = append(res, str2[i])
		}
	}
	return res
}

func main() {
	fmt.Println(ConcatAlter([]int{2, 4, 6, 8, 10}, []int{1, 3, 5, 7, 9, 11}))
	fmt.Println(ConcatAlter([]int{1, 2, 3}, []int{4, 5, 6, 7, 8, 9}))
	fmt.Println(ConcatAlter([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(ConcatAlter([]int{1, 2, 3}, []int{}))
}
