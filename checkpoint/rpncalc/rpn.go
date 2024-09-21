package main

import (
	"strconv"
)

// func main() {
// 	if len(os.Args) != 2 {
// 		return
// 	}

// 	str := os.Args[1]
// }

func rpn(str string) int {
	res := ""
	var nb int
	for _, ch := range str {
		if ch == ' ' {
			continue
		} else {
			res += string(ch)
		}
	}
	for i, ch := range res {
		switch ch {
		case '*':
			nb1, _ := strconv.Atoi(string(res[i-1]))
			nb2, _ := strconv.Atoi(string(res[i-2]))
			nb = nb1 * nb2
		case '+':
			nb1, _ := strconv.Atoi(string(res[i-1]))
			nb2, _ := strconv.Atoi(string(res[i-2]))
			nb = nb1 + nb2
		case '-':
			nb1, _ := strconv.Atoi(string(res[i-1]))
			nb2, _ := strconv.Atoi(string(res[i-2]))
			nb = nb1 - nb2
		}
	}
	return nb
}
