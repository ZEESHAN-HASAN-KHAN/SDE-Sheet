/* Given an array Arr of size N, print the second largest distinct element from an array. If the second largest element doesn't exist then return -1.

Example 1:

Input:
N = 6
Arr[] = {12, 35, 1, 10, 34, 1}
Output: 34
Explanation: The largest element of the
array is 35 and the second largest element
is 34.
Example 2:

Input:
N = 3
Arr[] = {10, 5, 10}
Output: 5
Explanation: The largest element of
the array is 10 and the second
largest element is 5.
Your Task:
You don't need to read input or print anything. Your task is to complete the function print2largest() which takes the array of integers arr and n as parameters and returns an integer denoting the answer. If 2nd largest element doesn't exist then return -1.

Expected Time Complexity: O(N)
Expected Auxiliary Space: O(1)

Constraints:
2 ≤ N ≤ 105
1 ≤ Arri ≤ 105

*/

package main

import (
	"fmt"
	"math"
)

func max(a int, b int) int {
	if a > b {
		return a
	}

	return b
}

func secondLargestElement(array []int) int {
	largest, secondLargest := math.MinInt, -1

	for _, element := range array {
		largest = max(element, largest)
	}

	for _, element := range array {
		if element != largest && element > secondLargest {
			secondLargest = element
		}
	}

	return secondLargest
}

func main() {
	var array []int = []int{12, 35, 1, 10, 34, 1}

	fmt.Printf("The second largest element is %d. \n", secondLargestElement(array))

}
