package chapter01

import "testing"

func TestURLify(t *testing.T) {
	tests := []struct {
		name string
		data string
		size int
		want string
	}{
		{
			name: "should return 'Mr%20John%20Smith' for string 'Mr John Smith    ', and size 13",
			data: "Mr John Smith    ",
			size: 13,
			want: "Mr%20John%20Smith",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := URLify(test.data, test.size)

			if got != test.want {
				t.Errorf("test %v failed: got %v, want %v", test.name, got, test.want)
			}
		})
	}
}
