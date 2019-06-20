package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("")
	s1 := []string{"Rum", "Rummy", "Russia"}
	// sort slice
	sort.Strings(s1)
	fmt.Println(s1)

	// reverse slice
	// sort.Sort(sort.Reverse(sort.StringSlice(s1)))
	// fmt.Println(s1)
	var s2 sort.StringSlice = s1
	fmt.Println(sort.Reverse(s2))

	// sort differnt type e.g. struct
	type person struct {
		name   string
		income int
	}
	p := []person{
		person{"abcd", 10},
		person{"else", 5},
		person{"someone", 8},
	}

	sort.Slice(p, func(i, j int) bool { return p[i].income < p[j].income })
	fmt.Println(p)
}
