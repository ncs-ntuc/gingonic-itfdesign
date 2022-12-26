package cart

import "bitbucket.org/niranjanawati/cart-mydesign/catalogue"

type ScanGo struct {
	UserID string              `json:"userid"`
	Items  []catalogue.Product `json:"items"`
}
