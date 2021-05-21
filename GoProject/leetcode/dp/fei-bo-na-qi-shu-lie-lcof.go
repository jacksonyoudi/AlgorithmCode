package dp

func fib(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 || n == 2 {
		return n
	}

	i := 3
	a1 := 1
	a2 := 2
	a3 := a1 + a2

	for i <= n {
		a1 = a2
		a2 = a3
		a3 = a1 + a2
		i++
	}

	return a3

}
