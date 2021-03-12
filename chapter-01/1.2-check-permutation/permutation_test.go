package permutation

import "testing"

func TestPermutation(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want bool
	}{
		{
			name: "should return true when s = 'god' and t = 'dog'",
			s:    "god",
			t:    "dog",
			want: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Permutation(test.s, test.t)

			if got != test.want {
				t.Errorf("test %v failed. got %v, want: %v", test.name, got, test.want)
			}
		})
	}
}
