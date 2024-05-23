package main

import "fmt"

func Chunk(slice []int, size int) {
if size == 0 {
return
}



answer := [][]int{}
//endRune:=0
for i:=0; i<len(slice); i+=size {
	endRune := i + size
if endRune > len(slice){
endRune = len(slice)
}
answer = append(answer, slice[i:endRune])
}
fmt.Println(answer)
}



func main() {
	Chunk([]int{}, 10)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 0)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 3)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 5)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 4)
}


