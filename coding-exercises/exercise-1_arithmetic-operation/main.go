package main

import "fmt"

func main() {

	var firstNumber, secondNumber int = 20, 10

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

func Addition(a, b int) int {
	// that function additions the a number with a

	return a + b
}
func Subtraction(a, b int) int {
	// that function subtracts the b number from a

	return a - b
}
func Multiplication(a, b int) int {
	// that function multiplies the a with b

	return a * b
}
func Division(a, b int) int {
	// that function divides the a number to b
	return a / b
}
