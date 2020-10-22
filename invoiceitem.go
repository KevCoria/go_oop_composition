package invoiceitem

//Item contains info of a invoiceitem
type Item struct {
	id      uint
	product string
	value   float64
}

//New returns a new item
func New(id uint, product string, value float64) Item {
	return Item{id, product, value}
}
