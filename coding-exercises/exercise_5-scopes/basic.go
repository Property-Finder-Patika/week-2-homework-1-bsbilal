package main

import "fmt"

type Celsius float32

const (
	boilingTemperature Celsius = 100
)

var currentTemperature Celsius = 20.2

func main() {

	fmt.Println("Current temperature is ", currentTemperature)
	fmt.Println("Boiling temperature is ", boilingTemperature)
	SetCurrentTemperature(32.7)
	fmt.Println("Current temperature is ", currentTemperature)

}
func SetCurrentTemperature(t Celsius) {
	currentTemperature = t
}
