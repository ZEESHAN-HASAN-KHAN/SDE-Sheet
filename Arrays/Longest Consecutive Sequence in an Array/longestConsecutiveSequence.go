/* Given an unsorted array of integers nums, return the length of the longest consecutive elements sequence.

You must write an algorithm that runs in O(n) time.

 

Example 1:

Input: nums = [100,4,200,1,3,2]
Output: 4
Explanation: The longest consecutive elements sequence is [1, 2, 3, 4]. Therefore its length is 4.
Example 2:

Input: nums = [0,3,7,2,5,8,4,6,0,1]
Output: 9
 

Constraints:

0 <= nums.length <= 105
-109 <= nums[i] <= 109
*/
package main

import "fmt"

func max(a int, b int) int {
    if a > b {
        return a
    }
    return b
}

func longestConsecutive(nums []int) int {
    ump := make(map[int]struct{}) // Initialize the map properly
    ans := 0

    for _, element := range nums {
        ump[element] = struct{}{}
    }

    for i := range ump {
        if _, ok := ump[i-1]; !ok { // Check if there's no predecessor
            x := i
            count := 1
            for {
                if _, ok := ump[x+1]; ok { // Check if there's a successor
                    x++
                    count++
                } else {
                    break
                }
            }
            ans = max(ans, count)
        }
    }
    return ans
}

func main() {
	nums := []int{100, 4, 200, 1, 3, 2}
	fmt.Println(longestConsecutive(nums))
}