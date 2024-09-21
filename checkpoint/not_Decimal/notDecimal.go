package main

import (
	"fmt"
)

func main(){
		fmt.Print(NotDecimal("0.1"))
		fmt.Print(NotDecimal("174.2"))
		fmt.Print(NotDecimal("0.1255"))
		fmt.Print(NotDecimal("1.20525856"))
		fmt.Print(NotDecimal("-0.0f00d00"))
		fmt.Print(NotDecimal(""))
		fmt.Print(NotDecimal(" "))
		fmt.Print(NotDecimal("-19.525856"))
		fmt.Print(NotDecimal("1952"))
	
}

func NotDecimal(dec string) string {
	result := ""

	if dec == ""{
		return "\n"
	}
	decCopy := dec
	newStr := ""
isNegative := false
	for _, ch := range dec {
		if ch == '-'{
			isNegative = true
			continue
		} else {
			newStr += string(ch)
		}
	}
dec = newStr
nStr := ""
	for _, ch := range dec {
		if ch == '.'{
			continue
		} else {
			nStr += string(ch)
		}
	}

	dec = nStr
	res := ""
	for i, ch := range dec {
		if dec[i] == '0' && res == "" {
			continue
		} else {
			res += string(ch)
		}
	}

	for _,ch := range res{
		if ch < '0' || ch > '9' {
			return decCopy + "\n"
		} else {
			result += string(ch)
		}
	}

	if isNegative{
		return "-" + result + "\n"	
	}
	return result + "\n"
}