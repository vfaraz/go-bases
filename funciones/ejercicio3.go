package main

import (
	"errors"
	"fmt"
)

func main() {
	salary, err := calculateSalary(120, "a")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("El salario es:", salary)
	}
}

func calculateSalary(minutes float64, category string) (salary float64, err error) {
	var perHour float64
	var percentBonus float64
	var bonus float64

	switch category {
	case "C":
		perHour = 1000
		percentBonus = 0
	case "B":
		perHour = 1500
		percentBonus = 0.20
	case "A":
		perHour = 3000
		percentBonus = 0.50
	default:
		err = errors.New("La categoria no existe")
	}
	salary = (minutes / 60.0) * perHour
	bonus = salary * percentBonus
	salary += bonus
	return
}
