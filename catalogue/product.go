package catalogue

type Denomination int

const (
	Box = iota
	Packet
	Sachet
	Bottle
	Crete
	Piece
)

// Product : any item that can be added to the cart
type Product interface {
	TotalPrice() float32
	DnmUnit() Denomination
	DeliveryCharge() float32
	ServiceCharge() float32
}
