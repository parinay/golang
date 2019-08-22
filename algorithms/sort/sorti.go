package main

import (
	"fmt"
	"sort"

	"github.com/parinay/golang/algorithms/sort/global"
)

// ByAge implements sort.Interface based on the Age field.
type ByAge []global.Family

func (a ByAge) Len() int { return len(a) }

func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func (a ByAge) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func sorti(f []global.Family) {

	sort.Sort(ByAge(f))
	fmt.Println(f)
}
