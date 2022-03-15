package main

func sumNums(n int) int {
	if n == 1 {
		return 1
	}
	result := n + sumNums(n-1)
	return result

}