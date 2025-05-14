package main

import "fmt"

type paymenter interface {
	pay(amount float32)
}

type payment struct {
	gateway paymenter
}

func (p payment) makePayment(amount float32) {
	p.gateway.pay(amount)
}

type razorpay struct{}

func (r razorpay) pay(amount float32) {
	fmt.Println("making payment using razorpay", amount)
}

type stripe struct{}

func (s stripe) pay(amount float32) {
	fmt.Println("making payment using stripe", amount)
}

type fakepayment struct{}

func (f fakepayment) pay(amount float32) {
	fmt.Println("making payment using fake gateway for testing purpose")
}

type paypal struct{}

func (p paypal) pay(amount float32) {
	fmt.Println("making payment using paypal", amount)
}

func main() {
	stripePaymentGw := stripe{}
	// razorpayPaymentGw := razorpay{}

	// fakeGw := fakepayment{}
	///	paypalGw := paypal{}
	newPayment := payment{
		gateway: stripePaymentGw,
	}
	newPayment.makePayment(100)
}
