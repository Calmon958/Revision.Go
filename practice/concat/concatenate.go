package main

func Concatenate(str []string) string {
	res := ""
	for _, ch := range str {
		res += ch
	}
	return res
}
