/* A peak element is an element that is strictly greater than its neighbors.

Given a 0-indexed integer array nums, find a peak element, and return its index. If the array contains multiple peaks, return the index to any of the peaks.

You may imagine that nums[-1] = nums[n] = -∞. In other words, an element is always considered to be strictly greater than a neighbor that is outside the array.

You must write an algorithm that runs in O(log n) time.

 

Example 1:

Input: nums = [1,2,3,1]
Output: 2
Explanation: 3 is a peak element and your function should return the index number 2.
Example 2:

Input: nums = [1,2,1,3,5,6,4]
Output: 5
Explanation: Your function can return either index number 1 where the peak element is 2, or index number 5 where the peak element is 6.
 

Constraints:

1 <= nums.length <= 1000
-231 <= nums[i] <= 231 - 1
nums[i] != nums[i + 1] for all valid i.
*/
package main

import "fmt"

func findPeakElement(nums []int) int {
    if len(nums) == 1 {
        return 0
    } else if nums[0] > nums[1] {
        return 0
    } else if nums[len(nums)-1] > nums[len(nums)-2] {
        return len(nums) - 1
    }

    var (
        low = 1
        high = len(nums) - 1
    )

    for low <= high {
        mid := (low+high)/2
        if nums[mid] > nums[mid-1] && nums[mid] > nums[mid+1] {
            return mid
        } else if nums[mid] > nums[mid-1] && nums[mid] < nums[mid+1] {
            low = mid + 1
        } else {
            high = mid - 1
        }
    }

    return -1
}

func main() {
	fmt.Println(findPeakElement([]int{1,2,3,1})) // 2
	fmt.Println(findPeakElement([]int{1,2,1,3,5,6,4})) // 5
}