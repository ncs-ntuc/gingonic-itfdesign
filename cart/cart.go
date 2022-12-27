package cart

import "bitbucket.org/niranjanawati/cart-mydesign/catalogue"

type CartType int

const (
	TypeOnlineCart = iota
	TypeOfflineCart
	TypeB2BCart
	TypePreOrderCart
)

type ICart interface {
	AddValue(float32)
}
type ICartItems interface {
	Picked() []catalogue.Product
	Chargeable() []catalogue.Product
}

func InitCart(ct CartType) ICart {
	switch ct {
	case TypeOnlineCart:
		return &OnlineCart{}
	case TypeOfflineCart:
		return &OnlineCart{}
	case TypeB2BCart:
		return &OnlineCart{}
	case TypePreOrderCart:
		return &OnlineCart{}
	}
	return nil
}

// ApplyDiscounts : this will run thru all the items and change the cart value
// crt 	: cart object over Icart interface
// incase of nil cart this will send back an error
func ApplyDiscounts(crt ICart) error {
	var value float32 = 0.0
	for _, item := range crt.(ICartItems).Picked() {
		value += item.Price()
	}
	crt.AddValue(value)
	return nil
}

func ApplyCharges(crt ICart) error {
	var value float32 = 0.0
	for _, item := range crt.(ICartItems).Chargeable() {
		value += (item.DeliveryCharge() + item.ServiceCharge())
	}
	crt.AddValue(value)
	return nil
}
