package main

import (
	"errors"
	"fmt"
)

const (
	small             = "Small"
	medium            = "Medium"
	large             = "Large"
	shippingCost      = 2500.0
	smallStorageCost  = 0.0
	mediumStorageCost = 0.03
	largeStorageCost  = 0.06
)

type products interface {
	price() (float64, error)
}

type product struct {
	ProductName string
	Price       float64
	Size        string
}

func (p product) price() (float64, error) {
	switch p.Size {
	case small:
		return p.Price + (p.Price * smallStorageCost), nil
	case medium:
		return p.Price + (p.Price * mediumStorageCost), nil
	case large:
		return p.Price + (p.Price * largeStorageCost) + shippingCost, nil
	default:
		return 0, errors.New("No existe el tamaño")
	}
}

func factory(price float64, size string, name string) products {
	return product{
		ProductName: name,
		Price:       price,
		Size:        size,
	}
}
func main() {
	newProduct := factory(20, medium, "Teclado")
	price, err := newProduct.price()
	if err != nil {
		panic("No existe el tamaño")
	} else {
		fmt.Printf("El precio total es $%.2f\n", price)
	}

}
