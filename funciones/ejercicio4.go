package main

import (
	"errors"
	"fmt"
)

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

type funcFloat64 func(...float64) float64

func main() {

	minFunc, err := operation(minimum)
	averageFunc, err := operation(average)
	maxFunc, err := operation(maximum)
	if err != nil {
		fmt.Println(err)
	} else {
		minValue := minFunc(2, 3, 3, 4, 10, 2, 4, 5)
		averageValue := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)
		maxValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)

		fmt.Println("El minimo es:", minValue)
		fmt.Println("El promedio es:", averageValue)
		fmt.Println("El maximo es:", maxValue)
	}

}

func operation(typeOp string) (function funcFloat64, err error) {
	switch typeOp {
	case minimum:
		return minFunc, err
	case average:
		return averageFunc, err
	case maximum:
		return maxFunc, err
	default:
		err = errors.New("No existe ese calculo")
		return
	}
}

func minFunc(ratings ...float64) (min float64) {
	min = ratings[0]
	for _, value := range ratings {
		if value < min {
			min = value
		}
	}
	return
}
func averageFunc(ratings ...float64) (average float64) {
	var add float64
	var count float64
	for _, value := range ratings {
		add += value
		count++
	}
	average = add / count
	return
}
func maxFunc(ratings ...float64) (max float64) {
	max = ratings[0]
	for _, value := range ratings {
		if value > max {
			max = value
		}
	}
	return
}
