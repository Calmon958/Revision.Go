package main

import "fmt"

func Replace(s, old, new string) string {
	if old == "" || old==new {
		return s
	}

	res := ""
	n := len(old)
	for i:=0; i<len(s); {
		if len(s[i:]) >= n && s[i:i+n] == old {
			res += new
			i+=n
		} else {
			res += string(s[i])
			i++
		}
	}
	fmt.Println(s)
	return res
}

func main(){
	fmt.Println(Replace("hello world", "world", "python"))
}