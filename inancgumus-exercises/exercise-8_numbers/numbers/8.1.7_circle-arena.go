package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		radius = 10.
		area   float64
	)
	area = math.Pi * radius * radius

	fmt.Printf("radius: %g -> area: %.2f\n",
		radius, area)
}
