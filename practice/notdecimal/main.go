package main

import (
	"fmt"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/solutions"
)

func NotDecimal(dec string) string {
	decBackup := dec
	result := ""
	if dec == "" {
		return "\n"
	}

	isNegative := false
	if dec[0] == '-' {
		isNegative = true
		if len(dec) > 1 {
			dec = dec[1:]
		}
	}

	// remove decimal point
	for _, char := range dec {
		if char == '.' {
			continue
		}
		result += string(char)
	}

	dec = result
	result = ""
	// remove zero before the number
	for i, c := range dec {
		if c == '0' {
			continue
		} else {
			result = dec[i:]
			break
		}
	}

	dec = result
	result = ""
	// print numbers and not characters
	for _, c := range dec {
		if c >= '0' && c <= '9' {
			result += string(c)
		} else {
			return decBackup + "\n"
		}
	}

	if isNegative {
		result = "-" + result
	}

	return result + "\n"
}

func main() {
	fmt.Print(NotDecimal("0.1"))
	fmt.Print(NotDecimal("174.2"))
	fmt.Print(NotDecimal("0.1255"))
	fmt.Print(NotDecimal("1.20525856"))
	fmt.Print(NotDecimal("-0.0f00d00"))
	fmt.Print(NotDecimal(""))
	fmt.Print(NotDecimal("-19.525856"))
	fmt.Print(NotDecimal("1952"))

	fmt.Println("-----------------GO TESTS-----------------")
	args := []string{"-19.525856", "00.02", "56s44", "", "415.458", "1.6", "165", "502.3254", "51.3+95.9", "-5.0f00d00", "-53.124"}

	for i := 0; i < len(args); i++ {
		challenge.Function("NotDecimal", NotDecimal, solutions.NotDecimal, args[i])
	}
}
