# [11. Container With Most Water](https://leetcode.com/problems/container-with-most-water/)

## Title
Given n non-negative integers a1, a2, ..., an, where each represents a point at coordinate (i, ai). n vertical lines are drawn such that the two endpoints of line i is at (i, ai) and (i, 0). Find two lines, which together with x-axis forms a container, such that the container contains the most water.

In other words, there are many vertical line segments on the x-axis at points 1, 2,..., n, the lengths are a1, a2, ..., an in order. Find two line segments to maximize the area between them and x. The area formula is Min(ai, aj) X |j-i|

## Problem solving ideas
The exhaustive method is O(n^2) complexity, which will trigger the time limit of the leetcode.

 The solution of O(n) complexity is to keep two pointers i and j; respectively point to the beginning and end of the length array. If ai is less than aj, move i backward (i++). Otherwise, move j forward (j--). If the current area is larger than the recorded area, replace it. The basis of this idea is that if the length of i is less than j, no matter how you move j, and the short board is at i, it is impossible to find a value larger than the currently recorded area. You can only find a new one that is larger than the currently recorded area. area.