/* Given an integer array nums, return the number of reverse pairs in the array.

A reverse pair is a pair (i, j) where:

0 <= i < j < nums.length and
nums[i] > 2 * nums[j].
 

Example 1:

Input: nums = [1,3,2,3,1]
Output: 2
Explanation: The reverse pairs are:
(1, 4) --> nums[1] = 3, nums[4] = 1, 3 > 2 * 1
(3, 4) --> nums[3] = 3, nums[4] = 1, 3 > 2 * 1
Example 2:

Input: nums = [2,4,3,5,1]
Output: 3
Explanation: The reverse pairs are:
(1, 4) --> nums[1] = 4, nums[4] = 1, 4 > 2 * 1
(2, 4) --> nums[2] = 3, nums[4] = 1, 3 > 2 * 1
(3, 4) --> nums[3] = 5, nums[4] = 1, 5 > 2 * 1
 

Constraints:

1 <= nums.length <= 5 * 104
-231 <= nums[i] <= 231 - 1
*/
package main

import "fmt"

func merge(nums []int, start int, mid int, end int) int {
    j := mid + 1
    count := 0

    for i:= start; i<=mid; i++ {
        for j<=end && int64(nums[i]) > 2*int64(nums[j]) {
            j++
        }
        count += j - (mid + 1)
    }

    left := start
    right := mid+1
    var temp []int

    for left <= mid && right <= end {
        if nums[left] < nums[right] {
            temp = append(temp, nums[left])
            left++
        } else {
            temp = append(temp, nums[right])
            right++
        }
    }

    for left <= mid {
        temp = append(temp, nums[left])
        left++
    }

    for right <=end {
        temp = append(temp, nums[right])
        right++
    }

    for k := 0; k < len(temp); k++ {
        nums[k + start] = temp[k]
    }

    return count
}



func mergeSort(nums []int, start int, end int) int {
    count := 0
    if start < end {
        mid := (start+end)/2
        count += mergeSort(nums, start, mid)
        count += mergeSort(nums, mid+1, end)
        count += merge(nums, start, mid, end)
    }

    return count
}

func reversePairs(nums []int) int {
    return mergeSort(nums, 0, len(nums)-1)
}

func main() {
	fmt.Println(reversePairs([]int{1,3,2,3,1})) // 2
	fmt.Println(reversePairs([]int{2,4,3,5,1})) // 3
}