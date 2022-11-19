package leetcode

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// Assuming the length of nums1 is small
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}

	low, high, k, nums1Mid, nums2Mid := 0, len(nums1), (len(nums1)+len(nums2)+1)>>1, 0, 0
	for low <= high {
		// nums1: ……………… nums1[nums1Mid-1] | nums1[nums1Mid] ……………………
		// nums2: ……………… nums2[nums2Mid-1] | nums2[nums2Mid] ……………………
		nums1Mid = low + (high-low)>>1 // The right side of the dividing line is mid, and the left side of the dividing line is mid-1
		nums2Mid = k - nums1Mid
		if nums1Mid > 0 && nums1[nums1Mid-1] > nums2[nums2Mid] { // The dividing line in nums1 is too many, so move to the left
			high = nums1Mid - 1
		} else if nums1Mid != len(nums1) && nums1[nums1Mid] < nums2[nums2Mid-1] { // The dividing line in nums1 is missing, so move to the right
			low = nums1Mid + 1
		} else {
			// Find a suitable division, need to output the final result
			// Divided into 2 cases of odd and even
			break
		}
	}
	midLeft, midRight := 0, 0
	if nums1Mid == 0 {
		midLeft = nums2[nums2Mid-1]
	} else if nums2Mid == 0 {
		midLeft = nums1[nums1Mid-1]
	} else {
		midLeft = max(nums1[nums1Mid-1], nums2[nums2Mid-1])
	}
	if (len(nums1)+len(nums2))&1 == 1 {
		return float64(midLeft)
	}
	if nums1Mid == len(nums1) {
		midRight = nums2[nums2Mid]
	} else if nums2Mid == len(nums2) {
		midRight = nums1[nums1Mid]
	} else {
		midRight = min(nums1[nums1Mid], nums2[nums2Mid])
	}
	return float64(midLeft+midRight) / 2
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a > b {
		return a
	}

	return b
}
