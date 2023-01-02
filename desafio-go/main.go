package main

import (
	"desafio-go-bases/internal/tickets"
	"fmt"
)

func main() {
	err := tickets.ReadTickets("./tickets.csv")
	if err != nil {
		panic(err)
	}
	totalTickets, err := tickets.GetTotalTickets("Argentina")
	if err != nil {
		panic(err)
	}
	fmt.Println("Total de tickets al pais:", totalTickets)
	timeTickets, err := tickets.GetMornings("Noche")
	if err != nil {
		panic(err)
	}
	fmt.Println("Tickets en el turno:", timeTickets)

	averageTickets, err := tickets.AverageDestination("Argentina")
	if err != nil {
		panic(err)
	}
	fmt.Println("Porcentaje:", averageTickets)
}
