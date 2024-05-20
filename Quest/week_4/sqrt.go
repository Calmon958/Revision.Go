package main

// import "fmt"

func Sqrt(nb int) int {
	square := 0
	for a := 0; a*a <= nb; a++ {
		if a*a == nb {
			return a
		}
	}
	return square
}

// func main() {
// 	fmt.Println(Sqrt(4))
// }
