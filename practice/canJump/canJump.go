package main

func CanJump(nb []uint) bool{
	End := 0
	n := len(nb)

	
	if n == 1 {
		return true
	} else if n == 0{
		return false
	}


	for i:=0; i<n; i++ {
		if i > End {
			return false
		}
		if i+int(nb[i]) > End {
			End = i+int(nb[i])
		}
		if End >= n-1 {
			return true
		}
	}
	return false
}