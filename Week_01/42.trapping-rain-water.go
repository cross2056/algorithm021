package leetcode

/*
 * @lc app=leetcode id=42 lang=golang
 *
 * [42] Trapping Rain Water
 */

// @lc code=start
func trap(height []int) int {
	if len(height) < 3 {
		return 0
	}
	stack := []int{}
	sum := 0
	for i := 0; i < len(height); i++ {
		if len(stack) == 0 {
			stack = append(stack, i)
			continue
		}
		if height[i] < height[stack[len(stack)-1]] {
			stack = append(stack, i)
		} else {
			for len(stack) > 0 && height[i] > height[stack[len(stack)-1]] {
				midIndx := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				if len(stack) > 0 {
					h := min(height[stack[len(stack)-1]], height[i]) - height[midIndx]
					w := i - stack[len(stack)-1] - 1
					sum += h * w
				}
			}
			stack = append(stack, i)
		}
	}
	return sum
}

// @lc code=end
