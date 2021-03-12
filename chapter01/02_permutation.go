package chapter01

// Given two strings, write a method to decide if one is a permutation of the
// other

// PermutationBySort checks if string s is a permutation of string t sorting both strings and comparing them
func PermutationBySort(s, t string) bool {

	if len(s) != len(t) {
		return false
	}
	return SortString(s) == SortString(t)
}

// PermutationBySet checks if string s is a permutation of string t using a map
//
func PermutationBySet(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	set := make(map[rune]int)

	for _, v := range s {
		set[v]++
	}

	for _, v := range t {
		set[v]--

		if set[v] < 0 {
			return false
		}
	}

	return true
}
