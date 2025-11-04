package gateway

import "fmt"

func NewPaymentGateway(vertical string) (Gateway, error) {
	switch vertical {
	case "paypal":
		return &paypal{}, nil
	case "razorpay":
		return &razorpay{}, nil
	case "stripe":
		return  &stripe{}, nil
	default:
		return nil, fmt.Errorf("unsupported payment gateway %s", vertical)
	}
}