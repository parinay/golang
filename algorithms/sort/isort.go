package main

import "fmt"

func insertSort(v []int) {
	fmt.Println(v)
	for j := 1; j < len(v); j++ {
		key := v[j]
		i := j - 1

		for i >= 0 && v[i] > key {
			v[i+1] = v[i]
			i--
		}

		v[i+1] = key
	}
	fmt.Println(v)
}
