package leetcode

import (
	"fmt"
	"testing"
)

func TestRotateArray(t *testing.T) {
	nums := []int{-1}
	rotate(nums, 2)
	if ans := fmt.Sprintf("%+v", nums); ans != "[5 6 7 1 2 3 4]" {
		t.Errorf("[1,2,3,4,5,6,7] 3 should be [5 6 7 1 2 3 4] but got %s", ans)
	}
}
