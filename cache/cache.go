package cache

type ICache interface {
	ReadCart(uid string) ([]byte, error)
}
