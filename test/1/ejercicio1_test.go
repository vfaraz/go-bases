package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateTax0(t *testing.T) {
	salary := 10000.0
	expectedResultTax := 0.0
	expectedResultPercent := 0.0

	tax, percent := CalculateTax(salary)

	assert.Equal(t, expectedResultTax, tax)
	assert.Equal(t, expectedResultPercent, percent)

}

func TestCalculateTax17(t *testing.T) {
	salary := 100000.0
	expectedResultTax := 17000.0
	expectedResultPercent := 17.0

	tax, percent := CalculateTax(salary)

	assert.Equal(t, expectedResultTax, tax)
	assert.Equal(t, expectedResultPercent, percent)

}
func TestCalculateTax27(t *testing.T) {
	salary := 200000.0
	expectedResultTax := 54000.0
	expectedResultPercent := 27.0

	tax, percent := CalculateTax(salary)

	assert.Equal(t, expectedResultTax, tax)
	assert.Equal(t, expectedResultPercent, percent)

}
