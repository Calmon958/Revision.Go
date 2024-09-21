package main

import "fmt"

func FifthAndSkip(str string) string {
	word := ""
	count := 0

	if str == "" {
		return "\n"
	}

	for _, ch := range str {
		if ch != ' ' {
			if (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') {
				count++
				if count%6 == 0 {
					word += string(' ')
				} else {
					word += string(ch)
				}
			} else {
				return "Invalid Input"
			}
		} else {
			continue
		}
	}
	return word
}

func main() {
	fmt.Println(FifthAndSkip("abcdefghijklmnopqrstuwxyz"))
	fmt.Println(FifthAndSkip("This is a short sentence"))
	fmt.Println(FifthAndSkip("1234"))
	fmt.Println(FifthAndSkip("Po65 4o"))
	fmt.Println(FifthAndSkip("w58tw7474abc"))
	fmt.Println(FifthAndSkip("-552"))
	fmt.Println(FifthAndSkip("8595485-52"))
	fmt.Println(FifthAndSkip("Kimetsu no Yaiba"))
	fmt.Println(FifthAndSkip("hello \\! n4ght cr3a8ure7 "))
	fmt.Println(FifthAndSkip("QKplq%QSw"))
	fmt.Println(FifthAndSkip("e 5£ @ 8* 7 =56 ;"))
	fmt.Println(FifthAndSkip("1234556789"))
}
