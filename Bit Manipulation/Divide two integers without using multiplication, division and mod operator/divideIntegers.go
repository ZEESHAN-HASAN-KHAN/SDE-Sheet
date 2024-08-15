/*
	Given two integers dividend and divisor, divide two integers without using multiplication, division, and mod operator.

The integer division should truncate toward zero, which means losing its fractional part. For example, 8.345 would be truncated to 8, and -2.7335 would be truncated to -2.

Return the quotient after dividing dividend by divisor.

Note: Assume we are dealing with an environment that could only store integers within the 32-bit signed integer range: [−231, 231 − 1]. For this problem, if the quotient is strictly greater than 231 - 1, then return 231 - 1, and if the quotient is strictly less than -231, then return -231.

Example 1:

Input: dividend = 10, divisor = 3
Output: 3
Explanation: 10/3 = 3.33333.. which is truncated to 3.
Example 2:

Input: dividend = 7, divisor = -3
Output: -2
Explanation: 7/-3 = -2.33333.. which is truncated to -2.

Constraints:

-231 <= dividend, divisor <= 231 - 1
divisor != 0
*/
package main

import "math"

func divide(dividend int, divisor int) int {
	if dividend == divisor {
		return 1
	}

	// Handle overflow case
	if dividend == math.MaxInt32 && divisor == -1 {
		return math.MinInt32 + 1
	}
	if dividend == -math.MaxInt32 && divisor == -1 {
		return math.MaxInt32
	}

	sign := true
	if dividend < 0 && divisor > 0 {
		sign = false
	} else if dividend > 0 && divisor < 0 {
		sign = false
	}

	div := int(math.Abs(float64(dividend)))
	dis := int(math.Abs(float64(divisor)))
	ans := 0
	for div >= dis {
		count := 0
		for div >= (dis << (count + 1)) {
			count++
		}

		ans += (1 << count)
		div -= (dis << count)
	}

	// Handle overflow
	if ans >= math.MaxInt32 {
		if sign {
			return math.MaxInt32
		}
		return math.MinInt32
	}

	if sign {
		return ans
	}

	return -ans
}

func main() {
	println(divide(10, 3)) // Output: 3
	println(divide(7, -3)) // Output: -2
	println(divide(1, 1)) // Output: 1
}