package orderengine

type Platform interface {
	Normal(invoice Invoice) string
	Express(invoice Invoice) string
	Subscription(invoice Invoice) string
}

type Builder interface {
	ApplyDiscount(discount int32) Builder
	AddTax(tax int32) Builder
	AddShippingInfo(info string) Builder
	GiftWrap() Builder
	Build() Invoice
}
