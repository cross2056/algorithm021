package leetcode

import (
	"strconv"
	"testing"
)

func TestTrap(t *testing.T) {
	heights := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	if ans := trap(heights); ans != 6 {
		t.Errorf("[0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1] should be 6 but got %s", strconv.Itoa(ans))
	}
}
