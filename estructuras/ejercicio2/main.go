package main

import (
	"fmt"
	"time"
)

type Person struct {
	ID          int
	Name        string
	DateOfBirth time.Time
}

type Employe struct {
	ID       int
	Position string
	Person   Person
}

func (e Employe) PrintEmployee() {
	fmt.Println(e)
}

func main() {
	person := Person{1, "Valentin", time.Date(1999, 7, 7, 0, 0, 0, 0, time.Local)}
	employe := Employe{1, "Software Developer", person}
	employe.PrintEmployee()

}
