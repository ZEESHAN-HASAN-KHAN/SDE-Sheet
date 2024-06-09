/* Given an array A of positive integers. Your task is to find the leaders in the array. An element of the array is a leader if it is greater than all the elements to its right side or if it is equal to the maximum element on its right side. The rightmost element is always a leader. 

Example 1:

Input:
n = 6
A[] = {16,17,4,3,5,2}
Output: 17 5 2
Explanation: The first leader is 17 
as it is greater than all the elements
to its right.  Similarly, the next 
leader is 5. The right most element 
is always a leader so it is also 
included.
Example 2:

Input:
n = 5
A[] = {10,4,2,4,1}
Output: 10 4 4 1
Explanation: 1 is the rightmost element and 4 is the element which is greater
than all the elements to its right. Similarly another 4 is the element that is equal to the greatest element to its right and 10 is the greatest element in the whole array.
Your Task:
You don't need to read input or print anything. The task is to complete the function leader() which takes array A and n as input parameters and returns an array of leaders in order of their appearance.

Expected Time Complexity: O(n)
Expected Auxiliary Space: O(n)

Constraints:
1 <= n <= 107
0 <= Ai <= 107

*/
package main

import (
	"fmt"
)

func leaders(A []int, n int) []int {
	var (
		ans []int
		max = A[n-1]
	)
	ans = append(ans, max)
	for i:=n-2; i>=0; i-- {
		if A[i] >= max {
			max = A[i]
			ans = append(ans, max)
		}
	}
	// reverse the array
	for i:=0; i<len(ans)/2; i++ {
		ans[i], ans[len(ans)-1-i] = ans[len(ans)-1-i], ans[i]
	}
	return ans
}

func main() {
	fmt.Println(leaders([]int{16,17,4,3,5,2}, 6))
	fmt.Println(leaders([]int{10,4,2,4,1}, 5))
}