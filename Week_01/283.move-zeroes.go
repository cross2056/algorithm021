package leetcode

/*
 * @lc app=leetcode id=283 lang=golang
 *
 * [283] Move Zeroes
 */

// @lc code=start
func moveZeroes(nums []int) {
	if nums == nil || len(nums) == 0 {
		return
	}
	instPostion := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[instPostion] = nums[i]
			instPostion++
		}
	}

	for i := instPostion; i < len(nums); i++ {
		nums[i] = 0
	}
}

// @lc code=end
