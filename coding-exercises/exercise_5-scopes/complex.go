package main

import "fmt"

type money float32

var beginMoney money = -20

func main() {
	fmt.Println("begin money is ", beginMoney)
	var beginMoney money = 0
	cleanMoney := 25
	var hardMoney money = 100

	fmt.Println(beginMoney == hardMoney)

	fmt.Println("second begin money is ", beginMoney)
	{
		cleanMoney = 30
		fmt.Println(cleanMoney)
		hardMoney := 30
		fmt.Println("hard money created")
		fmt.Println(hardMoney)
		var beginMoney money = 20
		fmt.Println("third begin money is ", beginMoney)

	}

}
