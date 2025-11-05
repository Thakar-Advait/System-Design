package orderengine

import (
	"fmt"
)

type Factory interface {
	NewInvoiceBuilder() Builder
	Service() Platform
}

func NewOrderFactory(vertical string) (Factory, error) {
	switch vertical {
	case "bank":
		return &BankFactory{}, nil
	case "paypal":
		return &PayPalFactory{}, nil
	case "stripe":
		return &StripeFactory{}, nil
	default:
		return nil, fmt.Errorf("unsupported platform")
	}
}
