package adapter

import "test/paypal"

type PaypalAdapter struct {
	Paypal *paypal.Paypal
}

// implementation of the methods
func (p *PaypalAdapter) Pay() {
	p.Paypal.InitRequest()
}
