package main

import (
	"errors"
	"fmt"
)

type CustomError struct {
	text string
}

func (c CustomError) Error() string {
	return c.text
}

var (
	MyError = CustomError{"Error: el salario es menor a 10.000"}
)

func main() {
	var salary int
	salary = 10001
	_, err := CheckSalary(salary)

	if errors.Is(err, MyError) {
		panic(err)
	}
	fmt.Println("El salario es mayor a 10000")
}

func CheckSalary(salary int) (b bool, err error) {
	if salary > 10000 {
		b = true
		return
	}
	err = CustomError{"Error: el salario es menor a 10.000"}
	return
}
