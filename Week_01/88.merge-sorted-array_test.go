package leetcode

import (
	"fmt"
	"testing"
)

func TestMerge(t *testing.T) {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3
	merge(nums1, m, nums2, n)
	if ans := fmt.Sprintf("%+v", nums1); ans != "[1 2 2 3 5 6]" {
		t.Errorf("nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3 should be [1 2 2 3 5 6] but got %s", ans)
	}
}
