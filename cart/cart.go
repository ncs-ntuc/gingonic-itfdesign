package cart

type CartType int

const (
	OnlineCart = iota
	OfflineCart
	B2BCart
	PreOrderCart
)

type ICart interface {
}

func InitCart(ct CartType) ICart {
	switch ct {
	case OfflineCart:
		return &ScanGo{}
	}
	return nil
}
