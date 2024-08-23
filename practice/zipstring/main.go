package main

import (
//"github.com/01-edu/z01"
"fmt"
)

func ZipString(s string) string {
if s == "" {
return ""
}
result := ""
count:=1

for i:=1; i<len(s); i++ {
if s[i]==s[i-1]{
count++
} else {
count = 1
result +=Itoa(count) + string(s[i-1])
}
}
result +=Itoa(count) + string(s[len(s)-1])
return result
}

func Itoa(nb int) string {
result := ""
for nb > 0 {
digit := nb%10
result += string(digit + '0') + result
nb/=10
}
return result
}

//package main

//import (
//	"fmt"
//)

func main() {
	args := []string{"aaaa","zzzzzZZZZZZ","","ziiiiipMeee","hhellloTthereYouuunggFelllas"}

	for _, arg := range args {
		fmt.Println(ZipString(arg))
	}
}
