package main

import "fmt"

func main() {
	result := convertToCelsius(305.15)
	fmt.Printf("Kelvin to Celsius = %.2f", result)
}

func convertToCelsius(kelvinValue float64) float64 {
	return kelvinValue - 273
}
