package adapter

import (
	"test/stripe"
)

type StripeAdapter struct {
	Stripe *stripe.Stripe
}

// implementation of the methods
func (s *StripeAdapter) Pay() {
	s.Stripe.SendInfoPay()
}
