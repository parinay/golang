package main

import "fmt"

type author struct {
	firstname string
	lastname  string
	bio       string
}

func (a author) fullname() string {
	return fmt.Sprintf("%s %s", a.firstname, a.lastname)
}

type post struct {
	title   string
	content string
	author
}

func (p post) details() {
	fmt.Println("Title", p.title)
	fmt.Println("Content:", p.content)
	fmt.Println("Author:", p.fullname())
	fmt.Println("Bio ", p.bio)
}

type website struct {
	posts []post
}

func (w website) contents() {
	fmt.Println("Contents of website")
	for _, v := range w.posts {
		v.details()
		fmt.Println()
	}
}

func main() {

	author1 := author{
		"Naveen",
		"Ramnathan",
		"Golang Enthusiast",
	}
	post1 := post{
		"Inheritance in Go",
		"Go supports composition instead of inheritance",
		author1,
	}
	post1.details()

	post2 := post{
		"Struct instead of Classes in Go",
		"Go does not support classes but methods can be added to structs",
		author1,
	}
	post3 := post{
		"Concurrency",
		"Go is a concurrent language and not a parallel one",
		author1,
	}
	w := website{
		posts: []post{post1, post2, post3},
	}
	w.contents()
}
