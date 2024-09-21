package main

import (
	"fmt"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/solutions"
)

func NotDecimal(dec string) string {
	decCopy := dec
result := ""
	for _, ch := range dec {
		if ch == '.' {
			continue
		} else {
			result += string(ch)
		}
	}
dec = result
	for i, ch := range dec {
		if ch == '0'{
			continue
		} else {
			dec = dec[i:]
			break
		}
	}
finDec := dec
res := ""
	for _, ch := range finDec{
		if ch>='0' && ch<='9' {
			res += string(ch)
		} else {
			return decCopy
		}
	}
return res
}

func main() {
	fmt.Println(NotDecimal("0.1"))
	fmt.Println(NotDecimal("174.2"))
	fmt.Println(NotDecimal("0.1255"))
	fmt.Println(NotDecimal("1.20525856"))
	fmt.Print(NotDecimal("0.0f00d00"))
	fmt.Println(NotDecimal(""))
	// fmt.Print(NotDecimal("-19.525856"))
	fmt.Println(NotDecimal("1952"))

	fmt.Println("-----------------GO TESTS-----------------")
	args := []string{"-19.525856", "00.02", "56s44", "", "415.458", "1.6", "165", "502.3254", "51.3+95.9", "-5.0f00d00", "-53.124"}

	for i := 0; i < len(args); i++ {
		challenge.Function("NotDecimal", NotDecimal, solutions.NotDecimal, args[i])
	}
}
