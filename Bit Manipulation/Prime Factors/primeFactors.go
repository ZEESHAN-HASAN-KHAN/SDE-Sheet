/* Given a number N. Find its unique prime factors in increasing order.
 

Example 1:

Input: N = 100
Output: 2 5
Explanation: 2 and 5 are the unique prime
factors of 100.
Example 2:

Input: N = 35
Output: 5 7
Explanation: 5 and 7 are the unique prime
factors of 35.
 

Your Task:
You don't need to read or print anything. Your task is to complete the function AllPrimeFactors() which takes N as input parameter and returns a list of all unique prime factors of N in increasing order.

 

Expected Time Complexity: O(N)
Expected Space Complexity: O(N)
 

Constraints:
1 <= N  <= 106
*/
package main

import (
	"fmt"
)

func seiveOfEratosthenes(n int) []int {
	prime := make([]bool, n+1)
	for i := 0; i < n+1; i++ {
		prime[i] = true
	}

	for p := 2; p*p <= n; p++ {
		if prime[p] == true {
			for i := p * p; i <= n; i += p {
				prime[i] = false
			}
		}
	}

	primes := []int{}
	for p := 2; p <= n; p++ {
		if prime[p] {
			primes = append(primes, p)
		}
	}
	return primes
}

func AllPrimeFactors(N int) []int {
	primes := seiveOfEratosthenes(N)
	factors := []int{}
	for _, prime := range primes {
		if N%prime == 0 {
			factors = append(factors, prime)
			for N%prime == 0 {
				N /= prime
			}
		}
	}
	return factors
}

func main() {
	fmt.Println(AllPrimeFactors(100))
	fmt.Println(AllPrimeFactors(35))
}