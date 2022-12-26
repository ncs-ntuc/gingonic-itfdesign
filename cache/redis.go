package cache

import (
	"fmt"

	"bitbucket.org/niranjanawati/cart-mydesign/cart"
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
		return nil, cmd.Err()
	}
	itemsAsStr, err := cmd.Result()
	if err != nil {
		return nil, err
	}
	return unjson([]byte(itemsAsStr)), nil
}
