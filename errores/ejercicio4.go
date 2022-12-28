package main

import (
	"errors"
	"fmt"
)

var (
	MyError = errors.New("Error: el salario es menor a 10.000")
)

func main() {
	var salary int
	salary = 100
	_, err := CheckSalary(salary)

	if errors.Is(err, MyError) {
		panic(err)
	}
	fmt.Println("El salario es mayor a 10.000")
}

func CheckSalary(salary int) (b bool, err error) {
	if salary > 10000 {
		b = true
		return
	}
	err = fmt.Errorf("Error:el mínimo imponible es de 10.000 y el salario ingresado es de: %d %w",
		salary, MyError)
	return
}
