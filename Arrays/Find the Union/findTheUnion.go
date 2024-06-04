/*
Given two sorted arrays of size n and m respectively, find their union. The Union of two arrays can be defined as the common and distinct elements in the two arrays. Return the elements in sorted order.

Example 1:

Input:
n = 5, arr1[] = {1, 2, 3, 4, 5}
m = 5, arr2 [] = {1, 2, 3, 6, 7}
Output:
1 2 3 4 5 6 7
Explanation:
Distinct elements including both the arrays are: 1 2 3 4 5 6 7.
Example 2:

Input:
n = 5, arr1[] = {2, 2, 3, 4, 5}
m = 5, arr2[] = {1, 1, 2, 3, 4}
Output:
1 2 3 4 5
Explanation:
Distinct elements including both the arrays are: 1 2 3 4 5.
Example 3:

Input:
n = 5, arr1[] = {1, 1, 1, 1, 1}
m = 5, arr2[] = {2, 2, 2, 2, 2}
Output:
1 2
Explanation:
Distinct elements including both the arrays are: 1 2.
Your Task:
You do not need to read input or print anything. Complete the function findUnion() that takes two arrays arr1[], arr2[], and their size n and m as input parameters and returns a list containing the union of the two arrays.

Expected Time Complexity: O(n+m).
Expected Auxiliary Space: O(n+m).

Constraints:
1 <= n, m <= 105
-109 <= arr1[i], arr2[i] <= 109
*/
package main

import (
	"fmt"
)

func findUnion(arr1 []int, arr2 []int) (union []int) {
	i, j := 0, 0

	for i < len(arr1) && j < len(arr2) {
		if arr1[i] <= arr2[j] {
			if len(union) == 0 || union[len(union)-1] != arr1[i] {
				union = append(union, arr1[i])
			}
			i++
		} else {
			if len(union) == 0 || union[len(union)-1] != arr2[j] {
				union = append(union, arr2[j])
			}
			j++
		}
	}

	for i < len(arr1) {
		if len(union) == 0 || union[len(union)-1] != arr1[i] {
			union = append(union, arr1[i])
		}
		i++
	}

	for j < len(arr2) {
		if len(union) == 0 || union[len(union)-1] != arr2[j] {
			union = append(union, arr2[j])
		}
		j++
	}

	return
}

func main() {
	arr1 := []int{1, 2, 3, 4, 5}
	arr2 := []int{1, 2, 3, 6, 7}

	union := findUnion(arr1, arr2)

	for _, element := range union {
		fmt.Printf("%d ", element)
	}
}
