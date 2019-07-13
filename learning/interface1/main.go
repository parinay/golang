package main

import "fmt"

// interface

type VowelsFinder interface {
	FindVowels() []rune
}

// type alias
type MyString string

// method with reciver type of Mystring
func (ms MyString) FindVowels() []rune {
	var vowels []rune

	for _, rune := range ms {
		if rune == 'a' || rune == 'e' || rune == 'i' || rune == 'o' || rune == 'u' {
			vowels = append(vowels, rune)

		}
	}
	return vowels
}
func main() {
	name := MyString("Thomas Anderson")
	var v VowelsFinder
	v = name

	fmt.Printf("Vowels are %c\n", v.FindVowels())
}
