package leetcode

import (
	"fmt"
	"testing"
)

/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

// @lc code=start
func TestTwoSum(t *testing.T) {
	if ans := fmt.Sprintf("%+v", twoSum([]int{2, 7, 11, 15}, 9)); ans != "[0 1]" {
		t.Errorf("[2,7,11,15] should be [0 1] but got %s", ans)
	}

	if ans := fmt.Sprintf("%+v", twoSum([]int{3, 2, 4}, 6)); ans != "[1 2]" {
		t.Errorf("[3,2,4], 6 should be [0 1] but got %s", ans)
	}

	if ans := fmt.Sprintf("%+v", twoSum([]int{3, 3}, 6)); ans != "[0 1]" {
		t.Errorf("[3,3], 6 should be [0 1] but got %s", ans)
	}

	if ans := fmt.Sprintf("%+v", twoSum([]int{2, 5, 5, 11}, 10)); ans != "[1 2]" {
		t.Errorf("[2,5,5,11], 10 should be [1 2] but got %s", ans)
	}
}

// @lc code=end
