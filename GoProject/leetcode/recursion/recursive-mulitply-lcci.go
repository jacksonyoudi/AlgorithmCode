package main

func multiply(A int, B int) int {
	if A == 0 || B == 0 {
		return 0
	}

	if A == 1 || B == 1 {
		return A + B - 1
	}

	if A > B {
		return B + multiply(A, B-1)
	} else {
		return A + multiply(A - 1 , B)
	}

}

