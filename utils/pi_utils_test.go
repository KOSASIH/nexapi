package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomPoint(t *testing.T) {
	x, y := RandomPoint()
	assert.True(t, x >= -1 && x <= 1)
	assert.True(t, y >= -1 && y <= 1)
}

func TestIsPointInCircle(t *testing.T) {
	tests := []struct {
		x, y float64
		want bool
	}{
		{0, 0, true},
		{1, 0, true},
		{0, 1, true},
		{-1, 0, true},
		{0, -1, true},
		{2, 0, false},
		{0, 2, false},
	}

	for _, tt := range tests {
		got := IsPointInCircle(tt.x, tt.y)
		assert.Equal(t, tt.want, got)
	}
}

func TestCalculatePI(t *testing.T) {
	tests := []struct {
		numPoints int
		want      float64
	}{
		{100, 3.14},
		{1000, 3.141},
		{10000, 3.1415},
	}

	for _, tt := range tests {
		got := CalculatePI(tt.numPoints)
		assert.InDelta(t, tt.want, got, 0.01)
	}
}

func TestCalculatePIParallel(t *testing.T) {
	tests := []struct {
		numPoints int
		numWorkers int
		want      float64
	}{
		{100, 2, 3.14},
		{1000, 4, 3.141},
		{10000, 8, 3.1415},
	}

	for _, tt := range tests {
		got := CalculatePIParallel(tt.numPoints, tt.numWorkers)
		assert.InDelta(t, tt.want, got, 0.01)
	}
}
