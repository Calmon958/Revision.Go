package main

import "fmt"

func Save(arg string, num int) []string {
	chunk := []string{}
	str := ""
	count := 0
	count2 := 0
	for i := 0; i < len(arg); i++ {
		count++
		if count <= 3 {
			str+=string(str[i])
		} else if count2 < 3 {
			continue
		} else{
			count = 0
			count2 = 0
		}
	}
	return chunk
}

func main() {
	fmt.Println(Save("12345789", 3))
}
