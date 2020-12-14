package leetcode

/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

// @lc code=start
// Solution1
func twoSum(nums []int, target int) []int {
	hash := map[int]int{}
	for i, n := range nums {
		if j, ok := hash[target-n]; ok {
			return []int{j, i}
		}
		hash[n] = i
	}

	return nil
}

// Solution2
// func twoSum(nums []int, target int) []int {
// 	for i := 0; i < len(nums)-1; i++ {
// 		for j := 1; j < len(nums); j++ {
// 			if i == j {
// 				continue
// 			}
// 			if nums[i]+nums[j] == target {
// 				return []int{i, j}
// 			}
// 		}
// 	}
// 	return nil
// }

// @lc code=end
