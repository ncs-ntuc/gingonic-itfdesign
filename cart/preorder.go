package cart

import "bitbucket.org/niranjanawati/cart-mydesign/catalogue"

type PreOrderCart struct {
	UserID string              `json:"userid"`
	Items  []catalogue.Product `json:"items"`
	Value  float32
}

func (prdr *PreOrderCart) Picked() []catalogue.Product {
	return prdr.Items
}

func (prdr *PreOrderCart) AddValue(v float32) {
	prdr.Value += v
}
func (prdr *PreOrderCart) Chargeable() []catalogue.Product {
	return []catalogue.Product{}
}
