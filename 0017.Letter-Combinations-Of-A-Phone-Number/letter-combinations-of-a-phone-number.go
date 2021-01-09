package leetcode

var m = map[byte][]string{
	'2': []string{"a", "b", "c"},
	'3': []string{"d", "e", "f"},
	'4': []string{"g", "h", "i"},
	'5': []string{"j", "k", "l"},
	'6': []string{"m", "n", "o"},
	'7': []string{"p", "q", "r", "s"},
	'8': []string{"t", "u", "v"},
	'9': []string{"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	if digits == "" {
		return nil
	}

	res := []string{""}
	// Let all the numbers in digits have a chance to iterate.
	for i := 0; i < len(digits); i++ {
		temp := []string{}
		// Let each element in ret be able to add new letters.
		for j := 0; j < len(res); j++ {
			// Add the letter corresponding to digits[i] to the end of ret[j] multiple times
			for k := 0; k < len(m[digits[i]]); k++ {
				temp = append(temp, res[j]+m[digits[i]][k])
			}
		}

		res = temp
	}

	return res
}
