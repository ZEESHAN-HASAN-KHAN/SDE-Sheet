/* Given an Integer N and a list arr. Sort the array using quick sort algorithm.
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


Expected Time Complexity: O(N*logN)
Expected Auxiliary Space: O(logN)

Constraints:
1 <= N <= 103
1 <= arr[i] <= 103
*/

package main

import (
	"fmt"
)

func partionArray(array []int, start int, end int) int {
	i := start
	j := end
	pivot := array[start]

	for i < j {
		for i < end && array[i] <= pivot {
			i++
		}

		for j > start && array[j] >= pivot {
			j--
		}

		if i < j {
			array[i], array[j] = array[j], array[i]
		}
	}

	array[start], array[j] = array[j], array[start]

	return j
}

func quickSort(array []int, start int, end int) {
	if start < end {
		partionIndex := partionArray(array, start, end)
		quickSort(array, start, partionIndex)
		quickSort(array, partionIndex+1, end)
	}
}

func main() {
	var array []int = []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	quickSort(array, 0, len(array)-1)

	for _, temp := range array {
		fmt.Printf("%d ", temp)
	}
}
