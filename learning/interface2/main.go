package main

import "fmt"

type Permanent struct {
	empId int
	bPay  int
	pf    int
}

type Contract struct {
	empId int
	bpay  int
}

type Freelancer struct {
	empId     int
	noOfHours int
}

// Interface that implements method CalculateSal()
// returns int
type SalaryCalc interface {
	CalculateSal() int
}

// method CalculateSal() value receiver
func (c Contract) CalculateSal() int {
	return c.bpay
}

// method CalculateSal() value receiver
func (p Permanent) CalculateSal() int {
	return p.bPay + p.pf
}

// method CalculateSal() value receiver
func (f Freelancer) CalculateSal() int {
	return f.noOfHours * 140
}

// functional vall with value slice of structs
// printing the type and value
func describe(s []SalaryCalc) {
	for _, v := range s {
		fmt.Printf("Interface type of %T value %v\n", v, v)
	}
}

// function call with struct slice returning int
func totalExpense(s []SalaryCalc) int {
	expense := 0
	// loops over SalaryCalc slice and as per the struct type
	// either Permanent{} OR {}Contract{} calls method
	// CalculateSal() and returns the expense per employee.
	// Easy to add many more suchi employye types OR methods
	// which implement the interface SalaryCalc{}
	for _, v := range s {
		expense = expense + v.CalculateSal()
	}
	// return total expense
	return expense
}
func main() {
	p1 := Permanent{100, 1000, 10}
	p2 := Permanent{101, 1050, 10}

	c1 := Contract{102, 800}

	f1 := Freelancer{103, 176}

	employees := []SalaryCalc{p1, p2, c1, f1}

	fmt.Printf("Total expense Per Month $%d\n", totalExpense(employees))
	describe(employees)

}
