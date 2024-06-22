/* Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.

Notice that the solution set must not contain duplicate triplets.

 

Example 1:

Input: nums = [-1,0,1,2,-1,-4]
Output: [[-1,-1,2],[-1,0,1]]
Explanation: 
nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0.
nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0.
nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0.
The distinct triplets are [-1,0,1] and [-1,-1,2].
Notice that the order of the output and the order of the triplets does not matter.
Example 2:

Input: nums = [0,1,1]
Output: []
Explanation: The only possible triplet does not sum up to 0.
Example 3:

Input: nums = [0,0,0]
Output: [[0,0,0]]
Explanation: The only possible triplet sums up to 0.
 

Constraints:

3 <= nums.length <= 3000
-105 <= nums[i] <= 105
*/

package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
    var (
        ans [][]int
    )

    sort.Ints(nums)

    for i := range nums {
        if i != 0 && nums[i] == nums[i-1] {
            continue
        }

        var (
            j = i+1
            k = len(nums) - 1
        )

        for j < k {
            sum := nums[i] + nums[j] + nums[k]
            if sum > 0 {
                k--
            } else if sum < 0 {
                j++
            } else {
                ans = append(ans, []int{nums[i], nums[j], nums[k]})
                j++
                k--

                for j < k && nums[j] == nums[j-1] {
                    j++
                }

                for j < k && nums[k] == nums[k+1] {
                    k--
                }
            }
        }
    }
    return ans
}

// test case
func main() {
	nums := []int{-1,0,1,2,-1,-4}
	fmt.Println(threeSum(nums)) //Output: [[-1,-1,2],[-1,0,1]]

	nums = []int{0,1,1}
	fmt.Println(threeSum(nums)) //Output: []

	nums = []int{0,0,0}
	fmt.Println(threeSum(nums)) //Output: [[0,0,0]]
}