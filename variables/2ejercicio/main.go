package main

import "fmt"

func main() {
	var temperature int8
	var humidity int8
	var pressure int16
	temperature = 32
	humidity = 47
	pressure = 1013
	var apellido string = "Gomez"

	fmt.Println("Temperatura:", temperature, "ºc")
	fmt.Println("Humedad:", humidity, "%")
	fmt.Println("Presión:", pressure, "hPa")
	fmt.Println("Presión:", apellido, "hPa")

}
