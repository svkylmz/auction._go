package main

import (
	"auction_go/domains"
	"fmt"
)

func main() {
	sadan := domains.NewCustomer("Şadan Kuzgun", 10000)
	sevket := domains.NewCustomer("Şevket Yılmaz", 10000)

	fmt.Println(sadan)
	fmt.Println(sevket)

	firstProduct := domains.NewProduct("100 yıllık şamdan", 5750.5, sadan, sevket)

	fmt.Println(firstProduct)
}
