/* You are given a positive number N. Using the concept of Sieve, compute its prime factorisation.

Example:

Input: 
N = 12246
Output: 
2 3 13 157
Explanation: 
2*3*13*157 = 12246 = N.
Your Task:

Comple the function findPrimeFactors(), which takes a positive number N as input and returns a vector consisting of prime factors. You should implement Sieve algorithm to solve this problem.

 

Expected Time Complexity: O(Nlog(log(N)).

Expected Auxiliary Space: O(N).

Constraints:

1<=N<=2*105
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

func findPrimeFactors(n int) []int {
	var (
		ans []int
		onlyPrime = seiveOfEratosthenes(n)
		iterator = 0
	)

	for n > 1 && iterator < len(onlyPrime) {
		for n % onlyPrime[iterator] == 0 {
			ans = append(ans, onlyPrime[iterator])
			n /= onlyPrime[iterator]
		}
		iterator++
	}

	return ans
}


func main() {
	fmt.Println(findPrimeFactors(12246)) // Output: [2 3 13 157]
}