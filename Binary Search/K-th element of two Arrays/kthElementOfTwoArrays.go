/* Given two sorted arrays arr1 and arr2 of size N and M respectively and an element K. The task is to find the element that would be at the kth position of the final sorted array.


Example 1:

Input:
arr1[] = {2, 3, 6, 7, 9}
arr2[] = {1, 4, 8, 10}
k = 5
Output:
6
Explanation:
The final sorted array would be -
1, 2, 3, 4, 6, 7, 8, 9, 10
The 5th element of this array is 6.
Example 2:
Input:
arr1[] = {100, 112, 256, 349, 770}
arr2[] = {72, 86, 113, 119, 265, 445, 892}
k = 7
Output:
256
Explanation:
Final sorted array is - 72, 86, 100, 112,
113, 119, 256, 265, 349, 445, 770, 892
7th element of this array is 256.

Your Task:
You don't need to read input or print anything. Your task is to complete the function kthElement() which takes the arrays arr1[], arr2[], its size N and M respectively and an integer K as inputs and returns the element at the Kth position.


Expected Time Complexity: O(Log(N) + Log(M))
Expected Auxiliary Space: O(Log (N))


Constraints:
1 <= N, M <= 106
0 <= arr1i, arr2i < INT_MAX
1 <= K <= N+M
*/

package main

import (
	"math"
)

func kthElement(nums1 []int, nums2 []int, n1 int, n2 int, k int) int {
	if n1 > n2 {
		return kthElement(nums2, nums1, n2, n1, k)
	}

	left := k
	low := max(0, k-n2)
	high := min(k, n1)

	for low <= high {
		mid1 := (low + high) >> 1
		mid2 := left - mid1

		l1, l2 := math.MinInt, math.MinInt
		r1, r2 := math.MaxInt, math.MaxInt

		if mid1 < n1 {
			r1 = nums1[mid1]
		}

		if mid2 < n2 {
			r2 = nums2[mid2]
		}

		if mid1-1 >= 0 {
			l1 = nums1[mid1-1]
		}

		if mid2-1 >= 0 {
			l2 = nums2[mid2-1]
		}

		if l1 <= r2 && l2 <= r1 {
			return max(l1, l2)
		} else if l1 > r2 {
			high = mid1 - 1
		} else {
			low = mid1 + 1
		}
	}

	return 0
}

func main() {
	nums1 := []int{2, 3, 6, 7, 9}
	nums2 := []int{1, 4, 8, 10}
	n1 := len(nums1)
	n2 := len(nums2)
	k := 5
	result := kthElement(nums1, nums2, n1, n2, k)
	println(result) // Output the result
}
