package integration

import (
	"context"
	"fmt"
	"math"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	piCalculationTimeout = 10 * time.Second
)

func TestPIIntegration(t *testing.T) {
	tests := []struct {
		name        string
		numPoints   int
		expectedPI float64
	}{
		{"Small", 100, 3.14},
		{"Medium", 1000, 3.141},
		{"Large", 10000, 3.1415},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), piCalculationTimeout)
			defer cancel()

			actualPI, err := CalculatePI(ctx, tt.numPoints)
			require.NoError(t, err)

			assert.InDelta(t, tt.expectedPI, actualPI, 0.01)
		})
	}
}

func BenchmarkPIIntegration(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), piCalculationTimeout)
		defer cancel()

		_, err := CalculatePI(ctx, 10000)
		if err!= nil {
			b.Errorf("Expected CalculatePI to succeed, but got error: %s", err)
		}
	}
}

func CalculatePI(ctx context.Context, numPoints int) (float64, error) {
	ch := make(chan float64, numPoints)
	errCh := make(chan error, 1)

	go func() {
		defer close(ch)
		defer close(errCh)

		for i := 0; i < numPoints; i++ {
			x := float64(i) / float64(numPoints-1)
			y := 4 * (1 - x*x)
			ch <- y
		}
	}()

	go func() {
		defer close(errCh)

		var sum float64
		for y := range ch {
			sum += y
		}

		pi := sum / float64(numPoints)
		errCh <- nil
	}()

	select {
	case <-ctx.Done():
		return 0, ctx.Err()
	case err := <-errCh:
		return 0, err
	case pi := <-ch:
		return pi, nil
	}
}
