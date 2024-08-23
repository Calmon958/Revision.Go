package main

import "fmt"

func Strlen(s string) int {
count := 0
for range s {
count++
}
return count
}

func main() {
fmt.Println(Strlen("Hello World!"))
}
