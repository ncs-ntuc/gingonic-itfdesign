package cart

import (
	"encoding/json"

	"bitbucket.org/niranjanawati/cart-mydesign/catalogue"
)

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

//	func (sg *OnlineCart) MarshalJSON() ([]byte, error) {
//		return nil, nil
//	}
//
// https://medium.com/@dynastymasra/override-json-marshalling-in-go-cb418102c60f
func (sg *OnlineCart) UnmarshalJSON(b []byte) error {
	// For now we are assuming everything is a Grocery item
	// moving on we can then make objects that mutate when unmarshalled as specific implementation of products
	// objects then can be pushed into cart which takes Product interface
	result := struct {
		UserId string               `json:"userid"`
		Items  []*catalogue.Grocery `json:"items"`
	}{}
	if err := json.Unmarshal(b, &result); err != nil {
		return err
	}
	sg.UserID = result.UserId
	for _, item := range result.Items {
		sg.Items = append(sg.Items, item)
	}
	return nil
}
