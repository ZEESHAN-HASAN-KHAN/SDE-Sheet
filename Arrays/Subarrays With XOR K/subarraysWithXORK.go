/* Given an array ‘A’ consisting of ‘N’ integers and an integer ‘B’, find the number of subarrays of array ‘A’ whose bitwise XOR( ⊕ ) of all elements is equal to ‘B’.



A subarray of an array is obtained by removing some(zero or more) elements from the front and back of the array.



Example:
Input: ‘N’ = 4 ‘B’ = 2
‘A’ = [1, 2, 3, 2]

Output: 3

Explanation: Subarrays have bitwise xor equal to ‘2’ are: [1, 2, 3, 2], [2], [2].
Detailed explanation ( Input/output format, Notes, Images )
Sample Input 1:
4 2
1 2 3 2
Sample Output 1 :
3
Explanation Of Sample Input 1:
Input: ‘N’ = 4 ‘B’ = 2
‘A’ = [1, 2, 3, 2]

Output: 3

Explanation: Subarrays have bitwise xor equal to ‘2’ are: [1, 2, 3, 2], [2], [2].
Sample Input 2:
4 3
1 2 3 3
Sample Output 2:
4
Sample Input 3:
5 6
1 3 3 3 5 
Sample Output 3:
2
Constraints:
1 <= N <= 10^3
1 <= A[i], B <= 10^9

Time Limit: 1-sec
*/
package main 

import (
	"fmt"
)

func subarraysWithXORK(A []int, B int) int {
	var (
		count = 0
		xor = 0
		ump = make(map[int]int)
	)

	for i := range A {
		xor ^= A[i]
		if xor == B {
			count++
		}
		if v, ok := ump[xor^B]; ok {
			count += v
		}
		ump[xor]++
	}
	return count
}

func main() {
	A := []int{1, 2, 3, 2}
	B := 2
	fmt.Println(subarraysWithXORK(A, B)) // 3

	A = []int{1, 2, 3, 3}
	B = 3
	fmt.Println(subarraysWithXORK(A, B)) // 4

	A = []int{1, 3, 3, 3, 5}
	B = 6
	fmt.Println(subarraysWithXORK(A, B)) // 2
}