/* Given an integer numRows, return the first numRows of Pascal's triangle.

In Pascal's triangle, each number is the sum of the two numbers directly above it as shown:


 

Example 1:

Input: numRows = 5
Output: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]
Example 2:

Input: numRows = 1
Output: [[1]]
 

Constraints:

1 <= numRows <= 30
*/

package main

import (
	"fmt"
)

func generate(numRows int) [][]int {
    var (
        ans [][]int
    )
    for i:=0;i<numRows;i++ {
        temp := make([]int, i+1)
        temp[0] = 1
        temp[i] = 1
        for j := 1; j <= i-1; j++ {
            temp[j] = ans[i-1][j-1] + ans[i-1][j]
        } 
        ans = append(ans, temp)
    }
    return ans
}

func main() {
	numRows := 5
	fmt.Println(generate(numRows)) //Output: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]

	numRows = 1
	fmt.Println(generate(numRows)) //Output: [[1]]
}