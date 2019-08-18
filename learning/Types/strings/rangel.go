package main

import "fmt"

func rangel() {
	const nihongo = "日本語"
	fmt.Printf("The string is - %s \n", nihongo)
	fmt.Printf("The quoted string is - %+q \n", nihongo)
	for index, runeValue := range nihongo {
		fmt.Printf("%#U %x starts at byte position %d\n", runeValue, runeValue, index)
	}
}
