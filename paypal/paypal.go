package paypal

type Paypal struct {
	Items      []string `json:"items"`
	AutoReturn string   `json:"auto_return"`
	BackUrls   string   `json:"back_url"`
}

func (p *Paypal) InitRequest() string {

	return "Paid with Paypal"
}
