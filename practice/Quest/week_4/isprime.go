package main

// import "fmt"

func IsPrime(nb int) bool {
	if nb < 2 {
		return false
	} else if nb ==2 || nb == 3 {
		return true
	} else if (nb%2 == 0 && nb > 2) || (nb%3 == 0 && nb >3) {
		return false
	}
	for i := 3; i*i <=nb; i++ {
		if i*i == nb {
			return false
		}
	}
	return true
}

// func main() {
// 	fmt.Println(IsPrime(49))
// }