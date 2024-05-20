package main

// import "fmt"

func BasicAtoi(s string) int {
	res := 0
	for _, value := range s {
		a := int(value - '0')
		res = res*10 + a
	}
	return res
}

// func main() {
// 	fmt.Println(BasicAtoi("12345"))
// 	fmt.Println(BasicAtoi("0000000012345"))
// 	fmt.Println(BasicAtoi("000000"))
// }
