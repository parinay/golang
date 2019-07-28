package main

import (
	"fmt"
	"reflect"
)

func main() {
	map1 := make(map[int]int)
	map2 := make(map[int]int)

	map1[1] = 10
	map2[1] = 10

	eq := reflect.DeepEqual(map1, map2)

	if eq {
		fmt.Println("Maps are equal")
	} else {
		fmt.Println("Maps are not equal")
	}
}
