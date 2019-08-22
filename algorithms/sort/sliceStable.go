package main

import (
	"fmt"
	"sort"

	"github.com/parinay/golang/algorithms/sort/global"
)

func sliceStable(f []global.Family) {

	sort.SliceStable(f, func(i, j int) bool {
		return f[i].Age < f[j].Age
	})

	fmt.Println("Family ", f)
}
