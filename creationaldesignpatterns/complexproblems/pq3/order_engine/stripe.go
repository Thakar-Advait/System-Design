package orderengine

import (
	"fmt"
	"sync"
)

type Stripe struct{}

type StripeFactory struct {
	once     sync.Once
	platform Platform
}

type StripeInvoiceBuilder struct {
	discount     int32
	tax          int32
	giftWrap     bool
	shippingInfo string
}

func (*StripeFactory) NewInvoiceBuilder() Builder {
	return &StripeInvoiceBuilder{}
}

func (s *StripeFactory) Service() Platform {
	s.once.Do(func() {
		s.platform = &Stripe{}
	})
	return s.platform
}

func (s *StripeInvoiceBuilder) ApplyDiscount(discount int32) Builder {
	s.discount = discount
	return s
}

func (s *StripeInvoiceBuilder) AddTax(tax int32) Builder {
	s.tax = tax
	return s
}

func (s *StripeInvoiceBuilder) GiftWrap() Builder {
	s.giftWrap = true
	return s
}

func (s *StripeInvoiceBuilder) AddShippingInfo(info string) Builder {
	s.shippingInfo = info
	return s
}

func (s *StripeInvoiceBuilder) Build() Invoice {
	return Invoice{
		Discount:     s.discount,
		Tax:          s.tax,
		ShippingInfo: s.shippingInfo,
		GiftWrap:     s.giftWrap,
	}
}

func (*Stripe) Normal(invoice Invoice) string {
	return fmt.Sprintf("Normal Stripe order placed for %v", invoice)
}

func (*Stripe) Express(invoice Invoice) string {
	return fmt.Sprintf("Express Stripe order placed for %v", invoice)
}

func (*Stripe) Subscription(invoice Invoice) string {
	return fmt.Sprintf("Subscription Stripe order placed for %v", invoice)
}
