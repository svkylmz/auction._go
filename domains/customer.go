package domains

var CustomerList map[string]*Customer

type Customer struct {
	Name    string
	Wallet  int
	Product []*Product
}

func init() {
	CustomerList = make(map[string]*Customer)
}

func NewCustomer(name string, wallet int) *Customer {
	customer := &Customer{
		Name:   name,
		Wallet: wallet,
	}
	return customer
}
