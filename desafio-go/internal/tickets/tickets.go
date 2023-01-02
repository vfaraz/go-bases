package tickets

import (
	"encoding/csv"
	"errors"
	"os"
	"strconv"
	"strings"
)

type Ticket struct {
	ID                 int
	Name               string
	Email              string
	DestinationCountry string
	FlightTime         string
	Price              int
}

var (
	tickets []Ticket
)

func ReadTickets(path string) error {
	fileOpen, err := os.Open(path)
	if err != nil {
		return err
	}
	file, err := csv.NewReader(fileOpen).ReadAll()
	if err != nil {
		return err
	}
	for _, line := range file {
		price, err := strconv.Atoi(line[5])
		if err != nil {
			return err
		}
		id, err := strconv.Atoi(line[0])
		if err != nil {
			return err
		}
		tickets = append(tickets, Ticket{
			ID:                 id,
			Name:               line[1],
			Email:              line[2],
			DestinationCountry: line[3],
			FlightTime:         line[4],
			Price:              price,
		})
	}
	return err
}
func GetTotalTickets(destination string) (i int, err error) {
	for _, ticket := range tickets {
		if ticket.DestinationCountry == destination {
			i++
		}
	}
	if i == 0 {
		err = errors.New("no hay tickets a ese destino")
	}
	return
}
func GetMornings(time string) (i int, err error) {
	for _, ticket := range tickets {
		hours, _ := strconv.Atoi(strings.Split(ticket.FlightTime, ":")[0])
		switch time {
		case "Madrugada":
			if hours >= 0 && hours <= 6 {
				i++
			}
		case "Mañana":
			if hours >= 7 && hours <= 12 {
				i++
			}
		case "Tarde":
			if hours >= 13 && hours <= 19 {
				i++
			}
		case "Noche":
			if hours >= 20 && hours <= 223 {
				i++
			}
		default:
			err = errors.New("Error:Error de horario")
		}
	}
	return
}
func AverageDestination(destination string) (average float64, err error) {
	ticketsDestination, err := GetTotalTickets(destination)
	if err != nil {
		return average, err
	}
	totalTicketsDestiantion := float64(ticketsDestination)
	totalTickets := float64(len(tickets))
	average = (totalTicketsDestiantion / totalTickets) * 100
	return
}
