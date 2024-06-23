/* Given an ascending sorted rotated array arr of distinct integers of size n. The array is right-rotated k times. Find the value of k.

Example 1:

Input:
n = 5
arr[] = {5, 1, 2, 3, 4}
Output: 1
Explanation: The given array is 5 1 2 3 4. 
The original sorted array is 1 2 3 4 5. 
We can see that the array was rotated 
1 times to the right.
Example 2:

Input:
n = 5
arr[] = {1, 2, 3, 4, 5}
Output: 0
Explanation: The given array is not rotated.
Your Task:
Complete the function findKRotation() which takes array arr and size n, as input parameters and returns an integer representing the answer. You don't have to print answers or take inputs.

Expected Time Complexity: O(log(n))
Expected Auxiliary Space: O(1)

Constraints:
1 <= n <=105
1 <= arri <= 107
*/
package main

import (
	"fmt"
)

func findKRotation(arr []int, n int) int {
	var (
		low = 0
		high = n-1
		minimum = arr[low]
		index = low
		prev = minimum
	)

	for low <= high {
		mid := (low+high)/2
		if arr[mid] >= arr[low] {
			minimum = min(minimum, arr[low])
			if minimum != prev {
				index = low
				prev = minimum
			}
			low = mid+1
		} else {
			minimum = min(minimum, arr[mid])
			if minimum != prev {
				index = mid
				prev = minimum
			}
			high = mid - 1
		}
	}

	return index
}

func main() {
	fmt.Println(findKRotation([]int{5, 1, 2, 3, 4}, 5))
	fmt.Println(findKRotation([]int{1, 2, 3, 4, 5}, 5))
}