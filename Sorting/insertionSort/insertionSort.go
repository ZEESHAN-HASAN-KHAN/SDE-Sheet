/* Given an Integer N and a list arr. Sort the array using insertion sort algorithm.
Example 1:

Input:
N = 5
arr[] = {4, 1, 3, 9, 7}
Output:
1 3 4 7 9
Example 2:

Input:
N = 10
arr[] = {10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
Output:
1 2 3 4 5 6 7 8 9 10


Expected Time Complexity: O(N^2).
Expected Auxiliary Space: O(1).

Constraints:
1 <= N <= 103
1 <= arr[i] <= 103
*/

package main

import (
	"fmt"
)

func insertionSort(array []int) {
	for i := 1; i < len(array); i++ {
		j := i
		for j > 0 && array[j] < array[j-1] {
			array[j-1], array[j] = array[j], array[j-1]
			j--
		}
	}
}

func main() {
	var array []int = []int{4, 1, 3, 9, 7}
	insertionSort(array)

	for _, temp := range array {
		fmt.Printf("%d ", temp)
	}
}
