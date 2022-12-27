package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateSalaryA(t *testing.T) {
	//ratings := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	minutes := 60.0
	category := "A"
	expectedSalary := 4500.0

	salary, _ := calculateSalary(minutes, category)

	assert.Equal(t, expectedSalary, salary)

}

func TestCalculateSalaryB(t *testing.T) {
	//ratings := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	minutes := 60.0
	category := "B"
	expectedSalary := 1800.0

	salary, _ := calculateSalary(minutes, category)

	assert.Equal(t, expectedSalary, salary)

}
func TestCalculateSalaryC(t *testing.T) {
	//ratings := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	minutes := 60.0
	category := "C"
	expectedSalary := 1000.0

	salary, _ := calculateSalary(minutes, category)

	assert.Equal(t, expectedSalary, salary)

}
