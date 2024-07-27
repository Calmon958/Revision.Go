package main

import (
	"fmt"
	"os"
)

func main() {
	num := os.Args[1]
	n := atoi(num)
	fmt.Println(addPrime(n))
}

func Prime(nb int) bool {
	if nb < 2 {
		return false
	}

	for i := 2; i<nb; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}

func addPrime(nb int) int {
	sum := 0
	for i := 2; i <=nb; i++ {
		if Prime(i) {
			sum += i
		}
	}
	return sum
}

func atoi(s string) int {
	nb := 0
	for _, char := range s {
		nb = nb*10 + int(char-'0')
	}
	return nb
}
