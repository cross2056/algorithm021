package leetcode

import (
	"fmt"
	"testing"
)

func TestPlusOne(t *testing.T) {
	if ans := fmt.Sprintf("%+v", plusOne([]int{4, 3, 2, 1})); ans != "[4 3 2 2]" {
		t.Errorf("[4, 3, 4, 1] should be [4 3 2 2] but got %s", ans)
	}
}

// @lc code=end
