package main

func DivMod(a, b int, div *int, mod *int) {
	*div = a / b
	*mod = a % b
}
