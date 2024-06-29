/* Problem statement
You are given an array 'arr' consisting of 'n' integers which denote the position of a stall.



You are also given an integer 'k' which denotes the number of aggressive cows.



You are given the task of assigning stalls to 'k' cows such that the minimum distance between any two of them is the maximum possible.



Print the maximum possible minimum distance.



Example:
Input: 'n' = 3, 'k' = 2 and 'arr' = {1, 2, 3}

Output: 2

Explanation: The maximum possible minimum distance will be 2 when 2 cows are placed at positions {1, 3}. Here distance between cows is 2.
Detailed explanation ( Input/output format, Notes, Images )
Sample Input 1 :
6 4
0 3 4 7 10 9


Sample Output 1 :
3


Explanation to Sample Input 1 :
The maximum possible minimum distance between any two cows will be 3 when 4 cows are placed at positions {0, 3, 7, 10}. Here distance between cows are 3, 4 and 3 respectively.


Sample Input 2 :
5 2
4 2 1 3 6


Sample Output 2 :
5


Expected time complexity:
Can you solve this in O(n * log(n)) time complexity?


Constraints :
2 <= 'n' <= 10 ^ 5
2 <= 'k' <= n
0 <= 'arr[i]' <= 10 ^ 9
Time Limit: 1 sec.
*/

package main

import (
	"sort"
)

func isAllotmentPossible(arr []int, k, mid int) bool {
	cows := 1
	prev := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i]-prev >= mid {
			cows++
			prev = arr[i]
		}
	}
	return cows >= k
}

func aggressiveCows(n, k int, arr []int) int {
	sort.Ints(arr)
	var (
		low  = 1
		high = arr[n-1] - arr[0]
	)

	for low <= high {
		mid := (low + high) / 2
		if isAllotmentPossible(arr, k, mid) {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return high
}

func main() {
	println(aggressiveCows(6, 4, []int{0, 3, 4, 7, 10, 9})) // 3
	println(aggressiveCows(5, 2, []int{4, 2, 1, 3, 6}))     // 5
}
