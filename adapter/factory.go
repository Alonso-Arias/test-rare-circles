package adapter

import (
	"test/paypal"
	"test/stripe"
)

// receive the payment method will be used
func Factory(s string) Adapter {
	switch s {
	case "Paypal":
		d := paypal.Paypal{}
		return &PaypalAdapter{&d}
	case "Stripe":
		d := stripe.Stripe{}
		return &StripeAdapter{&d}
	default:
		d := stripe.Stripe{}
		return &StripeAdapter{&d}
	}
}
