package domains

var ProductList map[string]*Product

type Product struct {
	Name     string
	InitCost float64
	Seller   *Customer
	Buyer    *Customer
}

func init() {
	ProductList = make(map[string]*Product)
}

func NewProduct(name string, initcost float64, seller, buyer *Customer) *Product {
	product := &Product{
		Name:     name,
		InitCost: initcost,
		Seller:   seller,
		Buyer:    buyer,
	}
	return product
}
