package chapter01

// Given a string, write a function to check if it is a permutation of palindrome.

// CheckPalindromePermutation verifies if a permutation is also a palindrome
// Time Complexity: O(N)
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

// CheckPalindromePermutation verifies if a permutation is also a palindrome
// This version uses only one loop
// Time Complexity O(N)
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
