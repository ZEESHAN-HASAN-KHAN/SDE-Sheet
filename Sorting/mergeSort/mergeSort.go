/*
Given an array arr[], its starting position l and its ending position r. Sort the array using merge sort algorithm.
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

Expected Time Complexity: O(nlogn)
Expected Auxiliary Space: O(n)

Constraints:
1 <= N <= 105
1 <= arr[i] <= 105
*/

package main

import (
	"fmt"
)

func merge(array []int, left int, mid int, right int) {

	// fmt.Printf("left = %d, right = %d, mid = %d", left, right, mid)

	// fmt.Print(" Array = ")

	// for ele := left; ele < right; ele++ {
	// 	fmt.Printf("%d ", array[ele])
	// }

	i := left
	j := mid + 1
	var temp []int

	for i <= mid && j <= right {
		if array[i] > array[j] {
			temp = append(temp, array[j])
			j++
		} else {
			temp = append(temp, array[i])
			i++
		}
	}

	for i <= mid {
		temp = append(temp, array[i])
		i++
	}

	for j <= right {
		temp = append(temp, array[j])
		j++
	}

	for k := range temp {
		array[k+left] = temp[k]
	}

	// fmt.Printf(" Sorted Array = ")
	// for ele := left; ele < right; ele++ {
	// 	fmt.Printf("%d ", array[ele])
	// }

	// fmt.Println()
}

func mergeSort(array []int, left int, right int) {
	if left < right {
		mid := (left + right) / 2
		mergeSort(array, left, mid)
		mergeSort(array, mid+1, right)
		merge(array, left, mid, right)
	}
}

func main() {
	var array []int = []int{4, 1, 3, 9, 7}
	mergeSort(array, 0, len(array)-1)

	for _, temp := range array {
		fmt.Printf("%d ", temp)
	}
}
