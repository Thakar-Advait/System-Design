package gateway

import "fmt"

type stripe struct{}

func (s *stripe) Process(amount int64) string {
	return fmt.Sprintf("Payment of amount %d processed by Stripe", amount)
}
