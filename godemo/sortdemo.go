package main

import (
	"sort"
	"fmt"
)

type byLength []string

func (s byLength) Len() int {
	return len(s)
}

func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// sort by length
func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func main() {
	fruits := []string{"1", "23", "kiwi", "1s3d", "4"}
	s := byLength(fruits)
	sort.Sort(s)
	fmt.Println(s)
}
