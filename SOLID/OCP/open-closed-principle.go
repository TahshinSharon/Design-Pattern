package main

import "fmt"


type PaymentMethod interface{
	Pay(amount int)
}

type BkashPayment struct{}

func(b BkashPayment) Pay(amount int){
	fmt.Printf("Paying through Bkash of Tk: %d\n", amount)
}
type DebitCardPayment struct{}
func(dc DebitCardPayment) Pay(amount int){
	fmt.Printf("Paying through DebitCard of Tk: %d\n", amount)
}
type PayPalPayment struct{}
func(dc PayPalPayment) Pay(amount int){
	fmt.Printf("Paying through PayPal of Tk: %d\n", amount)
}
type PaymentProcessor struct{}
func(pp PaymentProcessor) ProcessPayment(paymentMethod PaymentMethod,amount int){
	paymentMethod.Pay(amount)
}
func main(){
	bkash:=BkashPayment{}
	debit := DebitCardPayment{}

	paypal := PayPalPayment{}
	paymentProcessor := PaymentProcessor{}
	paymentProcessor.ProcessPayment(bkash, 500)
	paymentProcessor.ProcessPayment(debit, 1000)
	paymentProcessor.ProcessPayment(paypal, 700)
}