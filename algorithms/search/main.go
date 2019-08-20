package main

import "fmt"

func main() {

	s1 := []string{"p", "a", "r", "i", "n", "y"}
	fmt.Println("Find ", "y", "in ", s1)
	if linear1(s1, "y") != len(s1) {
		fmt.Println("Found !")
	} else {
		fmt.Println("Not Found !")
	}

	s2 := []string{"v", "a", "i", "b", "h", "a", "w", "i"}
	fmt.Println("Find ", "p", "in ", s2)
	if !linear2(s2, "p") {
		fmt.Println("Not found")
	} else {
		fmt.Println("Found!")
	}
	a := []int{1, 3, 6, 10, 15, 21, 28, 36, 45, 55}
	x := 6
	bsearch(a, x)
}
