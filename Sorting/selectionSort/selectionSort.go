/*
Selection Sort

Given an unsorted array of size N, use selection sort to sort arr[] in increasing order.

Example 1:

Input:
N = 5
arr[] = {4, 1, 3, 9, 7}
Output:
1 3 4 7 9
Explanation:
Maintain sorted (in bold) and unsorted subarrays.
Select 1. Array becomes 1 4 3 9 7.
Select 3. Array becomes 1 3 4 9 7.
Select 4. Array becomes 1 3 4 9 7.
Select 7. Array becomes 1 3 4 7 9.
Select 9. Array becomes 1 3 4 7 9.
Example 2:

Input:
N = 10
arr[] = {10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
Output:
1 2 3 4 5 6 7 8 9 10

Expected Time Complexity: O(N2)
Expected Auxiliary Space: O(1)


Constraints:
1 ≤ N ≤ 10^3

*/

package main

import (
	"fmt"
)

func selectionSort(array []int) {
	for i := range array {
		minPos := i
		for j := range array[i+1:] {
			j += i + 1 // Correct the index to refer to the original array
			if array[minPos] > array[j] {
				minPos = j
			}
		}
		// Swap the elements
		array[i], array[minPos] = array[minPos], array[i]
	}
}

func main() {
	var array []int = []int{4, 1, 3, 9, 7}
	selectionSort(array)

	for _, element := range array {
		fmt.Printf("%d ", element)
	}
}
