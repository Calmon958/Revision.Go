package main

import (
	"fmt"
)

func main() {
	fmt.Print(WordFlip("First second last"))
	// fmt.Print(WordFlip(""))
	// fmt.Print(WordFlip("     "))
	// fmt.Print(WordFlip(" hello  all  of  you! "))
}

func WordFlip(str string) string {
	if str == "" {
		return "Invalid Option"
	}

	result := ""
	res := []string{}
	strSlice := Split(str, '\n')
	for i := len(strSlice) - 1; i >= 0; i-- {
		res = append(res, strSlice[i])
	}
	result = Join(res, '\n')
	return result
}

func Split(str string, r rune) []string {
	ans := ""
	res := []string{}
	for _, ch := range str {
		if ch == ' ' && ans == "" {
			continue
		} else {
			if ch != ' ' {
				ans += string(ch)
			} else {
				res = append(res, ans)
				ans = ""
			}
		}
	}
	return res
}

func Join(strSlice []string, rb rune) string {
	ans := ""
	rb = '\n'
	for i := range strSlice {
		if i != len(strSlice)-1 {
			ans += strSlice[i] + string(rb)
		} else {
			ans += strSlice[i]
		}
	}

	return ans
}
