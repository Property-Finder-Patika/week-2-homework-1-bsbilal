package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("String to Int : ")
	fmt.Print(StringToInt("28"))

	fmt.Println("Int to String : ")
	fmt.Print(IntToString(28))

	b, err := strconv.ParseBool("true")
	fmt.Println(b)
	f, err := strconv.ParseFloat("3.1415", 32)
	fmt.Println(f)

	i, err := strconv.ParseInt("-42", 10, 32)
	fmt.Println(i)

	u, err := strconv.ParseUint("42", 10, 64)
	fmt.Println(u)

	if err == nil {

		fmt.Println("All worked")

	}

}

func StringToInt(s string) int {
	// string to int with Atoi func

	i, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return i
}

func IntToString(s int) string {
	// int to string with Itoa func

	i := strconv.Itoa(s)
	return i
}

//Inspired from https://pkg.go.dev/strconv#pkg-overview
