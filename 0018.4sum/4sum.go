package leetcode

import "sort"

func fourSum(nums []int, target int) [][]int {
	res := [][]int{}

	// After sorting, you can search by law
	sort.Ints(nums)

	for i := 0; i < len(nums)-3; i++ {
		// Avoid adding duplicate results
		// i>0 is to prevent nums[i-1] from overflowing
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i + 1; j < len(nums)-2; j++ {
			// Avoid adding duplicate results
			// nums[j-1] will not overflow
			// But if you don't add j> i+1, you will miss the answers where i and j are the same value
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			l, r := j+1, len(nums)-1
			for l < r {
				s := nums[i] + nums[j] + nums[l] + nums[r]
				switch {
				case s < target:
					l++
				case s > target:
					r--
				default:
					res = append(res, []int{nums[i], nums[j], nums[l], nums[r]})
					l, r = next(nums, l, r)
				}
			}
		}
	}

	return res
}

func next(nums []int, l, r int) (int, int) {
	for l < r {
		switch {
		case nums[l] == nums[l+1]:
			l++
		case nums[r] == nums[r-1]:
			r--
		default:
			l++
			r--
			return l, r
		}
	}

	return l, r
}
