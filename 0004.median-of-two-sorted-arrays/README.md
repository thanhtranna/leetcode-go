# [4. Median of Two Sorted Arrays](https://leetcode.com/problems/median-of-two-sorted-arrays/)


## Topic

There are two sorted arrays **nums1** and **nums2** of size m and n respectively.

Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).

You may assume **nums1** and **nums2** cannot be both empty.

**Example 1:**

    nums1 = [1, 3]
    nums2 = [2]
    
    The median is 2.0

**Example 2:**

    nums1 = [1, 2]
    nums2 = [3, 4]
    
    The median is (2 + 3)/2 = 2.5


## Topic idea


Given two ordered arrays nums1 and nums2 of size m and n.

Please find the median of these two ordered arrays, and the time complexity of the required algorithm is O(log(m + n)).

You can assume that nums1 and nums2 will not be empty at the same time.


## Problem solving ideas


-Given two ordered arrays, it is required to find the median of the ordered arrays after the two arrays are merged. The required time complexity is O(log (m+n)).
-The easiest way to think of this question is to merge the two arrays and then take the median. But the operation of merging ordered arrays is `O(max(n,m))`, which does not meet the meaning of the question. Seeing the time complexity of `log` given by the title, it is easy to think of binary search.
-Since we need to find the median of the arrays after the final merge, the total size of the two arrays is also known, so the middle position is also known. You only need to search for the position of the division in one array, and the position of the division in the other array can also be obtained. In order to minimize the time complexity, the array with the smaller length of the two arrays is searched binary.
-The key question is how to divide array 1 and array 2. In fact, it is how to split the array 1. First randomly generate a `midA` in two divisions. When does the dividing line meet the condition of the median? That is, the numbers on the left side of the line are less than the numbers on the right side, that is, `nums1[midA-1] ≤ nums2[midB] && nums2[midB-1] ≤ nums1[midA]`. If these conditions are not met, the cutting line needs to be adjusted. If `nums1[midA] <nums2[midB-1]`, it means that the number on the left of the line divided by `midA` is small, and the dividing line should be shifted to the right; if `nums1[midA-1]> nums2[midB] `, indicating that the number on the left side of the line divided by midA is large, and the dividing line should be shifted to the left. After many adjustments, the split line can always find a solution that meets the conditions.
-Assuming that the two dividing lines are now found, the subscripts of `array 1` on both sides of the dividing line are `midA-1` and `midA` respectively. The subscripts of `array 2` on both sides of the cut line are `midB-1` and `midB` respectively. Finally merge into the final array, if the length of the array is odd, then the median is `max(nums1[midA-1], nums2[midB-1])`. If the length of the array is even, the two numbers in the middle position are in order: `max(nums1[midA-1], nums2[midB-1])` and `min(nums1[midA], nums2[midB])`, Then the median is `(max(nums1[midA-1], nums2[midB-1]) + min(nums1[midA], nums2[midB])) / 2`. The icon is shown in the figure below:

    ![](https://img.halfrost.com/Leetcode/leetcode_4.png)