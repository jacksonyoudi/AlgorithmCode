package prime_number_of_set_bits_in_binary_representation

func countPrimeSetBits(left int, right int) int {
	m := 0
	for j := left; j <= right; j++ {
		if isPrime(hammingWeight(j)) {
			m++
		}
	}
	return m
}

func isPrime(x int) bool {
	// 0 和 1 都是
	if x < 2 {
		return false
	}
	for i := 2; i*i <= x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

func hammingWeight(n int) int {
	i := 0
	for ; n > 0; n &= (n - 1) {
		i++
	}
	return i
}
