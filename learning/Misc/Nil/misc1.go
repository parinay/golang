package main

import "fmt"

func misc1() {

	// x := nil
	// Error - use of untyped nil
	// fmt.Println(x)

	var y *int64 = nil
	fmt.Println(y)

	var a *int64 = nil
	var b *int64 = nil
	fmt.Println(a == b)
	// Error - invalid operation: a == b1 (mismatched types *int64 and *string)
	// var b1 *string = nil
	// fmt.Println(a==b1)
}
