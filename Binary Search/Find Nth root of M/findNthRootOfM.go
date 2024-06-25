/* You are given 2 numbers (n , m); the task is to find n√m (nth root of m).
 

Example 1:

Input: n = 2, m = 9
Output: 3
Explanation: 32 = 9
Example 2:

Input: n = 3, m = 9
Output: -1
Explanation: 3rd root of 9 is not
integer.
 

Your Task:
You don't need to read or print anyhting. Your task is to complete the function NthRoot() which takes n and m as input parameter and returns the nth root of m. If the root is not integer then returns -1.
 

Expected Time Complexity: O(n* log(m))
Expected Space Complexity: O(1)
 

Constraints:
1 <= n <= 30
1 <= m <= 109
*/
package main

import "fmt"

func multiply(num int, n int, m int) int {
	result:= 1
	for i:=0; i<n; i++ {
		result *= num
		if result > m {
			return 2
		}
	}

	if result == m {
		return 1
	}

	return 0
}

func NthRoot(n int, m int) int {
	var (
		low = 1
		high = m
	)

	for low <= high {
		mid := (low+high)/2
		result := multiply(mid, n, m)
		if result == 1 {
			return mid
		} else if result == 2 {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}

func main() {
	fmt.Println(NthRoot(2, 9)) // Output: 3
	fmt.Println(NthRoot(3, 9)) // Output: -1
}