package main

import "fmt"

func Slice(a []string, nbrs ...int) []string {
	start := 0
	end := 0

	if len(nbrs) == 1 {
		return a[start:]
	}
	if len(nbrs) > 1 {
		start = nbrs[0]
		end = nbrs[1]
	}

	if start < 0 {
		start = len(nbrs) + start
	}
	if end < 0 {
		end = len(nbrs) + end
	}

	if start >= len(a) || start >= end {
		return nil
	}//had forgoten this part
	return a[start:end]
}

func main() {
	a := []string{"coding", "algorithm", "ascii", "package", "golang"}
	fmt.Printf("%#v\n", Slice(a, 1))
	fmt.Printf("%#v\n", Slice(a, 2, 4))
	fmt.Printf("%#v\n", Slice(a, -3))
	fmt.Printf("%#v\n", Slice(a, -2))
	fmt.Printf("%#v\n", Slice(a, 2, 0))
}
