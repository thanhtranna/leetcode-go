package leetcode

func isMatch(s, p string) bool {
	sSize := len(s)
	pSize := len(p)

	dp := make([][]bool, sSize+1)
	for i := range dp {
		dp[i] = make([]bool, pSize+1)
	}

	/* dp[i][j] represents whether s[:i] can match p[:j] */

	dp[0][0] = true
	/**
	* According to the setting of the title, "" can match "a*b*c*"
	* So, you need to set the corresponding dp to true
	 */
	for j := 1; j < pSize && dp[0][j-1]; j += 2 {
		if p[j] == '*' {
			dp[0][j+1] = true
		}
	}

	for i := 0; i < sSize; i++ {
		for j := 0; j < pSize; j++ {
			if p[j] == '.' || p[j] == s[i] {
				/* p[j] and s[i] can match, so as long as the previous matches, it can match here */
				dp[i+1][j+1] = dp[i][j]
			} else if p[j] == '*' {
				/* At this point, the match of p[j] is related to the content of p[j-1]. */
				if p[j-1] != s[i] && p[j-1] != '.' {
					/**
					* p[j] cannot match s[i]
					* p[j-1:j+1] can only be treated as ""
					 */
					dp[i+1][j+1] = dp[i+1][j-1]
				} else {
					/**
					* p[j] matches s[i]
					* p[j-1;j+1] as "x*", there are three interpretations
					 */
					dp[i+1][j+1] = dp[i+1][j-1] || /* "x*" is interpreted as "" */
						dp[i+1][j] || /* "x*" is interpreted as "x" */
						dp[i][j+1] /* "x*" is interpreted as "xx..." */
				}
			}
		}
	}

	return dp[sSize][pSize]
}
