package main

import "fmt"

type Employee struct {
	Person
	EmployeeID uint
}

func (e *Employee) PrintInfo() {
	fmt.Println("Employee.Name:", e.Name)
	fmt.Println("Employee.Age:", e.Age)
	fmt.Println("Employee.ID:", e.EmployeeID)
}
