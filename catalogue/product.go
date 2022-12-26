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
}

type Grocery struct {
	Title   string       `json:"title"`
	Vendor  string       `json:"vendor"`
	Unit    Denomination `json:"unit"`
	PerUnit float32      `json:"price"`
}
