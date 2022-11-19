# [2. Add Two Numbers](https://leetcode.com/problems/add-two-numbers/)

## topic

You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example:

```
Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
```
Explanation: 342 + 465 = 807.


## Topic idea

Two linked lists in reverse order are required to be added from the lower order, and the result is also output in reverse order, and the return value is the head node of the linked list in reverse order.

## Problem solving ideas

Need to pay attention to the various carry issues.

Extreme situations, such as
```
Input: (9 -> 9 -> 9 -> 9 -> 9) + (1 ->)
Output: 0 -> 0 -> 0 -> 0 -> 0 -> 1
```

In order to unify the processing method, a virtual head node can be established first, and the Next of this virtual head node points to the real head, so that the head does not need to be processed separately, just a while loop. In addition, the condition for judging the termination of the loop does not need to be p.Next! = nil, so the last bit needs additional calculations, and the loop termination condition should be p != nil.