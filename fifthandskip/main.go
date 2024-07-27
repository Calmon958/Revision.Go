package main

import "fmt"

func FifthAndSkip(str string) string {
word:=""

for i:=0; i<len(str); i++ {
	if (str[i] >= 'a' && str[i]<='z') || (str[i]>='a' && str[i]<='z') {
		if (i+1)%5 != 0 {
			word[i]+=str[i]
		} else {
			word[i] = ' '
		}
	}
}
	return "Invalid Input"
}



func main() {
	fmt.Println(FifthAndSkip("abcdefghijklmnopqrstuwxyz"))
	fmt.Println(FifthAndSkip("This is a short sentence"))
	fmt.Println(FifthAndSkip("1234"))
}


