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
