package main

import "fmt"

type student struct {
	fname   string
	lname   string
	grade   string
	country string
}

func filter(s []student, f func(student) bool) []student {
	var r []student

	for _, v := range s {
		if f(v) == true {
			r = append(r, v)
		}
	}
	return r
}
func main() {
	s1 := student{
		fname:   "Naveen",
		lname:   "Ramanathan",
		grade:   "A",
		country: "India",
	}

	s2 := student{
		fname:   "Rama",
		lname:   "Gowda",
		grade:   "A+",
		country: "USA",
	}

	s := []student{s1, s2}

	f := filter(s, func(s student) bool {
		if s.grade == "A+" {
			return true
		}
		return false
	})
	fmt.Println(f)
}
