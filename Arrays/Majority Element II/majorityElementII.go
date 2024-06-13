/* Given an integer array of size n, find all elements that appear more than ⌊ n/3 ⌋ times.

 

Example 1:

Input: nums = [3,2,3]
Output: [3]
Example 2:

Input: nums = [1]
Output: [1]
Example 3:

Input: nums = [1,2]
Output: [1,2]
 

Constraints:

1 <= nums.length <= 5 * 104
-109 <= nums[i] <= 109
*/
package main

import "fmt"

func majorityElement(nums []int) []int {
	var (
		ans []int
		count1, count2 int
		candidate1, candidate2 int
	)
	for _, num := range nums {
		if count1 == 0 && candidate2 != num {
			candidate1 = num
		} else if count2 == 0 && candidate1 != num {
			candidate2 = num
		}
		if candidate1 == num {
			count1++
		} else if candidate2 == num {
			count2++
		} else {
			count1--
			count2--
		}
	}
	count1, count2 = 0, 0
	for _, num := range nums {
		if candidate1 == num {
			count1++
		} else if candidate2 == num {
			count2++
		}
	}
	if count1 > len(nums)/3 {
		ans = append(ans, candidate1)
	}
	if count2 > len(nums)/3 {
		ans = append(ans, candidate2)
	}
	return ans
}

func main() {
	nums := []int{3,2,3}
	fmt.Println(majorityElement(nums)) //Output: [3]

	nums = []int{1}
	fmt.Println(majorityElement(nums)) //Output: [1]

	nums = []int{1,2}
	fmt.Println(majorityElement(nums)) //Output: [1,2]
}