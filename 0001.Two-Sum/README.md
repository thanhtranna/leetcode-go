# [1. Two Sum](https://leetcode.com/problems/two-sum/)

## Topic

Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:

```
Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
```



## Topic idea

Find the number whose sum of 2 numbers is equal to the given value in the array, and return the subscript of the 2 numbers in the array as a result.

## Problem solving ideas

The optimal time complexity for this problem is O(n).

Scan the array in order, for each element, find the other half of the number in the map that can combine the given value, if found, just return the subscripts of the 2 numbers. If you can't find it, store the number in the map, wait for the "other half" number to be scanned, and then take it out to return the result.