package main

import (
	"errors"
	"fmt"
)

func main() {
	average, err := calculateAverage(10, 9, 9, 10, 10)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Promedio de notas:", average)
	}
}

func calculateAverage(ratings ...float64) (average float64, err error) {
	var add float64
	var count float64
	for _, value := range ratings {
		if value < 0 {
			err = errors.New("Las notas no pueden ser negativas")
		}
		count++
		add += value
	}
	average = add / count
	return
}
