package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"c", "b", "a"}
	sort.Strings(strs)
	fmt.Println("strs:", strs)

	ints := []int{2, 7, 6, 9}
	sort.Ints(ints)
	fmt.Println("ints:", ints)

	s := sort.IntsAreSorted(ints)
	fmt.Println("ints sorted:", s)
}
