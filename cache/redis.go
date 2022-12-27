package cache

import (
	"fmt"

	"bitbucket.org/niranjanawati/cart-mydesign/cart"
	"bitbucket.org/niranjanawati/cart-mydesign/errx"
	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
)

type RedisCache struct {
	Client *redis.Client
}

// ReadCart : given the uid this will fetch the items in the cart from the cache
func (rr *RedisCache) ReadCart(uid string, unjson func(byt []byte) cart.ICart) (cart.ICart, error) {
	key := fmt.Sprintf("cart-%s", uid)
	log.WithFields(log.Fields{
		"key": key,
	}).Info("Now reading cart from cache")
	cmd := rr.Client.Get(key)
	if cmd.Err() != nil {
		return nil, errx.Throw(errx.CacheQryErr, cmd.Err(), "RedisCache/ReadCart", "Check connection with cache", "Oops, something went wrong on the server")
	}
	itemsAsStr, err := cmd.Result()
	if err != nil {
		return nil, errx.Throw(errx.CacheQryErr, err, "RedisCache/ReadCart", "Check connection with cache", "Oops, something went wrong on the server")
	}
	return unjson([]byte(itemsAsStr)), nil
}
