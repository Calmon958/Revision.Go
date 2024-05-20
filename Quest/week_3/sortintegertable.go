package main

//import "fmt"

func SortIntegerTable(table []int) {
	counter := 0
	for range table {
		counter++
	}
	for a := 0; a < counter-1; a++ {
		for b := 0; b < (counter - a - 1); b++ {
			table[b+1], table[b] = table[b], table[b+1]
		}
	}
}

// func main() {
// 	s := []int{5, 4, 3, 2, 1, 0}
// 	SortIntegerTable(s)
// 	fmt.Println(s)
// }
