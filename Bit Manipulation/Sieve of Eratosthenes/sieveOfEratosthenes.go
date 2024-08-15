/* Given an integer n, return the number of prime numbers that are strictly less than n.

 

Example 1:

Input: n = 10
Output: 4
Explanation: There are 4 prime numbers less than 10, they are 2, 3, 5, 7.
Example 2:

Input: n = 0
Output: 0
Example 3:

Input: n = 1
Output: 0
 

Constraints:

0 <= n <= 5 * 106
*/
package main

import (
	"fmt"
)

func seiveOfEratosthenes(n int) int {
    prime := make([]int, n+1)
    count := 0;
    for i := 0; i<n+1; i++ {
        prime[i] = 1
    }

    for i:= 2; i*i <=n; i++ {
        if prime[i] == 1 {
            for j := i*i; j<=n; j+=i {
                prime[j] = 0;
            }
        }
    }

    for i:= 2; i <= n; i++ {
        if prime[i] == 1 {
            count++
        }
    }

    return count;
}

func countPrimes(n int) int {
    return seiveOfEratosthenes(n-1);
}

func main() {
	// Test cases
	fmt.Println(countPrimes(10)) // Output: 4
	fmt.Println(countPrimes(0)) // Output: 0
	fmt.Println(countPrimes(1)) // Output: 0
	fmt.Println(countPrimes(2)) // Output: 0
}