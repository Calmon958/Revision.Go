package main

import "fmt"

func rev(oct byte) byte {
	var result byte
	for i:=0; i<8; i++ {
		
		result = (result << 1 | oct & 1)
		oct >>= 1
	}
return result
}

func main() {
	// args := []byte{0x26, 0x27, 0x28, 0x29, 0xAA, 0xBB}

	fmt.Println(rev(0x26))
}
