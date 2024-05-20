package main

//import "fmt"

func IterativeFactorial(nb int) int {
	result := 1
	if nb == 1 || nb == 0 {
		return 0
	} else if nb > 1 && nb <= 30 {
		for a := 1; a <= nb; a++ {
			result = result * a
		}
	} else {
		return 0
	}
	return result
}

// func main() {
// 	arg := 4
// 	fmt.Println(IterativeFactorial(arg))
// }
