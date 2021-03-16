package chapter01

func CheckPalindromePermutation(s string) bool {
	letters := make(map[rune]int)

	for _, c := range s {
		if c == ' ' {
			continue
		}
		letters[c]++
	}

	isOdd := false

	for _, n := range letters {
		if n%2 != 0 {
			if isOdd {
				return false
			}
			isOdd = true
		}
	}
	return true
}

func CheckPalindromePermutationSimplified(s string) bool {
	letters := make(map[rune]int)
	odd := 0

	for _, c := range s {
		if c == ' ' {
			continue
		}
		letters[c]++

		if letters[c]%2 != 0 {
			odd++
		} else {
			odd--
		}
	}

	return odd <= 1
}
