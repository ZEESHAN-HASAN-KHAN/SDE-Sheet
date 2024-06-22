/* Given an m x n matrix, return all elements of the matrix in spiral order.

 

Example 1:


Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
Output: [1,2,3,6,9,8,7,4,5]
Example 2:


Input: matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
Output: [1,2,3,4,8,12,11,10,9,5,6,7]
 

Constraints:

m == matrix.length
n == matrix[i].length
1 <= m, n <= 10
-100 <= matrix[i][j] <= 100
*/
package main

import (
	"fmt"
)

func spiralOrder(matrix [][]int) []int {
    var (
        count = 0
        left = 0
        right = 0
        row = len(matrix)
        col = len(matrix[0])
        up = 0
        ans []int
        m = len(matrix)
        n = len(matrix[0])
    )

    for count < m*n {
        for i := left; i < col && count < m*n; i++ {
            ans = append(ans, matrix[left][i])
            count++
        }

        for i := left + 1; i < row && count < m*n; i++ {
            ans = append(ans, matrix[i][col-1])
            count++
        }

        for i := col-2 ; i >= right && count < m*n; i-- {
            ans = append(ans, matrix[row-1][i])
            count++
        }

        for i := row-2; i>= up+1 && count < m*n; i-- {
            ans = append(ans, matrix[i][up])
            count++;
        }

        left++
        right++
        up++
        row--
        col--
    }

    return ans
}

func main() {
	matrix := [][]int{{1,2,3},{4,5,6},{7,8,9}}
	fmt.Println(spiralOrder(matrix))
}