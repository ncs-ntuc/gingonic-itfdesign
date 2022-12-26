package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"bitbucket.org/niranjanawati/cart-mydesign/cache"
	"bitbucket.org/niranjanawati/cart-mydesign/cart"
	"github.com/gin-gonic/gin"
)

// CachedCart : middleware to get cart object unmarshalled from cache
// will connect to Redis and get the value of the cart for the user Id
// will inject the cart object into context for downstream middleware functions
// cache		: since we want the middleware to be agnostic of the underlying cart,we get an interface to cache
func CachedCart(cache cache.ICache) gin.HandlerFunc {
	return func(c *gin.Context) {
		// +++++++++ getting the url query parameter to get the cart type
		val := c.Query("cart") // gets the carttypefrom the url string
		crtType, err := strconv.Atoi(val)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		// +++++++++ reading from cache
		byt, err := cache.ReadCart(c.Param("userid"))
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		// ++++++++++++ unmarshalling the cart
		crt := cart.InitCart(cart.CartType(crtType))
		json.Unmarshal(byt, crt) // TODO: handle unmarshalling error later
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		c.Set("cachedcart", crt)
	}
}
func HndlUsrCart(c *gin.Context) {
	cartVal, ok := c.Get("cachedcart") // middleware wil capture the cart value
	if ok != true {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	cart := cartVal.(cart.ICart)
	if c.Request.Method == "GET" {
		/* This is when after the user is done shopping proceeds to get the cart
		detailed view. This will load all the items on the cart with their respective pricing, taxes and the validation
		*/
		c.JSON(http.StatusOK, cart)
		return
	}
}
