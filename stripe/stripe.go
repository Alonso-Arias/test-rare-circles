package stripe

type Stripe struct {
	Items    []string `json:"items"`
	Category string   `json:"category"`
}

func (p *Stripe) SendInfoPay() string {

	return "Paid with Stripe"
}
