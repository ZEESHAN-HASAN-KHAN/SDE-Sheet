/* Given an array of intervals where intervals[i] = [starti, endi], merge all overlapping intervals, and return an array of the non-overlapping intervals that cover all the intervals in the input.

 

Example 1:

Input: intervals = [[1,3],[2,6],[8,10],[15,18]]
Output: [[1,6],[8,10],[15,18]]
Explanation: Since intervals [1,3] and [2,6] overlap, merge them into [1,6].
Example 2:

Input: intervals = [[1,4],[4,5]]
Output: [[1,5]]
Explanation: Intervals [1,4] and [4,5] are considered overlapping.
 

Constraints:

1 <= intervals.length <= 104
intervals[i].length == 2
0 <= starti <= endi <= 104
*/
package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
    var (
        ans  [][]int
    )
    sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
    for _,element := range intervals {
        if len(ans) == 0 || ans[len(ans)-1][1] < element[0] {
            ans = append(ans, element)
        } else {
            ans[len(ans)-1][1] = max(ans[len(ans)-1][1], element[1])
        }
    }
    return ans
}

func main() {
	intervals := [][]int{{1,3},{2,6},{8,10},{15,18}}
	fmt.Println(merge(intervals)) // [[1,6],[8,10],[15,18]]

	intervals = [][]int{{1,4},{4,5}}
	fmt.Println(merge(intervals)) // [[1,5]]
}