/*
	You are given a sorted array ‘arr’ of length ‘n’, which contains positive integer positions of ‘n’ gas stations on the X-axis.

You are also given an integer ‘k’.

You have to place 'k' new gas stations on the X-axis.

You can place them anywhere on the non-negative side of the X-axis, even on non-integer positions.

Let 'dist' be the maximum value of the distance between adjacent gas stations after adding 'k' new gas stations.

Find the minimum value of dist.

Example:
Input: ‘n' = 7 , ‘k’=6, ‘arr’ = {1,2,3,4,5,6,7}.

Answer: 0.5

Explanation:
We can place 6 gas stations at 1.5, 2.5, 3.5, 4.5, 5.5, 6.5.

Thus the value of 'dist' will be 0.5. It can be shown that we can't get a lower value of 'dist' by placing 6 gas stations.

Note:
You will only see 1 or 0 in the output where:

	1 represents your answer is correct.
	0 represents your answer is wrong.

Answers within 10^-6 of the actual answer will be accepted.
Detailed explanation ( Input/output format, Notes, Images )
Sample Input 1:
5 4
1 2 3 4 5

Expected Answer:
0.5

Output Printed On Console:
1

Explanation of sample output 1:
k = 4, arr = {1,2,3,4,5}

One of the possible ways to place 4 gas stations is {1,1.5,2,2.5,3,3.5,4,4.5,5}

Thus the maximum difference between adjacent gas stations is 0.5.

Hence, the value of ‘dist’ is 0.5. It can be shown that there is no possible way to add 4 gas stations in such a way that the value of ‘dist’ is lower than this. (Note: 1 will be printed in the output for the correct answer).

Sample Input 2:
10 1
1 2 3 4 5 6 7 8 9 10

Expected Answer:
1

Output Printed On Console:
1

Explanation of sample output 2:
k = 1, arr = {1,2,3,4,5,6,7,8,9,10}

One of the possible ways to place 1 gas station is {1,1.5,2,3,4,5,6,7,8,9,10}

Thus the maximum difference between adjacent gas stations is still 1.

Hence, the value of ‘dist’ is 1. It can be shown that there is no possible way to add 1 gas station in such a way that the value of ‘dist’ is lower than this. (Note: 1 will be printed in the output for the correct answer).

Expected Time Complexity
Try solving this in O(n*log(A)), where A is the maximum value in array 'arr'.

Constraints:
2 <= n <= 10^5
1 <= k <= 10^6
1 <= arr[i] <= 10^9

Time Limit: 1 sec
*/
package main

import (
	"container/heap"
)

// A MaxHeap to store pairs of (distance, index)
type MaxHeap []pair

type pair struct {
	distance float64
	index    int
}

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i].distance > h[j].distance }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(pair))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}

func minimiseMaxDistance(arr []int, k int) float64 {
	h := &MaxHeap{}
	heap.Init(h)
	howMany := make([]int, len(arr)-1)

	for i := 0; i < len(arr)-1; i++ {
		heap.Push(h, pair{distance: float64(arr[i+1] - arr[i]), index: i})
	}

	for i := 1; i <= k; i++ {
		tp := heap.Pop(h).(pair)
		howMany[tp.index]++
		dist := float64(arr[tp.index+1] - arr[tp.index])
		ans := dist / float64(howMany[tp.index]+1)
		heap.Push(h, pair{distance: ans, index: tp.index})
	}

	return (*h)[0].distance
}

func main() {
	println(minimiseMaxDistance([]int{1, 2, 3, 4, 5}, 4))               
	println(minimiseMaxDistance([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 1)) 
}
