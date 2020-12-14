package leetcode

/*
 * @lc app=leetcode id=66 lang=golang
 *
 * [66] Plus One
 */

// @lc code=start
func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i]++
		if digits[i] <= 9 {
			break
		}
		digits[i] = digits[i] % 10
		if i == 0 {
			newDigits := []int{}
			newDigits = append(newDigits, 1)
			newDigits = append(newDigits, digits...)
			digits = newDigits
		}
	}
	return digits
}

// @lc code=end
