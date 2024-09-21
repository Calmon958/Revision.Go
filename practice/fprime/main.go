package main

import (
	"fmt"
	"os"
)

func main() {
	num := os.Args[1]
	numb := atoi(num)
	fmt.Println(primeFactor(numb))
}

func atoi(str string) int {
	nb := 0
	for _, char := range str {
		for i := '0'; i <= '9'; i++ {
			nb = nb*10 + int(char-'0')
		}
	}
	return nb
}

func itoa(nb int) string {
	str := ""
	for nb > 0 {
		digit := nb % 10
		str = itoa(digit+'0') + str
		nb /= 10
	}
	return str
}

func Prime(nb int) bool {
	for i := 2; i < nb; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}

func primeFactor(num int) string {
	result := ""
	for i := 2; i < num; i++ {
		for Prime(i) {
			if num%i == 0 {
				result = itoa(i)
				num /= i
			}
		}
	}
	return result
}

// package main

// import (
// 	"fmt"
// 	"os"
// 	"strconv"
// )

// func main() {
// 	if len(os.Args) < 2 {
// 		fmt.Println("Please provide a number as an argument.")
// 		return
// 	}

// 	num, err := strconv.Atoi(os.Args[1])
// 	if err != nil {
// 		fmt.Println("Invalid number:", os.Args[1])
// 		return
// 	}

// 	fmt.Println(primeFactor(num))
// }

// func primeFactor(num int) string {
// 	result := ""
// 	for i := 2; i <= num; i++ {
// 		for num%i == 0 {
// 			result += strconv.Itoa(i) + " "
// 			num /= i
// 		}
// 	}
// 	return result
// }

// func Prime(nb int) bool {
// 	if nb <= 1 {
// 		return false
// 	}
// 	for i := 2; i*i <= nb; i++ {
// 		if nb%i == 0 {
// 			return false
// 		}
// 	}
// 	return true
// }
