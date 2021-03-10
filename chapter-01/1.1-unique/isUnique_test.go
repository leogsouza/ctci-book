package unique

import "testing"

func TestUnique(t *testing.T) {
	tests := []struct {
		name string
		data string
		want bool
	}{
		{
			name: "should return true for string 'length'",
			data: "length",
			want: true,
		},
		{
			name: "should return false for string 'estimated'",
			data: "estimated",
			want: false,
		},
		{
			name: "should return true for string '|\\/#$%^&*'",
			data: "|\\/#$%^&*",
			want: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Unique(test.data)
			if got != test.want {
				t.Errorf("test: %v failed. got %v, want %v", test.name, got, test.want)
			}
		})
	}
}
