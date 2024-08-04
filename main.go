package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"time"

	"github.com/spf13/viper"
	"golang.org/x/exp/rand"

	"pi/utils"
)

var (
	configFile = flag.String("config", "config.yaml", "Path to the configuration file")
	numPoints  = flag.Int("num-points", 10000, "Number of points to use for PI calculation")
	numWorkers = flag.Int("num-workers", 4, "Number of workers to use for parallel PI calculation")
)

func main() {
	flag.Parse()

	// Load configuration from file
	viper.SetConfigFile(*configFile)
	err := viper.ReadInConfig()
	if err!= nil {
		log.Fatal(err)
	}

	// Initialize random number generator
	rand.Seed(time.Now().UnixNano())

	// Calculate PI using the Monte Carlo method
	startTime := time.Now()
	pi := utils.CalculatePIParallel(*numPoints, *numWorkers)
	elapsedTime := time.Since(startTime)

	fmt.Printf("PI: %.6f\n", pi)
	fmt.Printf("Error: %.6f\n", math.Abs(pi-math.Pi))
	fmt.Printf("Time: %v\n", elapsedTime)

	// Print configuration
	fmt.Println("Configuration:")
	fmt.Println(" Num points:", *numPoints)
	fmt.Println(" Num workers:", *numWorkers)
	fmt.Println(" Config file:", *configFile)
}
