package main

import "fmt"

func HashCode(dec string) string {
size := len(dec)
hashed := make([]rune, size)

for i, ch := range dec {
newChar := (int(ch)+size)%127
if newChar < 33 || newChar > 126 {
newChar+=33
}
hashed[i]=rune(newChar)
}
return string(hashed)
}


func main() {
	fmt.Println(HashCode("A"))
	fmt.Println(HashCode("AB"))
	fmt.Println(HashCode("BAC"))
	fmt.Println(HashCode("Hello World"))
}

