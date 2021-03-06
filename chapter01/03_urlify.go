package chapter01

// URLify replaces all spaces characters in a string with '%20'
// Time Complexity: O(N)
func URLify(s string, n int) string {
	spaces := 0
	for i := 0; i < n; i++ {
		if s[i] == ' ' {
			spaces++
		}
	}

	index := n + spaces*2
	runes := make([]rune, n+spaces*2)

	for i := n - 1; i >= 0; i-- {
		if s[i] == ' ' {
			runes[index-1] = '0'
			runes[index-2] = '2'
			runes[index-3] = '%'
			index -= 3
		} else {
			runes[index-1] = rune(s[i])
			index--
		}
	}

	return string(runes)
}
