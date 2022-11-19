package leetcode

func longestPalindrome(s string) string {
	if len(s) < 2 { // It must be a palindrome, return directly
		return s
	}

	// The first character index of the longest palindrome, and the length of the longest palindrome
	begin, maxLen := 0, 1

	// in the for loop
	// b represents the index number of the **first** character of the palindrome,
	// e represents the index number of the **tail** character of the palindrome,
	// i represents the index number of the first character of the palindrome `center paragraph`
	// in each for loop
	// Start with i and use the same characteristics of all characters in the `center section` to make b and e point to the beginning and end of the `center section` respectively
	// Then expand from `center section` to both sides, let b and e respectively point to the beginning and end of the longest palindrome centered on this `center section`
	for i := 0; i < len(s); { // Start looking for the longest palindrome with s[i] as the first character of `the middle paragraph`.
		if len(s)-i <= maxLen/2 {
			// Because i is the index number of the first character of the palindrome `center paragraph`
			// Assuming that the length of the longest palindrome that can be found at this time is l, then l <= (len(s)-i)*2-1
			// If len(s)-i <= maxLen/2 holds
			// Then, l <= maxLen-1
			// Then, l <maxLen
			// So, there is no need to look any further.
			break
		}

		b, e := i, i
		for e < len(s)-1 && s[e+1] == s[e] {
			e++
			// After the loop, s[b:e+1] is a string of the same string
		}

		// The first character of the next palindrome will only be s[e+1]
		// prepare for the next loop
		i = e + 1

		for e < len(s)-1 && b > 0 && s[e+1] == s[b-1] {
			e++
			b--
			// After the loop, s[b:e+1] is the longest palindrome that can be found this time.
		}

		newLen := e + 1 - b
		// If you create a new record, update the record
		if newLen > maxLen {
			begin = b
			maxLen = newLen
		}
	}

	return s[begin : begin+maxLen]
}
