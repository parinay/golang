package main

import "fmt"

type A struct {
	alpha int8
}

type B struct {
	a *A
}

type C struct {
	b *B
}

func (b *B) GetA() *A {
	println("_GetA()")
	if b != nil {
		println("GetA()")
		return b.a
	}
	return nil
}
func (c *C) GetB() *B {
	println("_GetB()")
	if c != nil {
		println("GetB()")
		return c.b
	}
	return nil
}
func misc2() {
	beta := C{}

	fmt.Println("beta.GetB()")
	fmt.Println(beta.GetB())
	fmt.Println("beta.GetB().GetA()")
	fmt.Println(beta.GetB().GetA())
}
