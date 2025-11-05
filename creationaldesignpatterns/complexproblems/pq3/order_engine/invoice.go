package orderengine

type Invoice struct {
	Tax          int32
	Discount     int32
	ShippingInfo string
	GiftWrap     bool
}
