package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

func main() {

	// Various ways to initializa struct in go
	var a Student
	a.Name = "abcd"
	a.Age = 10
	fmt.Printf("sturct a {%s %d}\n", a.Name, a.Age)

	//
	b := Student{
		Name: "efgh",
		Age:  11,
	}
	fmt.Println("struct b", b)

	//
	c := Student{"ijkl", 12}
	fmt.Println("struct c", c)

	// emptry struct
	d := Student{}
	fmt.Printf("struct d {%s %d}\n", d.Name, d.Age)

	// pointer to struct
	var pa = new(Student)
	pa.Name = "Alice"
	pa.Age = 6
	fmt.Printf("struct pa {%s %d}\n", pa.Name, pa.Age)

	fmt.Println(c == d)
}
