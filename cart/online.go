package cart

import "bitbucket.org/niranjanawati/cart-mydesign/catalogue"

type OnlineCart struct {
	UserID string              `json:"userid"`
	Items  []catalogue.Product `json:"items"`
	Value  float32
}

func (sg *OnlineCart) Picked() []catalogue.Product {
	return sg.Items
}

func (sg *OnlineCart) AddValue(v float32) {
	sg.Value += v
}
func (sg *OnlineCart) Chargeable() []catalogue.Product {
	result := []catalogue.Product{}
	for _, item := range sg.Items {
		if item.DnmUnit() == catalogue.Box || item.DnmUnit() == catalogue.Crete {
			result = append(result, item)
		}
	}
	return result
}
