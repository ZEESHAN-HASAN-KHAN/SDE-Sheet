/* Given an array containing N integers and an integer K., Your task is to find the length of the longest Sub-Array with the sum of the elements equal to the given value K.

 

Example 1:
 

Input :
A[] = {10, 5, 2, 7, 1, 9}
K = 15
Output : 4
Explanation:
The sub-array is {5, 2, 7, 1}.
Example 2:

Input : 
A[] = {-1, 2, 3}
K = 6
Output : 0
Explanation: 
There is no such sub-array with sum 6.
Your Task:
This is a function problem. The input is already taken care of by the driver code. You only need to complete the function lenOfLongSubarr() that takes an array (A), sizeOfArray (n),  sum (K)and returns the required length of the longest Sub-Array. The driver code takes care of the printing.

Expected Time Complexity: O(N).
Expected Auxiliary Space: O(N).

 

Constraints:
1<=N<=105
-105<=A[i], K<=105
*/

package main 

import "fmt"

func max(a int, b int) int {
	if a > b {
		return a
	}

	return b
}

func lenOfLongSubarr(arr []int, k int) (ans int){
	var (
		ump = make(map[int]int)
		sum = 0
	)

	for i, element := range arr {
		sum += element

		if sum == k {
			ans = max(ans, i+1)
		}

		rem := sum - k

		if _, ok := ump[rem]; ok {
			ans = max(ans, i - ump[rem])
		}

		ump[sum] = i
	}

	return
}

func main() {
	arr := []int{10, 5, 2, 7, 1, 9}
	k := 15

	fmt.Printf("Longest Sub-Array with sum %d: %d\n", k, lenOfLongSubarr(arr, k))
}