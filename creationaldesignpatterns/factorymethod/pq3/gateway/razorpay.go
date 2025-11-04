package gateway

import "fmt"

type razorpay struct{}

func (r *razorpay) Process(amount int64) string {
	return fmt.Sprintf("Payment of amount %d processed by Razorpay", amount)
}