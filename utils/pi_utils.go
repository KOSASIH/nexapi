package utils

import (
	"math"
	"math/rand"
	"time"
)

// PI represents the mathematical constant PI
const PI = 3.14159265359

// RandomPoint generates a random point within the unit circle
func RandomPoint() (x, y float64) {
	rand.Seed(time.Now().UnixNano())
	x = rand.Float64() * 2 - 1
	y = rand.Float64() * 2 - 1
	return
}

// IsPointInCircle checks if a point is within the unit circle
func IsPointInCircle(x, y float64) bool {
	return x*x+y*y <= 1
}

// CalculatePI calculates an approximation of PI using the Monte Carlo method
func CalculatePI(numPoints int) float64 {
	pointsInCircle := 0
	for i := 0; i < numPoints; i++ {
		x, y := RandomPoint()
		if IsPointInCircle(x, y) {
			pointsInCircle++
		}
	}
	return 4 * float64(pointsInCircle) / float64(numPoints)
}

// CalculatePIParallel calculates an approximation of PI using the Monte Carlo method in parallel
func CalculatePIParallel(numPoints int, numWorkers int) float64 {
	ch := make(chan int, numWorkers)
	for i := 0; i < numWorkers; i++ {
		go func() {
			pointsInCircle := 0
			for j := 0; j < numPoints/numWorkers; j++ {
				x, y := RandomPoint()
				if IsPointInCircle(x, y) {
					pointsInCircle++
				}
			}
			ch <- pointsInCircle
		}()
	}

	pointsInCircle := 0
	for i := 0; i < numWorkers; i++ {
		pointsInCircle += <-ch
	}

	return 4 * float64(pointsInCircle) / float64(numPoints)
}
