package chapter01

import "testing"

func TestCheckPalindromePermutation(t *testing.T) {
	tests := []struct {
		name string
		word string
		want bool
	}{
		{
			name: "'taco cat' is a palindrome permutation",
			word: "taco cat",
			want: true,
		},
		{
			name: "'tacos cat' is not a palindrome permutation",
			word: "tacos cat",
			want: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := CheckPalindromePermutation(test.word)

			if got != test.want {
				t.Errorf("CheckPalindromePermutation(%s) got %v, want %v", test.word, got, test.want)
			}
		})
	}
}
