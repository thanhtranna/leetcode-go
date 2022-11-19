# [6. ZigZag Conversion](https://leetcode.com/problems/zigzag-conversion/)

## Title

The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)

```text
P A H N
A P L S I I G
Y I R
```

And then read line by line: "PAHNAPLSIIGYIR"
Write the code that will take a string and make this conversion given a number of rows:

```go
func convert(text string, nRows int) string
```

convert("PAYPALISHIRING", 3) should return "PAHNAPLSIIGYIR".

## Problem solving ideas

After entering "ABCDEFGHIJKLMNOPQRSTUVWXYZ" and parameter 5, the answer is "AGMSYBFHLNRTXZCEIKOQUWDJPV",
According to the placement method of the topic, you can get:

```text
A I Q Y
B HJ PR XZ
C G K O S W
DF LN TV
E M U
```

As you can see, the index number of each line of characters in the original string is

1. Row 0, 0, 8, 16, 24
1. Row 1, 1, 7,9, 15,17, 23,25
1. Line 2, 2, 6, 10, 14, 18, 22
1. Line 3, 3, 5, 11,13, 19,21
1. Line 4, 4, 12, 20

Let p=numRows×2-2, the following law can be summarized

1. 0 rows, 0×p, 1×p,...
1. Row r, r, 1×p-r, 1×p+r, 2×p-r, 2×p+r,...
1. The last line, numRow-1, numRow-1+1×p, numRow-1+2×p,...

You only need to program to process each row in turn.