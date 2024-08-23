package main

// import "fmt"

func FindNextPrime(nb int) int {
	if nb <= 2 {
		return 0
	}
	for b := nb; ; b++ {
		if IsPrim(b) {
			return b
		}
	}
}

func IsPrim(n int) bool {
	if n < 2 {
		return false
	} else if n == 2 {
		return true
	} else if n%2 == 0 && n > 2 {
		return false
	}
	for a := 3; a*a <= n; a++ {
		if n%a == 0 {
			return false
		}
	}
	return true
}

// func main() {
// 	fmt.Println(FindNextPrime(14))
// 	fmt.Println(FindNextPrime(13))
// }
