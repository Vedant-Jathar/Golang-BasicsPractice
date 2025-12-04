package main

import "fmt"

type OrderStatus string

const (
	RECIEVED  OrderStatus = "recieved"
	CONFIRMED OrderStatus = "confirmed"
	DELIVERED OrderStatus = "delivered"
	PREPARED  OrderStatus = "prepared"
)

func changeOrderStatus(status OrderStatus) {
	fmt.Println("The status has been changed to: ", status)
}

func main() {
	changeOrderStatus(PREPARED)
}