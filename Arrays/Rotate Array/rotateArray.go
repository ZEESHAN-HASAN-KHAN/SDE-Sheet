/*
Given an integer array nums, rotate the array to the right by k steps, where k is non-negative.

Example 1:

Input: nums = [1,2,3,4,5,6,7], k = 3
Output: [5,6,7,1,2,3,4]
Explanation:
rotate 1 steps to the right: [7,1,2,3,4,5,6]
rotate 2 steps to the right: [6,7,1,2,3,4,5]
rotate 3 steps to the right: [5,6,7,1,2,3,4]
Example 2:

Input: nums = [-1,-100,3,99], k = 2
Output: [3,99,-1,-100]
Explanation:
rotate 1 steps to the right: [99,-1,-100,3]
rotate 2 steps to the right: [3,99,-1,-100]


Constraints:

1 <= nums.length <= 105
-231 <= nums[i] <= 231 - 1
0 <= k <= 105
*/

package main

import (
	"fmt"
)

func reverse(arr []int, start int, end int) {
	for start < end {
		arr[start], arr[end] = arr[end], arr[start]
		start++
		end--
	}
}

func rotate(nums []int, k int) {
	k = k % len(nums)

	reverse(nums, 0, len(nums)-k-1)
	reverse(nums, 0, len(nums)-1)
	reverse(nums, 0, k-1)
}

func main() {
	arr := []int{-1, -100, 3, 99}
	k := 2
	rotate(arr, k)

	for _, element := range arr {
		fmt.Printf("%d ", element)
	}
}
