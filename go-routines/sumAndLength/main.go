package main

import(
	"fmt"
	// "channel"
)

func main() {
	//channels are passed before the data type using chan
c := make(chan int) //integer channel
b:= make(chan int) //integer channel
// a := make(chan string)
go strLen("Hello", c)
go strLen("Golden era", b)
x, y := <-c, <-b
fmt.Println(x, y, x+y)
}

func strLen(s string, b chan int){
	b <- len(s)
}