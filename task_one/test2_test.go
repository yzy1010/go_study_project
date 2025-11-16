package main

import "testing"

func TestSingleNumber(t *testing.T) {
	cases := []struct {
		nums []int
		want int
	}{
		{[]int{2, 2, 1}, 1},
		{[]int{4, 1, 2, 1, 2}, 4},
		{[]int{0, 1, 0}, 1},
		{[]int{7}, 7},
	}

	for _, c := range cases {
		got := singleNumber(c.nums)
		if got != c.want {
			t.Fatalf("singleNumber(%v) == %d, want %d", c.nums, got, c.want)
		}
	}
}
