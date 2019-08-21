package main

import (
	"fmt"
	"sort"
)

func bsearch(a []int, x int) bool {
	i := sort.Search(len(a), func(i int) bool {
		return a[i] >= x
	})

	if i < len(a) && a[i] == x {
		fmt.Printf("Found %d at index %d in %v\n", x, i, a)
		return true
	} else {
		fmt.Printf("%d not found in %v\n", x, a)
		return false
	}
}

func bsearchStr(a []string, y string) bool {
	i := sort.Search(len(a), func(i int) bool { return y <= a[i] })
	if i < len(a) && a[i] == y {
		fmt.Printf("Found %s at index %d in %v\n", y, i, a)
		return true
	} else {
		fmt.Printf("Not Found %s in %v\n", y, a)
	}
	return false
}
