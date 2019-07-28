package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	fmt.Println(person{"Bob", 20})

	fmt.Println(person{"abcd", 30})

	fmt.Println(&person{name: "Ann", age: 40})

	fmt.Println(person{name: "fred"})

	s := person{name: "sean", age: 50}
	fmt.Println("s.name", s.name)
	fmt.Println("s.age", s.age)

	fmt.Println("Without a pointer")
	sp := s

	fmt.Println(s)
	p := &s
	fmt.Printf("Address of s:%p\n", p)
	fmt.Println(sp)
	q := &sp
	fmt.Printf("Address of sp:%p\n", q)

	fmt.Println("With a pointer")
	// with pominter
	sp1 := &s

	// sp1.age = 51
	// sp1.name = "efgh"

	fmt.Println(s)
	fmt.Printf("Address of s:%p\n", p)
	fmt.Println(*sp1)
	r := &sp1
	fmt.Printf("Address of sp1:%p\n", r)
}
