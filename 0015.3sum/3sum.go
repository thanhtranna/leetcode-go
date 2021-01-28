package leetcode

import "sort"

func threeSum(nums []int) [][]int {
	// After sorting, you can search by law
	sort.Ints(nums)

	res := [][]int{}
	for i := range nums {
		if nums[i] > 0 {
			break
		}
		// Avoid adding duplicate results
		// i>0 is to prevent nums[i-1] from overflowing
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		l, r := i+1, len(nums)-1

		for l < r {
			s := nums[i] + nums[l] + nums[r]
			switch {
			case s < 0:
				// the smaller l needs to be larger
				l++
			case s > 0:
				// The larger r needs to be smaller
				r--
			default:
				res = append(res, []int{nums[i], nums[l], nums[r]})
				// To avoid repeated additions, both l and r need to be moved to different elements.
				l, r = next(nums, l, r)
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
