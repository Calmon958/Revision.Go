package main

import "fmt"

func BubbleSort2(arr []int, compute func(int, int) bool) {
	size := len(arr)
	swapped := 1
	for i := 0; i < (size-1) && swapped == 1; i++ {
		swapped = 0
		for j := 0; j < size-i-1; j++ {
			if compute(arr[j], arr[j+1]) {
				arr[j+1], arr[j] = arr[j], arr[j+1]
				swapped = 1
			}
		}
	}
}

func main() {
	// more := true
	arr := []int{64, 34, 25, 12, 22}
	BubbleSort2(arr, more)
	fmt.Println(arr)
}

//insertion sorting

func insertionSort(arr []int, compute func(int, int) bool) {
	size := len(arr)
	var temp, i,j int
	for i = 1;i<size;i++{
		temporary = arr[i]
	}
}
