/* Given an array of integers nums and an integer k, return the total number of subarrays whose sum equals to k.

A subarray is a contiguous non-empty sequence of elements within an array.

 

Example 1:

Input: nums = [1,1,1], k = 2
Output: 2
Example 2:

Input: nums = [1,2,3], k = 3
Output: 2
 

Constraints:

1 <= nums.length <= 2 * 104
-1000 <= nums[i] <= 1000
-107 <= k <= 107
*/
package main 

import (
	"fmt"
)

func subarraySum(nums []int, k int) int {
    var (
        count = 0
        ump  = make(map[int]int)
        sum = 0
    )

    ump[0] = 1

    for _,element := range nums {
        sum += element
        
        if _,ok := ump[sum-k]; ok {
            count += ump[sum-k]
        }

        ump[sum]++
    }

    return count
}

func main() {
	nums := []int{1,1,1}
	k := 2
	fmt.Println(subarraySum(nums,k)) //Output: 2

	nums = []int{1,2,3}
	k = 3
	fmt.Println(subarraySum(nums,k)) //Output: 2
}