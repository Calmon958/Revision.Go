package main

import "fmt"

func Atoi(s string) int {
	sign := 1
	count := 0
	a := 0
	for i, value := range s {
		if value == '-' || value == '+' {
			if i != 0 {
				return 0
			} else if i == 0 && value == '-' {
				sign = -1
			}
			count++
		} else if value < '0' || value > '9' || count > 1 {
			return 0
		}
		b := 0
		for c := '1'; c <= value; c++ {
			b++
		}
		a = a*10 + b
	}
	return sign * a
}

func main() {
	fmt.Println(Atoi("12345"))
	fmt.Println(Atoi("0000000012345"))
	fmt.Println(Atoi("012 345"))
	fmt.Println(Atoi("Hello World!"))
	fmt.Println(Atoi("+1234"))
	fmt.Println(Atoi("-1234"))
	fmt.Println(Atoi("++1234"))
	fmt.Println(Atoi("--1234"))
}
