# [13. Roman to Integer](https://leetcode.com/problems/roman-to-integer/)

## Title
Given a roman numeral, convert it to an integer.

Input is guaranteed to be within the range from 1 to 3999.

## Problem solving ideas
This question is an inverse transformation of [12. Integer to Roman](./Algorithms/0012.integer-to-roman). ~~Same problem-solving ideas~~

For this question, the most critical information is
> Right plus and left minus, the left minus number must be one digit, for example, 8 is written as VIII, not IIX.

Problem solving ideas
1. Process the string from right to left.
1. When the number represented by the current character is smaller than the character on the right, the total number represented by the current character is subtracted.
1. Otherwise, add the total number represented by the current character.
## to sum up
Grasp key information and avoid stereotypes.