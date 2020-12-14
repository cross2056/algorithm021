package leetcode

/*
 * @lc app=leetcode id=88 lang=golang
 *
 * [88] Merge Sorted Array
 */

// @lc code=start
func merge(nums1 []int, m int, nums2 []int, n int) {
	p1 := m - 1
	p2 := n - 1
	ptail := m + n - 1
	for p1 >= 0 && p2 >= 0 {
		if nums1[p1] > nums2[p2] {
			nums1[ptail] = nums1[p1]
			p1--
		} else {
			nums1[ptail] = nums2[p2]
			p2--
		}
		ptail--
	}
	for p2 >= 0 && ptail >= 0 {
		nums1[ptail] = nums2[p2]
		p2--
		ptail--
	}
}

// @lc code=end
