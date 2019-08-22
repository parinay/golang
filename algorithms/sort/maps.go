package main

import (
	"fmt"
	"sort"
)

func maps() {
	m := map[string]int{"Alice": 2, "Cecil": 1, "Bob": 3}

	keys := make([]string, 0, len(m))
	values := make([]int, 0, len(m))

	for k, v := range m {
		keys = append(keys, k)
		values = append(values, v)
	}

	// sorting []slice keys
	sort.Strings(keys)
	sort.Ints(values)

	// print key, value ordered by key
	for _, k := range keys {
		fmt.Println(k, m[k])
	}
	for k, v := range values {
		fmt.Println(k, v)
	}
	// print maps ( default they're sorted by key
	fmt.Println(m)

}
