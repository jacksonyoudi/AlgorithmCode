package bit


const (
	one = 0xaaaaaaaa // 10101

)


// x > 0 and x & (x - 1) == 0
func isPowerOfFour(n int) bool {

	return (n > 0) && (n & (n-1) == 0) && (n & one != 0)

}