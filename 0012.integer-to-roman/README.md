# [12. Integer to Roman](https://leetcode.com/problems/integer-to-roman/)

## Title
Given an integer, convert it to a roman numeral.

Input is guaranteed to be within the range from 1 to 3999.

The following is about [Roman numerals] on Wikipedia (https://zh.wikipedia.org/zh-cn/%E7%BD%97%E9%A9%AC%E6%95%B0%E5%AD%97) explanation of.

There are 7 Roman numerals, namely Ⅰ (1), Ⅴ (5), Ⅹ (10), Ⅼ (50), Ⅽ (100), Ⅾ (500) and Ⅿ (1000). According to the following rules, any positive integer can be expressed. It should be noted that there is no "0" in Roman numerals, which has nothing to do with the carry system. It is generally believed that Roman numerals are only used for counting, not for calculation.

1. Repeat several times: A Roman numeral is repeated several times, which means several times the number.
1. Right plus left minus:
    1. Write the smaller Roman numeral to the right of the larger Roman numeral, which means that the larger number plus the smaller number.
    1. Write a smaller Roman numeral to the left of a larger Roman numeral to indicate that the larger number decreases the number.
    1. There are restrictions on the numbers that can be subtracted from the left, only limited to I, X, and C. For example, 45 cannot be written as VL, only XLV
    1. However, when subtracting to the left, one bit value cannot be crossed. For example, 99 cannot be represented by IC (100-1), but XCIX ([100-10]+[10-1]). (Equivalent to Arabic numerals for each digit separately.)
    1. The left minus number must be one digit, for example, 8 is written as VIII, not IIX.
    1. The number added to the right cannot exceed three consecutive digits. For example, 14 is written as XIV instead of XIIII. (See the "Digital Restrictions" item below.)
1. Add line and multiply thousands:
    1. Adding a horizontal line or subscript Ⅿ above the Roman numeral means that this number is multiplied by 1000, which is 1000 times the original number.
    1. Similarly, if there are two horizontal lines above, it is 1000000 times the original number.
1. Digital restrictions:
    1. The same number can only appear three times in a row, such as 40 can not be expressed as XXXX, but as XL.

## Problem solving ideas
The description of Roman numerals seems very complicated. But for programming, the most critical one is
> However, when subtracting to the left, one bit value cannot be crossed. For example, 99 cannot be represented by IC (100-1), but XCIX ([100-10]+[10-1]). (Equivalent to Arabic numerals for each digit separately.)

Explain that Roman numerals, like Arabic numerals, are represented by bits. As long as the Roman numerals corresponding to the Arabic numerals on "one ten hundred thousand" are listed, you can freely combine to get the Roman numerals.

## to sum up
Table-driven method is a programming mode mentioned in "Code Encyclopedia", which can deal with complex query situations.