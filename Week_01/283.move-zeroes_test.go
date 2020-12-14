package leetcode

import (
	"fmt"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	input := []int{0, 1, 0, 3, 12}
	moveZeroes(input)
	if ans := fmt.Sprintf("%+v", input); ans != "[1 3 12 0 0]" {
		t.Errorf("[0,1,0,3,12] should be [1,3,12,0,0] but got %s", ans)
	}
}
