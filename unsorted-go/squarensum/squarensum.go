package squarensum

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

//
//Complete the square sum function so that it squares each number passed into it and then sums the results together.
//
//For example, for [1, 2, 2] it should return 9
// 1 * 1 + 2 * 2 + 2 * 2 = 9

func SquareSum(numbers []int) int {
	return 0
}

func testSummaryProcessorWithMockStore(t *testing.T) {
	var _ = Describe("Sample Tests", func() {
		It("Testing [1,2]", func() { Expect(SquareSum([]int{1, 2})).To(Equal(5)) })
		It("Testing [0,3,4,5]", func() { Expect(SquareSum([]int{0, 3, 4, 5})).To(Equal(50)) })
		It("Testing []", func() { Expect(SquareSum([]int{})).To(Equal(0)) })
	})
}
