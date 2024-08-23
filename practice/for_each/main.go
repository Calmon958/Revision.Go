package main


import (
	"fmt"
//	"math/rand"
//	"time"
)


func ForEach(f func(int), a []int) {
	for _, nb := range a {
		f(nb)
	}
}


/*
func IntBetween(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

func IntSliceBetween(min, max, length int) []int {
	rand.Seed(time.Now().UnixNano())
	slice := make([]int, length)
	for i := range slice {
		slice[i] = IntBetween(min, max)
	}
	return slice
}


func main() {
	functions := []func(int){add0, add1, add2, add3}

	type node struct {
		f func(int)
		a []int
	}

	table := []node{{
		f: add0,
		a: []int{1, 2, 3, 4, 5, 6},
	}}

	// 15 random slice of random ints with a random function from selection
	// for i := 0; i < 15; i++ {
		function := functions[IntBetween(0, len(functions)-1)]
		table = append(table, node{
			f: function,
			a: IntSliceBetween(-1000000, 1000000, 15),
		})
	// }

	for _, arg := range table {
		fmt.Println(ForEach(arg.f, arg.a))
	}
}

func add0(i int) {
	fmt.Println(i)
}

func add1(i int) {
	fmt.Println(i + 1)
}

func add2(i int) {
	fmt.Println(i + 2)
}
*/
func add3(i int) {
	fmt.Println(i + 3)
//fmt.Print(",")
}

func  PrintNbr(i int) {
fmt.Println(i)
}
//fmt.Println()
func main() {
	a := []int{1, 2, 3, 4, 5, 6}
ForEach(PrintNbr, a)

	ForEach(add3, a)
}

