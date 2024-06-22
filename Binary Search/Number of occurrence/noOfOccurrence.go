/* Given a sorted array Arr of size N and a number X, you need to find the number of occurrences of X in Arr.

Example 1:

Input:
N = 7, X = 2
Arr[] = {1, 1, 2, 2, 2, 2, 3}
Output: 4
Explanation: 2 occurs 4 times in the
given array.
Example 2:

Input:
N = 7, X = 4
Arr[] = {1, 1, 2, 2, 2, 2, 3}
Output: 0
Explanation: 4 is not present in the
given array.
Your Task:
You don't need to read input or print anything.
Your task is to complete the function count() which takes the array of integers arr, n, and x as parameters and returns an integer denoting the answer.
If x is not present in the array (arr) then return 0.

Expected Time Complexity: O(logN)
Expected Auxiliary Space: O(1)

Constraints:
1 ≤ N ≤ 105
1 ≤ Arr[i] ≤ 106
1 ≤ X ≤ 106
*/
package main

import (
	"fmt"
)

func upperBound(nums []int, target int) int {
    var (
        low = 0
        high = len(nums) - 1
        upper = -1
    )

    for low <= high {
        mid := (low+high)/2
        if nums[mid] == target {
            upper = mid
            low = mid+1
        } else if nums[mid] <= target {
            low = mid+1
        } else if nums[mid] > target {
            high = mid-1
        }
    }
    return upper 
}


func lowerBound(nums []int, target int) int {
    var (
        low = 0
        high = len(nums) - 1
        lower = -1
    )

    for low <= high {
        mid := (low+high)/2
        if nums[mid] == target {
            lower = mid
            high = mid - 1
        } else if nums[mid] > target {
            high = mid-1
        } else if nums[mid] < target {
            low = mid+1
        }
    }
    return lower 
}

func count(arr []int, n int, x int) int {
	lower := lowerBound(arr, x)
	upper := upperBound(arr, x)
	if lower == -1 {
		return 0
	}
	return upper - lower + 1
}

func main() {
	arr := []int{1, 1, 2, 2, 2, 2, 3}
	n := 7
	x := 2
	fmt.Println(count(arr, n, x)) // 4
}