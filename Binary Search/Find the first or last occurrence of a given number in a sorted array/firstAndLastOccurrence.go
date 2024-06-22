/* Given an array of integers nums sorted in non-decreasing order, find the starting and ending position of a given target value.

If target is not found in the array, return [-1, -1].

You must write an algorithm with O(log n) runtime complexity.

 

Example 1:

Input: nums = [5,7,7,8,8,10], target = 8
Output: [3,4]
Example 2:

Input: nums = [5,7,7,8,8,10], target = 6
Output: [-1,-1]
Example 3:

Input: nums = [], target = 0
Output: [-1,-1]
 

Constraints:

0 <= nums.length <= 105
-109 <= nums[i] <= 109
nums is a non-decreasing array.
-109 <= target <= 109
*/
package main

import (
	"fmt"
)

func upperBound(nums []int, target int) int {
    var (
        low = 0
        high = len(nums) - 1
        upper = -1
    )

    for low <= high {
        mid := (low+high)/2
        if nums[mid] == target {
            upper = mid
            low = mid+1
        } else if nums[mid] <= target {
            low = mid+1
        } else if nums[mid] > target {
            high = mid-1
        }
    }
    return upper 
}


func lowerBound(nums []int, target int) int {
    var (
        low = 0
        high = len(nums) - 1
        lower = -1
    )

    for low <= high {
        mid := (low+high)/2
        if nums[mid] == target {
            lower = mid
            high = mid - 1
        } else if nums[mid] > target {
            high = mid-1
        } else if nums[mid] < target {
            low = mid+1
        }
    }
    return lower 
}


func searchRange(nums []int, target int) []int {
    lower := lowerBound(nums, target)
    upper := upperBound(nums, target)

    return []int{lower, upper}
}

func main() {
	fmt.Println(searchRange([]int{5,7,7,8,8,10}, 8)) // Output: [3,4]
	fmt.Println(searchRange([]int{5,7,7,8,8,10}, 6)) // Output: [-1,-1]
	fmt.Println(searchRange([]int{}, 0)) // Output: [-1,-1]
}