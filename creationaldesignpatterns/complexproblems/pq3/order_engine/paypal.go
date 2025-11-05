package orderengine

import (
	"fmt"
	"sync"
)

type PayPal struct{}

type PayPalFactory struct {
	once    sync.Once
	plaform Platform
}

func (*PayPalFactory) NewInvoiceBuilder() Builder {
	return &PayPalInvoiceBuilder{}
}

func (p *PayPalFactory) Service() Platform {
	p.once.Do(func() {
		p.plaform = &PayPal{}
	})
	return p.plaform
}

type PayPalInvoiceBuilder struct {
	discount     int32
	tax          int32
	shippingInfo string
	giftWrap     bool
}

func (p *PayPalInvoiceBuilder) ApplyDiscount(discount int32) Builder {
	p.discount = discount
	return p
}

func (p *PayPalInvoiceBuilder) AddTax(tax int32) Builder {
	p.tax = tax
	return p
}

func (p *PayPalInvoiceBuilder) AddShippingInfo(info string) Builder {
	p.shippingInfo = info
	return p
}

func (p *PayPalInvoiceBuilder) GiftWrap() Builder {
	p.giftWrap = true
	return p
}

func (p *PayPalInvoiceBuilder) Build() Invoice {
	return Invoice{
		Discount:     p.discount,
		Tax:          p.tax,
		GiftWrap:     p.giftWrap,
		ShippingInfo: p.shippingInfo,
	}
}

func (*PayPal) Normal(invoice Invoice) string {
	return fmt.Sprintf("Paypal order placed for %v", invoice)
}

func (*PayPal) Express(invoice Invoice) string {
	return fmt.Sprintf("Express Paypal order placed for %v", invoice)
}

func (*PayPal) Subscription(invoice Invoice) string {
	return fmt.Sprintf("Subscription Paypal order placed for %v", invoice)
}
