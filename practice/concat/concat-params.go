package main

func ConcatParams(args []string) string {
	res := ""
	for _, ch := range args {
		if ch != args[len(args)-1] {
			res += ch + "\n"
		} else {
			res += ch
		}
	}
	return res
}

// func main() {
// 	fmt.Println(ConcatParams([]string{"a", "b"}))
// }
