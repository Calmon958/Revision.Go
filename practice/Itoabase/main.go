package main

import "fmt"

func ItoaBase(value, base int) string {
if base < 2 || base > 16 {
return ""
}
if  value  == 0 {
return "0"
}
digit := "0123456789ABCDEF"
res := ""
isNegative := value < 0
if isNegative {
value=-value
}
for value >0 {
reminder := value%base
res =string(digit[reminder]) + res
value/=base
}
if isNegative {
res = "-" + res
}
return res
}


func main() {
	fmt.Println(ItoaBase(255, 16))
	fmt.Println(ItoaBase(-255, 16))
	fmt.Println(ItoaBase(10, 2))
	fmt.Println(ItoaBase(-10, 2))
	fmt.Println(ItoaBase(0, 10))
}
