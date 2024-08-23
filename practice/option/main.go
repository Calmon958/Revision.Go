package main

import (
	"fmt"
	"os"
	// "uint"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Print("option: -abcdefghijklmnopqrstuvwxyz")
		return
	}
	arg := os.Args[1:]
	Option(arg)
}

func Option(str []string) {
	var optionInt uint32
	for _, arg := range str {
		if len(arg) > 0 && arg[0] == '-' {
			if len(arg) > 1 && Contain(arg) {
				fmt.Print("option: -abcdefghijklmnopqrstuvwxyz")
				return
			}
			if arg[0] == '-' && len(arg) == 1 {
				fmt.Println("Invalid option")
				return
			}
			for _, ch := range arg[1:] {
				if ch >= 'a' && ch <= 'z' {
					bitPosition := ch - 'a'
					optionInt |= (1 << bitPosition)
				} else {
					fmt.Println("Invalid option")
					return
				}
			}
		} else {
			fmt.Println("Invalid option")
			return
		}
	}
	bitString := Binary(optionInt, 32)
	for i := 0; i < 32; i += 8 {
		if (i + 8) < 32 {
			fmt.Print(bitString[i:i+8] + " ")
			// return
		} else {
			fmt.Print(bitString[i:i+8])
			// return
		}
	}
	fmt.Println()
}

func Contain(str string) bool {
	return str[:2] == "-h"
}

func Binary(nbr uint32, size int) string {
	binary := make([]byte, size)
	for i := 0; i < size; i++ {
		if nbr&(1<<(size-i-1)) != 0 {
			binary[i] = '1'
		} else {
			binary[i] = '0'
		}
	}
	return string(binary)
}
