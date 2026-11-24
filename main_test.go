package main

import "testing"

func TestKnapsack01(t *testing.T) {
	cases := []struct {
		weights, values []int
		capacity        int
		want            int
	}{
		{[]int{1, 3, 4, 5}, []int{1, 4, 5, 7}, 7, 9},      // items 3+4
		{[]int{10, 20, 30}, []int{60, 100, 120}, 50, 220}, // classic
		{[]int{5}, []int{10}, 4, 0},                       // nothing fits
		{[]int{2, 2}, []int{3, 3}, 2, 3},                  // single use each
		{nil, nil, 10, 0},
		{[]int{1}, []int{5}, 0, 0},
	}
	for _, c := range cases {
		if got := Knapsack01(c.weights, c.values, c.capacity); got != c.want {
			t.Errorf("Knapsack01(%v,%v,%d) = %d, want %d", c.weights, c.values, c.capacity, got, c.want)
		}
	}
}
