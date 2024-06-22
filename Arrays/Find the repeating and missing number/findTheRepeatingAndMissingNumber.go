/*
	Given an unsorted array arr of size n of positive integers. One number 'A' from set {1, 2,....,N} is missing and one number 'B' occurs twice in array. Find these two numbers.

Your task is to complete the function findTwoElement() which takes the array of integers arr and n as parameters and returns an array of integers of size 2 denoting the answer (The first index contains B and second index contains A)

# Examples

Input: n = 2 arr[] = {2, 2}
Output: 2 1
Explanation: Repeating number is 2 and smallest positive missing number is 1.
Input: n = 3 arr[] = {1, 3, 3}
Output: 3 2
Explanation: Repeating number is 3 and smallest positive missing number is 2.
Expected Time Complexity: O(n)
Expected Auxiliary Space: O(1)

Constraints:
2 ≤ n ≤ 105
1 ≤ arr[i] ≤ n
*/
package main

import (
	"fmt"
)

func findTwoElement(arr []int, n int) (int, int) {
	sumA := int64(n * (n + 1) / 2)
	sumR := int64(0)
	for _, i := range arr {
		sumR += int64(i)
	}

	powerA := int64(n * (n + 1) * (2*n + 1) / 6)
	powerR := int64(0)

	for _, i := range arr {
		powerR += int64(i) * int64(i)
	}

	eq1 := sumA - sumR             // X - Y
	eq2 := (powerA - powerR) / eq1 // X + Y

	X := int((eq1 + eq2) / 2)
	Y := int(eq2 - int64(X))

	return Y, X
}

func main() {
	arr := []int{4, 3, 6, 2, 1, 1}
	n := 6
	repeating, missing := findTwoElement(arr, n)
	fmt.Printf("Missing number: %d, Repeating number: %d\n", missing, repeating)
}
