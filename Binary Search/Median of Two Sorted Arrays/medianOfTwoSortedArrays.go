/* Given two sorted arrays nums1 and nums2 of size m and n respectively, return the median of the two sorted arrays.

The overall run time complexity should be O(log (m+n)).

 

Example 1:

Input: nums1 = [1,3], nums2 = [2]
Output: 2.00000
Explanation: merged array = [1,2,3] and median is 2.
Example 2:

Input: nums1 = [1,2], nums2 = [3,4]
Output: 2.50000
Explanation: merged array = [1,2,3,4] and median is (2 + 3) / 2 = 2.5.
 

Constraints:

nums1.length == m
nums2.length == n
0 <= m <= 1000
0 <= n <= 1000
1 <= m + n <= 2000
-106 <= nums1[i], nums2[i] <= 106
*/
package main

import (
	"math"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)

	if m > n {
		return findMedianSortedArrays(nums2, nums1)
	}

	low, high := 0, m
	medianPos := (m + n + 1) / 2

	for low <= high {
		mid1 := (low + high) / 2
		mid2 := medianPos - mid1

		l1 := math.MinInt
		l2 := math.MinInt
		r1 := math.MaxInt
		r2 := math.MaxInt

		if mid1 < m {
			r1 = nums1[mid1]
		}

		if mid2 < n {
			r2 = nums2[mid2]
		}

		if mid1-1 >= 0 {
			l1 = nums1[mid1-1]
		}

		if mid2-1 >= 0 {
			l2 = nums2[mid2-1]
		}

		if l1 <= r2 && l2 <= r1 {
			if (m+n)%2 != 0 {
				return float64(max(l1, l2))
			}
			return float64(max(l1, l2)+min(r1, r2)) / 2.0
		}

		if l1 > r2 {
			high = mid1 - 1
		} else {
			low = mid1 + 1
		}
	}

	return 0
}

func main() {
	println(findMedianSortedArrays([]int{1, 3}, []int{2})) // 2.0
	println(findMedianSortedArrays([]int{1, 2}, []int{3, 4})) // 2.5
}