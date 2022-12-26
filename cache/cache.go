package cache

import "bitbucket.org/niranjanawati/cart-mydesign/cart"

type ICache interface {
	ReadCart(uid string, jsonify func(byt []byte) cart.ICart) (cart.ICart, error)
}
