package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestPickPeaks(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "PickPeaks Suite")
}

func dotest(array []int, posPeaks PosPeaks) {
	var ans = PickPeaks(array)
	Expect(ans).To(Equal(posPeaks))
}

var _ = Describe("Test Example", func() {
	It("should support finding peaks", func() {
		dotest(
			[]int{1, 2, 3, 6, 4, 1, 2, 3, 2, 1},
			PosPeaks{Pos: []int{3, 7}, Peaks: []int{6, 3}},
		)
	})

	It("should support finding peaks, but should ignore peaks on the edge of the array", func() {
		dotest(
			[]int{3, 2, 3, 6, 4, 1, 2, 3, 2, 1, 2, 3},
			PosPeaks{Pos: []int{3, 7}, Peaks: []int{6, 3}},
		)
	})

	It("should support finding peaks; if the peak is a plateau, it should only return the position of the first element of the plateau", func() {
		dotest(
			[]int{3, 2, 3, 6, 4, 1, 2, 3, 2, 1, 2, 2, 2, 1},
			PosPeaks{Pos: []int{3, 7, 10}, Peaks: []int{6, 3, 2}},
		)
	})

	It("should support finding peaks; if the peak is a plateau, it should only return the position of the first element of the plateau", func() {
		dotest(
			[]int{2, 1, 3, 1, 2, 2, 2, 2, 1},
			PosPeaks{Pos: []int{2, 4}, Peaks: []int{3, 2}},
		)
	})

	It("should support finding peaks, but should ignore peaks on the edge of the array", func() {
		dotest(
			[]int{2, 1, 3, 1, 2, 2, 2, 2},
			PosPeaks{Pos: []int{2}, Peaks: []int{3}},
		)
	})
})
