package main

import (
	"fmt"
	"testing"
)

func simpletest(nums []int, expected int) {
	v := HighestRank(nums)
	if v != expected {
		fmt.Println("failed test:", v, "expecting", expected)
	}
}

func TestHighestRank(t *testing.T) {
	simpletest([]int{12, 10, 8, 12, 3, 3, 3, 10, 12}, 12)
	simpletest([]int{12, 10, 3, 12, 3, 3, 3, 10, 12}, 3)
	simpletest([]int{12}, 12)
	simpletest([]int{12, 14, 12}, 12)
}
