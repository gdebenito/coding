package main

// Complete the method which returns the number which is most frequent in the
// given input array. If there is a tie for most frequent number, return the
// largest number among them.
func HighestRank(nums []int) int {
	m := make(map[int]int)

	for _, v := range nums {
		m[v] = m[v] + 1
	}

	maxN := 0
	maxR := 0

	for k, v := range m {
		if v > maxR {
			maxN = k
			maxR = v
		}
		if v == maxR && k > maxN {
			maxN = k
		}
	}

	return maxN

}
