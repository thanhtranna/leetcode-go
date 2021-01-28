# [5. Longest Palindromic Substring](https://leetcode.com/problems/longest-palindromic-substring/)

## Title
Given a string s, find the longest palindromic substring in s. You may assume that the maximum length of s is 1000.

Example:
```
Input: "babad"
Output: "bab"
Note: "aba" is also a valid answer.
```
Example:
```
Input: "cbbd"
Output: "bb"
```
## Problem solving ideas
The title requires finding the longest palindrome in the string.
Of course, we can use the following program to determine whether the string s[i:j+1] is a palindrome.
```go
func isPalindromic(s *string, i, j int) bool {
    for i< j {
        if (*s)[i] != (*s)[j] {
            return false
        }
        i++
        j--
    }
    return true
}
```
However, this does not take advantage of the characteristics of palindrome, assuming that the length of palindrome is l and x is any character
1. When l is an odd number, the `center section` of the palindrome will only be "x", or "xxx", or "xxxxx", or...
1. When l is an even number, the `center section` of the palindrome will only be, "xx", or "xxxx", or "xxxxxx", or...
1. Any two palindrome substrings of the same string will not overlap each other in the middle section.

## to sum up
Making full use of the characteristics of the search object can speed up the search.