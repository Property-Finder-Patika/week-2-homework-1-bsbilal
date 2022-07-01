package main

import "fmt"

type money float32

var beginMoney money = -20

func main() {
	fmt.Println("begin money is ", beginMoney)
	var beginMoney money = 0

	fmt.Println("second begin money is ", beginMoney)
	{
		var beginMoney money = 20
		fmt.Println("third begin money is ", beginMoney)

	}

}
