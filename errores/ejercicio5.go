package main

import (
	"errors"
	"fmt"
)

func main() {
	salary, err := Salary(80, 2000)
	if err != nil {
		panic(err)
	}
	fmt.Printf("El salario es %.2f\n", salary)
}

var (
	ErrLess80Hours = errors.New("Error: el trabajador no puede haber trabajado menos de 80 hs mensuales")
)

func Salary(hours float64, perHour float64) (salary float64, err error) {
	salary = hours * perHour
	if salary >= 150000 {
		tax := salary * 0.1
		salary = salary - tax
	}
	if hours < 80 {
		err := ErrLess80Hours
		return salary, err
	}
	return
}
