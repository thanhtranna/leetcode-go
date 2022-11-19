# [3. Longest Substring Without Repeating Characters](https://leetcode.com/problems/longest-substring-without-repeating-characters/)

## topic

Given a string, find the length of the longest substring without repeating characters.



Example 1:

```c
Input: "abcabcbb"
Output: 3 
Explanation: The answer is "abc", with the length of 3. 
```

Example 2:

```c
Input: "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.
```

Example 3:

```c
Input: "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3. 
             Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
```

## Topic idea


Search for the longest substring without repeated letters in a string.

## Problem solving ideas

This question is similar to question 438, question 3, question 76, and question 567, and the ideas used are all "sliding windows".

The right boundary of the sliding window keeps moving to the right, as long as there are no repeated characters, the window boundary continues to expand to the right. Once the repeated characters appear, the left boundary needs to be reduced until the repeated characters move out of the left boundary, and then continue to move the right boundary of the sliding window. By analogy, each movement needs to calculate the current length and determine whether the maximum length needs to be updated, and the final maximum value is what is required in the question.