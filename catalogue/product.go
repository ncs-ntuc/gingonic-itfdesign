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
	Price() float32
	DnmUnit() Denomination
	DeliveryCharge() float32
	ServiceCharge() float32
}

// TODO: we havent added the quantity here - we assume for now the user will add only 1 qty of each prodcut
type Grocery struct {
	Title    string       `json:"title"`
	Vendor   string       `json:"vendor"`
	Unit     Denomination `json:"unit"`
	PerUnit  float32      `json:"price"`
	Discount float32      `json:"discount"`
}

func (grc *Grocery) Price() float32 {
	return grc.PerUnit * grc.Discount
}
func (grc *Grocery) DnmUnit() Denomination {
	return grc.Unit
}
func (grc *Grocery) DeliveryCharge() float32 {
	if grc.Unit == Box || grc.Unit == Crete {
		return grc.PerUnit * 0.2
	} else {
		return 0.0
	}
}
func (grc *Grocery) ServiceCharge() float32 {
	if grc.Unit == Box || grc.Unit == Crete {
		return grc.PerUnit * 0.15
	} else {
		return 0.0
	}
}
