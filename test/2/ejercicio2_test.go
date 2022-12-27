package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateAverage(t *testing.T) {
	//ratings := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expectedAverage := 5.5

	average, _ := calculateAverage(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	assert.Equal(t, expectedAverage, average)

}
