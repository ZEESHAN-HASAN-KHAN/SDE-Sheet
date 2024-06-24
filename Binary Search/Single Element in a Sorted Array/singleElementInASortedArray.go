/* You are given a sorted array consisting of only integers where every element appears exactly twice, except for one element which appears exactly once.

Return the single element that appears only once.

Your solution must run in O(log n) time and O(1) space.

 

Example 1:

Input: nums = [1,1,2,3,3,4,4,8,8]
Output: 2
Example 2:

Input: nums = [3,3,7,7,10,11,11]
Output: 10
 

Constraints:

1 <= nums.length <= 105
0 <= nums[i] <= 105
*/
package main

import "fmt"

func singleNonDuplicate(nums []int) int {
    if len(nums) == 1 {
        return nums[0]
    } else if (nums[0] != nums[1]) {
        return nums[0]
    } else if (nums[len(nums)-1] != nums[len(nums)-2]) {
        return nums[len(nums)-1]
    }

    var (
        low = 0
        high = len(nums) - 2
    )

    for low <= high {
        mid := (low+high)/2
        if nums[mid] != nums[mid+1] && nums[mid] != nums[mid-1] {
            return nums[mid]
        } else if (mid%2 == 0 && nums[mid+1] == nums[mid]) || (mid%2 == 1 && nums[mid-1] == nums[mid]) {
            low = mid + 1
        } else {
            high = mid - 1
        }
    }
    return -1
}

func main() {
	fmt.Println(singleNonDuplicate([]int{1,1,2,3,3,4,4,8,8})) // 2
	fmt.Println(singleNonDuplicate([]int{3,3,7,7,10,11,11})) // 10
}