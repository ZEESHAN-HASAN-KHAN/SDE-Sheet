/* You are given two integers L and R, your task is to find the XOR of elements of the range [L, R].

Example:

Input: 
L = 4, R = 8 
Output:
8 
Explanation:
4 ^ 5 ^ 6 ^ 7 ^ 8 = 8
Your Task:

Your task is to complete the function findXOR() which takes two integers l and r and returns the XOR of numbers from l to r.

Expected Time Complexity: O(1).

Expected Auxiliary Space: O(1).

Constraints:

1<=l<=r<=109

*/
package main

import (
	"fmt"
)

func nThXor(n int) int {
	switch n % 4 {
	case 0:
		return n
	case 1:
		return 1
	case 2:
		return n + 1
	case 3:
		return 0
	}
	return 0
}

func findXOR(l int, r int) int {
	return nThXor(l-1) ^ nThXor(r)
}

func main() {
	fmt.Println(findXOR(4, 8)) // 8
}