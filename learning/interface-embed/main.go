package main

import "fmt"

type SalaryCalc interface {
	DisplaySalary()
}

type LeaveCalc interface {
	CalculateLeavesLeft() int
}

type EmployeeOperations interface {
	SalaryCalc
	LeaveCalc
}

type Employee struct {
	fname          string
	lname          string
	basicPay       int
	pf             int
	totalLeaves    int
	leavesApproved int
}

func (e Employee) DisplaySalary() {
	fmt.Printf("%s %s has salary $%d\n", e.fname, e.lname, e.basicPay+e.pf)
}

func (e Employee) CalculateLeavesLeft() int {
	return e.totalLeaves - e.leavesApproved
}

func main() {
	e := Employee{
		fname:          "naveen",
		lname:          "R",
		basicPay:       10000,
		pf:             200,
		totalLeaves:    30,
		leavesApproved: 5,
	}
	var empOp EmployeeOperations = e

	empOp.DisplaySalary()
	fmt.Println("Leaves Left =", empOp.CalculateLeavesLeft())
}
