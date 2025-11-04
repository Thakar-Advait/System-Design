package gateway

import "fmt"

type paypal struct{}

func (p *paypal) Process(amount int64) string {
	return fmt.Sprintf("Payment of amount %d processed by Paypal", amount)
}