package main

import "fmt"

func CountIf(f func(string) bool, tab []string) int {
count:=0
for _, char := range tab {
if f(char) {
count++
}
}
return count
}


func main() {
	tab1 := []string{"Hello", "how", "are", "you"}
	tab2 := []string{"This","1", "is", "4", "you"}
	answer1 :=CountIf(IsNumeric, tab1)
	answer2 :=CountIf(IsNumeric, tab2)
	fmt.Println(answer1)
	fmt.Println(answer2)
}



func IsNumeric(s string) bool {
	for _, r := range s {
		if r < '0' || r > '9' {
			return false
		}
	}
	return true
}
