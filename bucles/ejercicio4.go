package main

import "fmt"

func main() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26,
		"Brenda": 19, "Darío": 44, "Pedro": 30}

	fmt.Println(employees)
	fmt.Println("La edad de Benjamin es:", employees["Benjamin"])
	c := 0
	for _, value := range employees {
		//fmt.Println(key, value)
		if value > 21 {
			c++
		}
	}
	fmt.Println("Hay", c, "empleados con más de 21 años")

	//Agregamos y eliminamos un empleado
	employees["Federico"] = 25
	delete(employees, "Pedro")
	fmt.Println(employees)
}
