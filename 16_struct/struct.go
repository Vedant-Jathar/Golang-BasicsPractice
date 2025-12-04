package main

import (
	"fmt"
	"time"
)

type order struct {
	id        int
	amount    float32
	status    string
	createdAt time.Time
}

func (o *order) changeStatus(status string) {
	o.status = status
}

func (o order) getAmount() float32 {
	return o.amount
}

func main() {

	myOrder1 := order{
		id:        1,
		amount:    89,
		status:    "delivered",
		createdAt: time.Now(),
	}

	myOrder1.changeStatus("recieved")

	fmt.Println(myOrder1)
	fmt.Println(myOrder1.getAmount())
}