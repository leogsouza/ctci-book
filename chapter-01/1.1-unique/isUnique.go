package unique

//  Implement an algorithm to determine if a string has all unique characters.
// What if you cannot use additional data structures?

// Unique checks if the all characters from a string are unique
// Good approach using map as auxiliary data structure
// Time complexity: O(N)
func Unique(s string) bool {
	chars := make(map[rune]bool)

	for _, c := range s {
		if chars[c] {
			return false
		}

		chars[c] = true
	}

	return true
}

// NaiveUnique checks if the all characters from a string are unique
// Brute force approach - O(N^2)
func NaiveUnique(s string) bool {

	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				return false
			}
		}
	}
	return true
}
