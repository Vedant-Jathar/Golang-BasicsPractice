package main

import "fmt"

type paymenter interface {
	pay(amount float32)
}

type paymentGateway struct {
	gateway paymenter
}

type stripe struct{}

func (s stripe) pay(amount float32) {
	fmt.Println("Making payment via stripe")
}

type razorpay struct{}

func (r razorpay) pay(amount float32) {
	fmt.Println("Making payment via razorpay")
}

func main() {
	stripeGateway := stripe{}
	// razorpayGateway := razorpay{}
	paymentGateway := paymentGateway{
		gateway: stripeGateway,
	}
	paymentGateway.gateway.pay(90)
}