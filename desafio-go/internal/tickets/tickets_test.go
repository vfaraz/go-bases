package tickets

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTotalTickets1(t *testing.T) {
	err := ReadTickets("/Users/vfaraz/Desktop/go-bootcamp/gobases/desafio-go-bases/tickets.csv")
	if err != nil {
		//assert.error?
		err = errors.New("no existe el archivo")
		panic(err)
	}
	destination := "Brazil"

	total, _ := GetTotalTickets(destination)
	expectedTotal := 45
	assert.Equal(t, total, expectedTotal)
}

func TestGetTotalTickets2(t *testing.T) {
	destination := "Argentina"

	total, _ := GetTotalTickets(destination)
	expectedTotal := 15
	assert.Equal(t, total, expectedTotal)
}

func TestGetMornings1(t *testing.T) {
	time := "Madrugada"

	total, _ := GetMornings(time)
	expectedTotal := 304

	assert.Equal(t, total, expectedTotal)
}

func TestGetMornings2(t *testing.T) {
	time := "Noche"

	total, _ := GetMornings(time)
	expectedTotal := 151

	assert.Equal(t, total, expectedTotal)
}

func TestAverageDestination1(t *testing.T) {
	destination := "Brazil"

	percent, _ := AverageDestination(destination)
	expectedPercent := 4.5

	assert.Equal(t, percent, expectedPercent)
}

func TestAverageDestination2(t *testing.T) {
	destination := "Argentina"

	percent, _ := AverageDestination(destination)
	expectedPercent := 1.5

	assert.Equal(t, percent, expectedPercent)
}
