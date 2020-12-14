package leetcode

/*
 * @lc app=leetcode id=189 lang=golang
 *
 * [189] Rotate Array
 */

// @lc code=start
// solution 1
func rotate(nums []int, k int) {
	k %= len(nums)
	reverse(nums, 0, len(nums)-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, len(nums)-1)
}

func reverse(nums []int, start int, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

// solution 2
// func rotate(nums []int, k int) {
// 	for t := 0; t < k; t++ {
// 		last := nums[len(nums)-1]
// 		for i := 0; i < len(nums); i++ {
// 			tmp := nums[i]
// 			nums[i] = last
// 			last = tmp
// 		}
// 	}
// }

// public void reverse(int[] nums, int start, int end) {
// 	while (start < end) {
// 		int temp = nums[start];
// 		nums[start] = nums[end];
// 		nums[end] = temp;
// 		start++;
// 		end--;
// 	}
// }

// @lc code=end
