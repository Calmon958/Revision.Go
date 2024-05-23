package main

import (
	"fmt"
	"os"
)

func main() {

	args := os.Args[1]

	str := ""
	for i := len(args) - 1; i >= 0; i-- {
		if args[i] != ' ' {
			str = string(args[i]) + str
		} else {
			if str != "" {
				break
			}

		}
	}
	fmt.Println(str)

}

// s := ''

// s := ''
// for _,char := range args {
// 	for i := len(args) -1; i >= 0; i-- {
// 		if char[i] == ' '
// 	}

// }
// }
// for _,char := range args {
// 	for i := len(args) -1; i >= 0; i-- {
// 		if char[i] == ' '
// 	}

// }
// }
