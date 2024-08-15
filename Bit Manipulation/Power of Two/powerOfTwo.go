/* Given an integer n, return true if it is a power of two. Otherwise, return false.

An integer n is a power of two, if there exists an integer x such that n == 2x.

 

Example 1:

Input: n = 1
Output: true
Explanation: 20 = 1
Example 2:

Input: n = 16
Output: true
Explanation: 24 = 16
Example 3:

Input: n = 3
Output: false
 

Constraints:

-231 <= n <= 231 - 1
*/

package main

func isPowerOfTwo(n int) bool {
    oneCount := 0

    for n > 0 {
        if n & 1 == 1{
            oneCount++
        }

        if oneCount > 1 {
            return false
        }

        n = n >> 1
    }

    if oneCount == 1 {
        return true
    }

    return false
}

func main() {
	println(isPowerOfTwo(1))
	println(isPowerOfTwo(16))
	println(isPowerOfTwo(3))	
}