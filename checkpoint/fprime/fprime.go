package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	arg := os.Args[1]

	b, _ := strconv.Atoi(arg)

	if isPrime(b) && primeFactor(b) == "" {
		fmt.Println(b)
	} else {
		fmt.Println(primeFactor(b))
	}
}

func isPrime(nb int) bool {
	// fmt.Println("getting: ", nb)
	for i := 2; i*i < nb; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}

func primeFactor(nb int) string {
	result := ""

	for i := 2; i <= nb; i++ {
		for nb%i == 0 {
			if result != "" {
				result += "*"
			}
			result += Itoa(i)
			nb /= i
		}
	}
	return result
}

func Itoa(nb int) string {
	str := ""
	for nb > 0 {
		digit := nb % 10
		str = string(digit+'0') + str
		nb /= 10
	}
	return str
}
