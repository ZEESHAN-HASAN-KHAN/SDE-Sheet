/* Given an array of integers. Find the Inversion Count in the array. 

Inversion Count: For an array, inversion count indicates how far (or close) the array is from being sorted. If the array is already sorted then the inversion count is 0.
If an array is sorted in the reverse order then the inversion count is the maximum. 
Formally, two elements arr[i] and arr[j] form an inversion if arr[i] > arr[j] and i < j.

Examples:

Input: n = 5, arr[] = {2, 4, 1, 3, 5}  Output: 3
Explanation: The sequence 2, 4, 1, 3, 5 
has three inversions (2, 1), (4, 1), (4, 3).
Input: n = 5, arr[] = {2, 3, 4, 5, 6}  Output: 0
Explanation: As the sequence is already 
sorted so there is no inversion count.
Input: n = 3, arr[] = {10, 10, 10}  Output: 0
Explanation: As all the elements of array 
are same, so there is no inversion count.
Expected Time Complexity: O(nLogn).
Expected Auxiliary Space: O(n).

Constraints:
1 ≤ n ≤ 5*105
1 ≤ arr[i] ≤ 1018
*/

package main

import (
	"fmt"
)

func merge(arr []int64, start, mid, end int) int64 {
	i := start
	j := mid + 1
	count := int64(0)
	temp := make([]int64, 0, end-start+1)

	for i <= mid && j <= end {
		if arr[i] <= arr[j] {
			temp = append(temp, arr[i])
			i++
		} else {
			temp = append(temp, arr[j])
			j++
			count += int64(mid - i + 1)
		}
	}

	for i <= mid {
		temp = append(temp, arr[i])
		i++
	}

	for j <= end {
		temp = append(temp, arr[j])
		j++
	}

	for k := 0; k < len(temp); k++ {
		arr[start+k] = temp[k]
	}

	return count
}

func mergeSort(arr []int64, start, end int) int64 {
	count := int64(0)
	if start < end {
		mid := (start + end) / 2
		count += mergeSort(arr, start, mid)
		count += mergeSort(arr, mid+1, end)
		count += merge(arr, start, mid, end)
	}
	return count
}

func inversionCount(arr []int64, n int) int64 {
	return mergeSort(arr, 0, n-1)
}

func main() {
	arr := []int64{4, 3, 6, 2, 1, 5}
	n := len(arr)
	inversions := inversionCount(arr, n)
	fmt.Printf("Number of inversions: %d\n", inversions)
}
