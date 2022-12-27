package main

import "fmt"

func main() {
	var salary float64
	salary = 1000000
	tax, percent := CalculateTax(salary)
	fmt.Printf("El impuesto de un sueldo de $%.1f es de %.1f (%.1f porciento)\n",
		salary, tax, percent)
}

func CalculateTax(salary float64) (tax float64, percent float64) {
	if salary > 50000 {
		percent += 17
	}
	if salary > 150000 {
		percent += 10
	}
	tax = (salary / 100) * percent
	return
}
