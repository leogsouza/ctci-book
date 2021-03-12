package chapter01

import "sort"

type SortableString []rune

func (s SortableString) Less(a, b int) bool {
	return s[a] < s[b]
}

func (s SortableString) Swap(a, b int) {
	s[a], s[b] = s[b], s[a]
}

func (s SortableString) Len() int {
	return len(s)
}

func SortString(s string) string {
	ss := SortableString(s)
	sort.Sort(ss)

	return string(ss)
}