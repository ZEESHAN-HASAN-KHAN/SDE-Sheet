/* Given an array nums with n objects colored red, white, or blue, sort them in-place so that objects of the same color are adjacent, with the colors in the order red, white, and blue.

We will use the integers 0, 1, and 2 to represent the color red, white, and blue, respectively.

You must solve this problem without using the library's sort function.

 

Example 1:

Input: nums = [2,0,2,1,1,0]
Output: [0,0,1,1,2,2]
Example 2:

Input: nums = [2,0,1]
Output: [0,1,2]
 

Constraints:

n == nums.length
1 <= n <= 300
nums[i] is either 0, 1, or 2.
*/

package main

import "fmt"

func sortColors(nums []int)  {
    var (
        i = 0
        j = len(nums)-1
    )

    for k :=0; k<len(nums) && k<=j; k++ {
        if nums[k] == 0 {
            nums[k], nums[i] = nums[i], nums[k]
            i++
        } else if nums[k] == 2 {
            nums[k] , nums[j] = nums[j], nums[k]
            k--
            j--
        }
    }
}

func main() {
	nums := []int{2,0,2,1,1,0}
	sortColors(nums)
	fmt.Println(nums)
}