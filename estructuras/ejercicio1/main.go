package main

import "fmt"

type Product struct {
	ID          int
	Name        string
	Price       float64
	Description string
	Category    string
}

var Products []Product

func (p *Product) Save(product Product) {
	Products = append(Products, product)
}

func (p Product) GetAll() {
	fmt.Println(Products)
}

func getById(id int) (product Product) {
	for _, p := range Products {
		if p.ID == id {
			return p
		}
	}
	return
}

func main() {
	product := Product{1, "Vaso", 100, "Vaso de vidrio", "Bazar"}
	product.Save(product)
	product.GetAll()
	productid := getById(1)
	fmt.Println(productid)
}
