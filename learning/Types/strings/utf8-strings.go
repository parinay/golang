package main

import "fmt"

func utf8Strings() {
	const placeOfInterest = `⌘`

	fmt.Printf("plain string: %s \n", placeOfInterest)

	fmt.Printf("quoted string: %+q \n", placeOfInterest)

	fmt.Printf("hex bytes: ")
	for i := 0; i < len(placeOfInterest); i++ {
		fmt.Printf("%x ", placeOfInterest[i])
	}
	fmt.Printf("\n")
}
