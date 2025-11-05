package orderengine

import (
	"fmt"
	"sync"
)

type Bank struct{}

type BankFactory struct {
	once     sync.Once
	platform Platform
}

type BankInvoiceBuilder struct {
	discount     int32
	tax          int32
	shippingInfo string
	giftWrap     bool
}

func (b *BankInvoiceBuilder) ApplyDiscount(discount int32) Builder {
	b.discount = discount
	return b
}

func (b *BankInvoiceBuilder) AddTax(tax int32) Builder {
	b.tax = tax
	return b
}

func (b *BankInvoiceBuilder) AddShippingInfo(info string) Builder {
	b.shippingInfo = info
	return b
}

func (b *BankInvoiceBuilder) GiftWrap() Builder {
	b.giftWrap = true
	return b
}

func (b *BankInvoiceBuilder) Build() Invoice {
	return Invoice{
		Discount:     b.discount,
		Tax:          b.tax,
		GiftWrap:     b.giftWrap,
		ShippingInfo: b.shippingInfo,
	}
}

func (*BankFactory) NewInvoiceBuilder() Builder {
	return &BankInvoiceBuilder{}
}

func (b *BankFactory) Service() Platform {
	b.once.Do(func() {
		b.platform = &Bank{}
	})
	return b.platform
}

func (*Bank) Normal(invoice Invoice) string {
	return fmt.Sprintf("Normal Bank order placed for %v", invoice)
}

func (*Bank) Express(invoice Invoice) string {
	return fmt.Sprintf("Express Bank order placed for %v", invoice)
}

func (*Bank) Subscription(invoice Invoice) string {
	return fmt.Sprintf("Subscription Bank order placed for %v", invoice)
}
