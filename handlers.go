package main

/* All the gin handler function implementations go here.
the handlers that are customizable in compile time are delegates that return handlers
Author 	: niranjan.awati@ntucenterprise.sg
Date 	: 27-DEC-2022
*/
import (
	"encoding/json"
	"net/http"
	"strconv"

	"bitbucket.org/niranjanawati/cart-mydesign/cache"
	"bitbucket.org/niranjanawati/cart-mydesign/cart"
	"bitbucket.org/niranjanawati/cart-mydesign/errx"
	"github.com/gin-gonic/gin"
)

// CachedCart : middleware to get cart object unmarshalled from cache
// will connect to Redis and get the value of the cart for the user Id
// will inject the cart object into context for downstream middleware functions
// Will Apply surcharges and also discounts on cart
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
		crt, err := cache.ReadCart(c.Param("userid"), func(byt []byte) cart.ICart {
			crt := cart.InitCart(cart.CartType(crtType))
			if json.Unmarshal(byt, crt) != nil {
				return nil
			}
			return crt
		})
		if err != nil {
			// c.AbortWithStatus(http.StatusBadRequest)
			// return
			errx.DigestErr(c, err)
			return
		}
		// Calculate the discount and service charges here
		if err := cart.ApplyDiscounts(crt); err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		if err := cart.ApplyCharges(crt); err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		c.Set("cachedcart", crt)
	}
}
func HndlUsrCart(c *gin.Context) {
	cartVal, ok := c.Get("cachedcart") // middleware wil capture the cart value
	if !ok {
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
