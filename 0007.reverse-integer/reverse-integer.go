package leetcode

import "math"

func reverse(x int) int {
	sign := 1

	// handle negative numbers
	if x < 0 {
		sign = -1
		x = -1 * x
	}

	res := 0
	for x > 0 {
		// Take the end of x
		temp := x % 10
		// put at the beginning of res
		res = res*10 + temp
		// x remove the end
		x = x / 10
	}

	// restore the symbol of x to res
	res = sign * res

	// deal with res overflow problem
	if res > math.MaxInt32 || res < math.MinInt32 {
		res = 0
	}

	return res
}
