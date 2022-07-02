package main

import "fmt"

func main() {

	var firstNumber, secondNumber int = 10, 20

	fmt.Println("Addition for + ")
	fmt.Println("Subtraction for - ")
	fmt.Println("Multiplication for * ")
	fmt.Println("Division for / ")
	fmt.Println("Please choose arithmetic operation with char input:")

	var charInput string
	readCount, err := fmt.Scanln(&charInput)
	if err != nil || readCount < 1 || readCount > 1 {
		fmt.Println("The program faced an error")
	}

	fmt.Println("The result is : ")
	switch charInput {
	case "+":
		fmt.Println(Addition(firstNumber, secondNumber))
	case "-":
		fmt.Println(Subtraction(firstNumber, secondNumber))
	case "/":
		fmt.Println(Division(firstNumber, secondNumber))
	case "*":
		fmt.Println(Multiplication(firstNumber, secondNumber))
	}

}

func Addition(a, b int) float32 {
	// that function additions the a number with b, with float conversion
	return float32(a) + float32(b)
}
func Subtraction(a, b int) uint {
	// that function subtracts the b number from a, with uint conversion

	return uint(a) - uint(b)
}
func Multiplication(a, b int) int8 {
	// that function multiplies the a with b, with int8 conversion

	return int8(a) * int8(b)
}
func Division(a, b int) float64 {
	// that function divides the a number to b, with float64 conversion
	return float64(a) / float64(b)
}
