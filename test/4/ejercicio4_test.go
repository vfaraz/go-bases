package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateMin(t *testing.T) {
	ratings := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expectedMin := 1.0

	min := minFunc(ratings...)

	assert.Equal(t, expectedMin, min)

}

func TestCalculateMax(t *testing.T) {
	ratings := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expectedMax := 10.0

	max := maxFunc(ratings...)

	assert.Equal(t, expectedMax, max)

}
func TestCalculateAvergae(t *testing.T) {
	ratings := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expectedAverage := 5.5

	average := averageFunc(ratings...)

	assert.Equal(t, expectedAverage, average)

}
