package main

import "fmt"

func main() {
	m := make(map[string]string)

	m["k1"] = "Bengaluru"
	m["k2"] = "Chennai"

	fmt.Println("map:", m)

	v1 := m["k1"]

	fmt.Println("v1:", v1)

	fmt.Println("len:", len(m))
	delete(m, "k1")
	fmt.Println("map:", m)

	_, prs := m["k1"]
	fmt.Println("prs:", prs)

	prs1, _ := m["k1"]
	fmt.Println("prs1:", prs1)

	//
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
