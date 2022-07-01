package main

import (
	"fmt"
)

const (
	canceled = iota
	pending
	accepted
	shipped
	delivered
)

func main() {

	var numberInput int
	fmt.Println("please enter your order number:")
	fmt.Scan(&numberInput)
	TrackOrder(numberInput)
}
func TrackOrder(i int) {
	// this func returns of order result
	switch i {
	case canceled:
		fmt.Println("This order has cancelled.")
	case pending:
		fmt.Println("Order is awaiting approval.")

	case accepted:
		fmt.Println("This order has accepted.")

	case shipped:
		fmt.Println("Ya! Order has shipped.")

	case delivered:
		fmt.Println("Your order has delivered.")

	default:
		fmt.Println("Package not found")
	}
}
