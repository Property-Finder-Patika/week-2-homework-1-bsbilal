package main

import "fmt"

func main() {
	var letter byte = 'A'
	fmt.Println("an english letter:", letter)

	var unicode rune
	unicode = 'Ç'
	fmt.Println("a non-english letter:", unicode)

	var year uint16 = 2040
	fmt.Println("a year in 4 digits like 2040:", year)

	var month uint8 = 6
	fmt.Println("a month in 2 digits: 1 to 12:", month)

	var lightSpeed uint32 = 670616629 // miles
	fmt.Println("the speed of the light:", lightSpeed)

	var angle float32 = 35.8
	fmt.Println("angle of a circle:", angle)

	var pi float64
	pi = 3.141592653589793
	fmt.Println("PI number:", pi)
}
