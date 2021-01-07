package leetcode

import (
	"bytes"
)

func convert(s string, numRows int) string {
	if numRows == 1 || len(s) <= numRows {
		return s
	}
	res := bytes.Buffer{}
	// p pace
	p := numRows*2 - 2

	// Process the first line
	for i := 0; i < len(s); i += p {
		res.WriteByte(s[i])
	}

	// Process the middle row
	for r := 1; r <= numRows-2; r++ {
		// Add the first character of the r line
		res.WriteByte(s[r])

		for k := p; k-r < len(s); k += p {
			res.WriteByte(s[k-r])
			if k+r < len(s) {
				res.WriteByte(s[k+r])
			}
		}
	}

	// Process the last line
	for i := numRows - 1; i < len(s); i += p {
		res.WriteByte(s[i])
	}

	return res.String()
}
