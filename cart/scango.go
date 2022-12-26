package cart

import "bitbucket.org/niranjanawati/cart-mydesign/catalogue"

type ScanGo struct {
	UserID string
	Items  []catalogue.Product
}
