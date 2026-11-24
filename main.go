package main

import "fmt"

// Knapsack01 returns the maximum value packable into the given capacity,
// each item usable at most once. One DP row, iterated backwards so an
// item cannot be counted twice within its own round.
func Knapsack01(weights, values []int, capacity int) int {
	dp := make([]int, capacity+1) // dp[c]: best value within capacity c
	for i := range weights {
		for c := capacity; c >= weights[i]; c-- {
			if v := dp[c-weights[i]] + values[i]; v > dp[c] {
				dp[c] = v
			}
		}
	}
	return dp[capacity]
}

func main() {
	fmt.Println(Knapsack01([]int{1, 3, 4, 5}, []int{1, 4, 5, 7}, 7))
}
