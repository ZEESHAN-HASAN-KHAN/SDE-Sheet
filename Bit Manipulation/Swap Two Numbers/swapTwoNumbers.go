/* Swap given two numbers and print them. (Try to do it without a temporary variable.) and return it.

Example 1:

Input: a = 13, b = 9
Output: 9 13
Explanation: after swapping it
becomes 9 and 13.
Example 2:

Input: a = 15, b = 8
Output: 8 15
Explanation: after swapping it
becomes 8 and 15.
Your Task:  
You don't need to read input or print anything. Your task is to complete the function get() which takes a, b as inputs and returns the list of integers contains the new value of a and b after swap.

Expected Time Complexity: O(1)
Expected Auxiliary Space: O(1)

Constraints:
1 ≤ a, b ≤ 106
*/
package main

import "fmt"

func get(a, b int) []int {
	a = a ^ b
	b = a ^ b
	a = a ^ b
	return []int{a, b}
}

func main() {
	fmt.Println(get(13, 9)) // Output: [9 13]
	fmt.Println(get(15, 8)) // Output: [8 15]
}