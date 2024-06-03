/* 
Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.

Note that you must do this in-place without making a copy of the array.

 

Example 1:

Input: nums = [0,1,0,3,12]
Output: [1,3,12,0,0]
Example 2:

Input: nums = [0]
Output: [0]
 

Constraints:

1 <= nums.length <= 104
-231 <= nums[i] <= 231 - 1
*/

package main

import (
	"fmt"
)

func moveZeroes(nums []int)  {
    i , j := 0, 0
    for i < len(nums) && j < len(nums) {
        for j < len(nums) && nums[j] != 0 {
            j++
        }
        i=j
        for i < len(nums) && nums[i] == 0 {
            i++
        }

        if i < len(nums) && j < len(nums) && j < i {
            nums[i], nums[j] = nums[j], nums[i]
        }
    }
}

func main() {
	nums := []int{0,1,0,3,12}
	moveZeroes(nums)

	for _, element := range nums {
		fmt.Printf("%d ", element)
	}
}

















