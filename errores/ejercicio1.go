package main

import (
	"fmt"
)

type CustomError struct {
	text string
}

func (c CustomError) Error() string {
	return c.text
}

func main() {
	var salary int
	salary = 100000
	err := CustomError{"error: el salario ingresado no alcanza el mínimo imponible"}
	if salary < 150000 {
		panic(err)
	}
	fmt.Println("Debe pagar impuestos")

}
