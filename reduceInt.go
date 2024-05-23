package main
import "fmt"

func ReduceInt(a []int, f func(int, int) int) {
answer := a[0]
for i := 1; i < len(a); i++ {
answer = f(answer, a[i])
//z01.PrintRune(char)
}
fmt.Println(answer)
}

func main() {
	mul := func(acc int, cur int) int {
		return acc * cur
	}
	sum := func(acc int, cur int) int {
		return acc + cur
	}
	div := func(acc int, cur int) int {
		return acc / cur
	}
	as := []int{500, 20, 10, 5}
	ReduceInt(as, mul)
	ReduceInt(as, sum)
	ReduceInt(as, div)
}
