package leetcode

import (
	"strconv"
	"testing"
)

/*
 * @lc app=leetcode id=26 lang=golang
 *
 * [26] Remove Duplicates from Sorted Array
 */

// @lc code=start
func TestRemoveDuplicates(t *testing.T) {

	if ans := removeDuplicates([]int{1, 1, 2}); ans != 2 {
		t.Errorf("[1, 1, 2] should be 2 but got %s", strconv.Itoa(ans))
	}
	if ans := removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}); ans != 5 {
		t.Errorf("[0, 0, 1, 1, 1, 2, 2, 3, 3, 4] should be 5 but got %s", strconv.Itoa(ans))
	}
}

// @lc code=end
